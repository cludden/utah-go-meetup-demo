package order_test

import (
	"context"
	"testing"
	"time"

	gocmp "github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/temporalio/reference-app-orders-go/app/order"
	"github.com/temporalio/reference-app-orders-go/app/shipment"
	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"github.com/temporalio/reference-app-orders-go/internal/interceptors"
	billingv1mocks "github.com/temporalio/reference-app-orders-go/mocks/gen/oms/billing/v1"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkflowSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	ctx      context.Context
	env      *testsuite.TestWorkflowEnvironment
	billing  *billingv1.TestWorkerClient
	orders   *orderv1.TestWorkerClient
	shipment *shipmentv1.TestWorkerClient
}

func TestWorkflowSuite(t *testing.T) {
	suite.Run(t, new(WorkflowSuite))
}

func (s *WorkflowSuite) SetupTest() {
	s.T().Setenv("TEMPORAL_DEBUG", "true")
	s.env = s.NewTestWorkflowEnvironment()
	s.env.SetWorkerOptions(worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{interceptors.NewValidation()},
	})
	s.ctx = context.Background()
	s.billing = billingv1.NewTestWorkerClient(s.env, billingv1mocks.NewMockBillingWorkflows(s.T()), nil)
	s.orders = orderv1.NewTestWorkerClient(s.env, &order.Workflows{}, &order.Activities{})
	s.shipment = shipmentv1.NewTestWorkerClient(s.env, &shipment.Workflows{}, nil)
}

func (s *WorkflowSuite) TestOrderWorkflow() {
	orderInput := orderv1.CreateOrderInput{
		Id:         "1234",
		CustomerId: "1234",
		Items: []*omsv1.Item{
			{Sku: "test1", Quantity: 1},
			{Sku: "test2", Quantity: 3},
		},
	}

	s.env.
		OnActivity(orderv1.CreateOrderActivityName, mock.Anything, mock.Anything).
		Return(&orderv1.CreateOrderResult{Order: &omsv1.Order{
			Id:         orderInput.Id,
			CustomerId: orderInput.CustomerId,
			ReceivedAt: timestamppb.Now(),
		}}, nil)

	s.env.
		OnWorkflow(billingv1.ChargeWorkflowName, mock.Anything, mock.Anything).
		Return(&billingv1.ChargeResult{Success: true}, nil)
	s.env.
		OnActivity(orderv1.UpdateOrderStatusActivityName, mock.Anything, mock.Anything).
		Return(nil)

	s.env.
		OnWorkflow(shipmentv1.ShipmentWorkflowName, mock.Anything, mock.Anything).
		Return(&shipmentv1.CreateShipmentResult{CourierReference: "test"}, nil).
		Times(2)

	_, err := s.orders.Order(context.Background(), &orderInput)
	s.Require().NoError(err)

	s.env.AssertWorkflowNumberOfCalls(s.T(), shipmentv1.ShipmentWorkflowName, 2)
}

func (s *WorkflowSuite) TestOrderShipmentStatus() {
	orderInput := orderv1.CreateOrderInput{
		Id:         "1234",
		CustomerId: "1234",
		Items: []*omsv1.Item{
			{Sku: "test1", Quantity: 1},
		},
	}

	order, err := s.orders.OrderAsync(s.ctx, &orderInput)
	s.Require().NoError(err)

	s.env.
		OnActivity(orderv1.CreateOrderActivityName, mock.Anything, mock.Anything).
		Return(&orderv1.CreateOrderResult{Order: &omsv1.Order{
			Id:         orderInput.Id,
			CustomerId: orderInput.CustomerId,
			ReceivedAt: timestamppb.Now(),
		}}, nil)

	s.env.
		OnWorkflow(billingv1.ChargeWorkflowName, mock.Anything, mock.Anything).
		Return(&billingv1.ChargeResult{Success: true}, nil)

	s.env.
		OnActivity(orderv1.UpdateOrderStatusActivityName, mock.Anything, mock.Anything).
		Return(nil)

	s.env.
		OnWorkflow(shipmentv1.ShipmentWorkflowName, mock.Anything, mock.Anything).
		Return(func(ctx workflow.Context, input *shipmentv1.CreateShipmentInput) (*shipmentv1.CreateShipmentResult, error) {
			order.ShipmentStatusUpdated(s.ctx, &shipmentv1.UpdateShipmentStatusInput{
				ShipmentId: input.GetId(),
				Status:     omsv1.ShipmentStatus_SHIPMENT_STATUS_DELIVERED,
				UpdatedAt:  timestamppb.Now(),
			})
			return &shipmentv1.CreateShipmentResult{CourierReference: "test"}, nil
		}).
		Times(2)

	_, err = order.Get(s.ctx)
	s.Require().NoError(err)

	status, err := s.orders.GetStatus(context.Background(), "", "")
	s.Require().NoError(err)

	f := status.GetOrder().GetFulfillments()[0]
	s.Require().Equal(omsv1.ShipmentStatus_SHIPMENT_STATUS_DELIVERED.String(), f.GetShipment().GetStatus().String())
}

func (s *WorkflowSuite) TestOrderAmendWithUnavailableItems() {
	orderInput := &orderv1.CreateOrderInput{
		Id:         "1234",
		CustomerId: "1234",
		Items: []*omsv1.Item{
			{Sku: "Adidas", Quantity: 1},
			{Sku: "test2", Quantity: 3},
		},
	}

	order, err := s.orders.OrderAsync(s.ctx, orderInput)
	s.Require().NoError(err)

	receivedAt := timestamppb.Now()
	s.env.
		OnActivity(orderv1.CreateOrderActivityName, mock.Anything, mock.Anything).
		Return(&orderv1.CreateOrderResult{Order: &omsv1.Order{
			Id:         orderInput.Id,
			CustomerId: orderInput.CustomerId,
			ReceivedAt: receivedAt,
		}}, nil)

	s.env.
		OnWorkflow(billingv1.ChargeWorkflowName, mock.Anything, mock.Anything).
		Return(&billingv1.ChargeResult{Success: true}, nil)

	s.env.
		OnActivity(orderv1.UpdateOrderStatusActivityName, mock.Anything, mock.Anything).
		Return(nil)

	s.env.
		OnWorkflow(shipmentv1.ShipmentWorkflowName, mock.Anything, mock.Anything).
		Return(&shipmentv1.CreateShipmentResult{CourierReference: "test"}, nil)

	s.env.RegisterDelayedCallback(func() {
		status, err := s.orders.GetStatus(s.ctx, "", "")
		s.Require().NoError(err)

		s.Require().Empty(gocmp.Diff(&omsv1.Order{
			Id:         "1234",
			CustomerId: "1234",
			Status:     omsv1.OrderStatus_ORDER_STATUS_CUSTOMER_ACTION_REQUIRED,
			ReceivedAt: receivedAt,
			Fulfillments: []*omsv1.Fulfillment{
				{
					CustomerId: orderInput.GetCustomerId(),
					OrderId:    orderInput.GetId(),
					Id:         "1234:1",
					Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_UNAVAILABLE,
					Items: []*omsv1.Item{
						{Sku: "Adidas", Quantity: 1},
					},
				},
				{
					CustomerId: orderInput.GetCustomerId(),
					OrderId:    orderInput.GetId(),
					Id:         "1234:2",
					Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_PENDING,
					Location:   "Warehouse A",
					Items: []*omsv1.Item{
						{Sku: "test2", Quantity: 3},
					},
				},
			},
		}, status.GetOrder(), protocmp.Transform()))
	}, time.Second*1)

	s.env.RegisterDelayedCallback(func() {
		order.CustomerAction(s.ctx, &orderv1.CustomerActionInput{
			Action: omsv1.CustomerAction_CUSTOMER_ACTION_AMEND,
		})
	}, time.Second*2)

	result, err := order.Get(s.ctx)
	s.Require().NoError(err)

	s.Require().Len(result.GetOrder().GetFulfillments(), 2)

	f := result.GetOrder().GetFulfillments()[0]
	s.Require().Equal(omsv1.FulfillmentStatus_FULFILLMENT_STATUS_CANCELLED.String(), f.GetStatus().String())

	f = result.GetOrder().GetFulfillments()[1]
	s.Require().Equal(omsv1.PaymentStatus_PAYMENT_STATUS_SUCCESS.String(), f.GetPayment().GetStatus().String())
	s.Require().Equal(f.GetId(), f.GetShipment().GetId())

	s.env.AssertWorkflowNumberOfCalls(s.T(), shipmentv1.ShipmentWorkflowName, 1)
}

func (s *WorkflowSuite) TestOrderCancelWithUnavailableItems() {
	orderInput := &orderv1.CreateOrderInput{
		Id:         "1234",
		CustomerId: "1234",
		Items: []*omsv1.Item{
			{Sku: "Adidas", Quantity: 1},
			{Sku: "test2", Quantity: 3},
		},
	}

	order, err := s.orders.OrderAsync(s.ctx, orderInput)
	s.Require().NoError(err)

	s.env.
		OnActivity(orderv1.CreateOrderActivityName, mock.Anything, mock.Anything).
		Return(&orderv1.CreateOrderResult{Order: &omsv1.Order{
			Id:         orderInput.Id,
			CustomerId: orderInput.CustomerId,
			ReceivedAt: timestamppb.Now(),
		}}, nil)

	s.env.OnActivity(orderv1.UpdateOrderStatusActivityName, mock.Anything, mock.Anything).Return(func(ctx context.Context, input *orderv1.UpdateOrderStatusInput) error {
		return nil
	})

	s.env.RegisterDelayedCallback(func() {
		order.CustomerAction(s.ctx, &orderv1.CustomerActionInput{
			Action: omsv1.CustomerAction_CUSTOMER_ACTION_CANCEL,
		})
	}, 1)

	_, err = order.Get(s.ctx)
	s.Require().NoError(err)
}

func (s *WorkflowSuite) TestOrderCancelAfterTimeout() {
	orderInput := orderv1.CreateOrderInput{
		Id:         "1234",
		CustomerId: "1234",
		Items: []*omsv1.Item{
			{Sku: "Adidas", Quantity: 1},
			{Sku: "test2", Quantity: 3},
		},
	}

	s.env.
		OnActivity(orderv1.CreateOrderActivityName, mock.Anything, mock.Anything).
		Return(&orderv1.CreateOrderResult{Order: &omsv1.Order{
			Id:         orderInput.Id,
			CustomerId: orderInput.CustomerId,
			ReceivedAt: timestamppb.Now(),
		}}, nil)

	s.env.
		OnActivity(orderv1.UpdateOrderStatusActivityName, mock.Anything, mock.Anything).
		Return(nil).
		Times(2)

	result, err := s.orders.Order(context.Background(), &orderInput)
	s.Require().NoError(err)

	s.Require().Equal(omsv1.OrderStatus_ORDER_STATUS_TIMED_OUT.String(), result.GetOrder().GetStatus().String())
}
