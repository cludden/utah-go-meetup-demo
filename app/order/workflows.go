package order

import (
	"fmt"
	"time"

	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"github.com/temporalio/reference-app-orders-go/internal/temporalutil"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Aggressively low for demo purposes.
const customerActionTimeout = 3 * time.Minute

type Workflows struct{}

type OrderWorkflow struct {
	*orderv1.OrderWorkflowInput
	logger log.Logger
	order  *omsv1.Order
}

func (w *Workflows) Order(ctx workflow.Context, input *orderv1.OrderWorkflowInput) (orderv1.OrderWorkflow, error) {
	return &OrderWorkflow{
		OrderWorkflowInput: input,
		logger:             log.With(workflow.GetLogger(ctx), "orderId", input.Req.GetId(), "customerId", input.Req.GetCustomerId()),
	}, nil
}

func (wf *OrderWorkflow) Execute(ctx workflow.Context) (*orderv1.CreateOrderResult, error) {
	// write order record to database
	order, err := orderv1.CreateOrder(ctx, wf.Req)
	if err != nil {
		return nil, err
	}
	wf.order = order.GetOrder()

	// locate and reserve items, create fulfillment for each reservation
	fulfillments, err := orderv1.ReserveItems(ctx, &orderv1.ReserveItemsInput{
		OrderId:    wf.order.GetId(),
		CustomerId: wf.Req.GetCustomerId(),
		Items:      wf.Req.GetItems(),
	})
	if err != nil {
		return nil, err
	}
	wf.order.Fulfillments = fulfillments.GetFulfillments()

	// if there is problem with the oder, await customer action
	if wf.customerActionRequired() {
		err = wf.UpdateOrderStatus(ctx, &orderv1.UpdateOrderStatusInput{Status: omsv1.OrderStatus_ORDER_STATUS_CUSTOMER_ACTION_REQUIRED})
		if err != nil {
			return nil, err
		}

		action, err := wf.waitForCustomer(ctx)
		if err != nil {
			return nil, err
		}

		switch action {
		case omsv1.CustomerAction_CUSTOMER_ACTION_CANCEL:
			err := wf.UpdateOrderStatus(ctx, &orderv1.UpdateOrderStatusInput{Status: omsv1.OrderStatus_ORDER_STATUS_CANCELLED})
			wf.cancelAllFulfillments()
			return &orderv1.CreateOrderResult{Order: wf.order}, err
		case omsv1.CustomerAction_CUSTOMER_ACTION_TIMED_OUT:
			err := wf.UpdateOrderStatus(ctx, &orderv1.UpdateOrderStatusInput{Status: omsv1.OrderStatus_ORDER_STATUS_TIMED_OUT})
			wf.cancelAllFulfillments()
			return &orderv1.CreateOrderResult{Order: wf.order}, err
		case omsv1.CustomerAction_CUSTOMER_ACTION_AMEND:
			wf.cancelUnavailableFulfillments()
		default:
			return nil, fmt.Errorf("unhandled customer action %q", action)
		}
	}

	// mark order as processing
	if err := wf.UpdateOrderStatus(ctx, &orderv1.UpdateOrderStatusInput{Status: omsv1.OrderStatus_ORDER_STATUS_PROCESSING}); err != nil {
		return nil, err
	}

	// process status updates from shipment service in the background
	workflow.Go(ctx, wf.handleShipmentStatusUpdates)

	// process each fulfillment in a separate coroutine and wait for all to
	// be delivered
	g := temporalutil.NewErrGroup(ctx)
	for _, f := range wf.order.GetFulfillments() {
		g.Go(ctx, func(ctx workflow.Context) error {
			return wf.processFulfillment(ctx, f)
		})
	}
	if err := g.Wait(ctx); err != nil {
		wf.logger.Error("failure processing fulfillments", "error", err)
	}

	// update order status
	status := omsv1.OrderStatus_ORDER_STATUS_COMPLETED
	if wf.allFulfillmentsFailed() {
		status = omsv1.OrderStatus_ORDER_STATUS_FAILED
	}
	if err := wf.UpdateOrderStatus(ctx, &orderv1.UpdateOrderStatusInput{Status: status}); err != nil {
		return nil, err
	}

	return &orderv1.CreateOrderResult{Order: wf.order}, nil
}

// GetStatus returns information about the current order status
func (wf *OrderWorkflow) GetStatus() (*orderv1.GetOrderResult, error) {
	return &orderv1.GetOrderResult{Order: wf.order}, nil
}

// UpdateOrderStatus updates the status of the workflow
func (wf *OrderWorkflow) UpdateOrderStatus(ctx workflow.Context, input *orderv1.UpdateOrderStatusInput) error {
	if input.GetId() == "" {
		input.Id = wf.Req.GetId()
	}
	if err := orderv1.UpdateOrderStatus(ctx, input); err != nil {
		return err
	}
	wf.order.Status = input.GetStatus()
	return nil
}

// allFulfillmentsFailed determines if all fulfillments have a failed status
func (wf *OrderWorkflow) allFulfillmentsFailed() bool {
	failures := 0
	for _, f := range wf.order.GetFulfillments() {
		if f.Status == omsv1.FulfillmentStatus_FULFILLMENT_STATUS_FAILED {
			failures++
		}
	}

	return failures >= 1 && failures == len(wf.order.GetFulfillments())
}

// cancelUnavailableFulfillments updates the status of any fulfillment with
// status UNAVAILABLE to CANCELLED
func (wf *OrderWorkflow) cancelUnavailableFulfillments() {
	wf.logger.Info("Cancelling unavailable fulfillments")

	for _, f := range wf.order.GetFulfillments() {
		if f.Status == omsv1.FulfillmentStatus_FULFILLMENT_STATUS_UNAVAILABLE {
			f.Status = omsv1.FulfillmentStatus_FULFILLMENT_STATUS_CANCELLED
		}
	}
}

// cancelAllFulfillments updates the status of all fulfillments to CANCELLED
func (wf *OrderWorkflow) cancelAllFulfillments() {
	wf.logger.Info("Cancelling all fulfillments")

	for _, f := range wf.order.GetFulfillments() {
		f.Status = omsv1.FulfillmentStatus_FULFILLMENT_STATUS_CANCELLED
	}
}

// customerActionRequired determines if any fulfillments have status UNAVAILABLE
func (wf *OrderWorkflow) customerActionRequired() bool {
	for _, f := range wf.order.GetFulfillments() {
		if f.GetStatus() == omsv1.FulfillmentStatus_FULFILLMENT_STATUS_UNAVAILABLE {
			return true
		}
	}
	return false
}

// handleShipmentStatusUpdates listens for ShipmentStatusUpdated signals from the
// courier and updates the corresponding fulfillment
func (wf *OrderWorkflow) handleShipmentStatusUpdates(ctx workflow.Context) {
	for {
		signal, _ := wf.ShipmentStatusUpdated.Receive(ctx)
		wf.logger.Info("Shipment status updated", "shipmentId", signal.GetShipmentId(), "status", signal.GetStatus().String())
		for _, f := range wf.order.GetFulfillments() {
			if f.GetShipment().GetId() == signal.GetShipmentId() {
				f.Shipment.Status = signal.Status
				f.Shipment.UpdatedAt = signal.UpdatedAt
				break
			}
		}
	}
}

// processFulfillment charges the customer for the items contained in this
// fulfillment, creates a shipment and waits until it has been delivered
func (wf *OrderWorkflow) processFulfillment(ctx workflow.Context, f *omsv1.Fulfillment) error {
	if f.Status == omsv1.FulfillmentStatus_FULFILLMENT_STATUS_CANCELLED {
		return nil
	}

	// update fulfillment details
	logger := log.With(wf.logger, "fulfillment", f.GetId())
	f.Status = omsv1.FulfillmentStatus_FULFILLMENT_STATUS_PROCESSING
	f.Payment = &omsv1.Payment{Status: omsv1.PaymentStatus_PAYMENT_STATUS_PENDING}

	// charge the customer for the items in this fulfillment
	charge, err := billingv1.ChargeChild(ctx, &billingv1.ChargeInput{
		CustomerId:    f.GetCustomerId(),
		OrderId:       wf.order.GetId(),
		FulfillmentId: f.GetId(),
		Reference:     f.GetId(),
		Items:         f.GetItems(),
	})
	if err != nil {
		f.Payment.Status = omsv1.PaymentStatus_PAYMENT_STATUS_FAILED
		return fmt.Errorf("failure charging customer: %w", err)
	}

	// update fulfillment payment summary
	logger.Info("Payment processed", "total", f.Payment.GetTotal(), "status", f.Payment.GetStatus().String(), "auth_code", charge.GetAuthCode())
	f.Payment.SubTotal = charge.GetSubTotal()
	f.Payment.Tax = charge.GetTax()
	f.Payment.Shipping = charge.GetShipping()
	f.Payment.Total = charge.GetTotal()
	if charge.GetSuccess() {
		f.Payment.Status = omsv1.PaymentStatus_PAYMENT_STATUS_SUCCESS
	} else {
		f.Payment.Status = omsv1.PaymentStatus_PAYMENT_STATUS_FAILED
	}

	f.Shipment = &omsv1.Shipment{
		Id:        f.GetId(),
		Status:    omsv1.ShipmentStatus_SHIPMENT_STATUS_PENDING,
		UpdatedAt: timestamppb.New(workflow.Now(ctx)),
	}

	// process shipment
	shipment, err := shipmentv1.ShipmentChild(ctx, &shipmentv1.CreateShipmentInput{
		CustomerId:    wf.order.GetCustomerId(),
		OrderId:       wf.order.GetId(),
		FulfillmentId: f.GetId(),
		RequestorWid:  workflow.GetInfo(ctx).WorkflowExecution.ID,
		Id:            f.GetShipment().GetId(),
		Items:         f.GetItems(),
	})
	if err != nil {
		return fmt.Errorf("failure processing shipment: %w", err)
	}
	logger.Info("Shipment processed", "courier", shipment.GetCourierReference(), "status", f.GetShipment().GetStatus().String())

	f.Status = omsv1.FulfillmentStatus_FULFILLMENT_STATUS_COMPLETED
	return nil
}

func (wf *OrderWorkflow) waitForCustomer(ctx workflow.Context) (action omsv1.CustomerAction, err error) {
	wf.logger.Info("Waiting for customer action")

	timerCtx, cancelTimer := workflow.WithCancel(ctx)
	workflow.NewSelector(ctx).
		AddFuture(workflow.NewTimer(timerCtx, customerActionTimeout), func(f workflow.Future) {
			if err = f.Get(timerCtx, nil); err != nil {
				return
			}

			wf.logger.Info("Timed out waiting for customer action", "timeout", customerActionTimeout)
			action = omsv1.CustomerAction_CUSTOMER_ACTION_TIMED_OUT
		}).
		AddReceive(wf.CustomerAction.Channel, func(workflow.ReceiveChannel, bool) {
			action = wf.CustomerAction.ReceiveAsync().GetAction()
			wf.logger.Info("Received customer action", "action", action.String())
			cancelTimer()
		}).
		Select(ctx)
	return action, err
}
