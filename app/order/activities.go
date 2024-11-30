package order

import (
	"context"
	"fmt"
	"strings"

	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// Activities implements the order package's Activities.
// Any state shared by the worker among the activities is stored here.
type Activities struct {
	BillingURL string
	OrderURL   string
	DB         *gorm.DB
}

// CreateOrder stores the Order status to the database.
func (a *Activities) CreateOrder(ctx context.Context, order *orderv1.CreateOrderInput) (*orderv1.CreateOrderResult, error) {
	m := &Order{
		ID:         order.GetId(),
		CustomerID: order.GetCustomerId(),
		Status:     omsv1.OrderStatus_ORDER_STATUS_PENDING.String(),
	}
	if err := a.DB.WithContext(ctx).Create(m).Error; err != nil {
		return nil, err
	}
	return &orderv1.CreateOrderResult{
		Order: &omsv1.Order{
			Id:         order.GetId(),
			CustomerId: order.GetCustomerId(),
			Status:     omsv1.OrderStatus_ORDER_STATUS_PENDING,
			ReceivedAt: timestamppb.New(m.ReceivedAt),
		},
	}, nil
}

// UpdateOrderStatus stores the Order status to the database.
func (a *Activities) UpdateOrderStatus(ctx context.Context, input *orderv1.UpdateOrderStatusInput) error {
	if err := a.DB.WithContext(ctx).Model(&Order{}).Where("id = ?", input.GetId()).Updates(&Order{
		Status: input.GetStatus().String(),
	}).Error; err != nil {
		return err
	}
	return nil
}

// ReserveItems reserves items to satisfy an order. It returns a list of reservations for the items.
// Any unavailable items will be returned in a Reservation with Available set to false.
// In a real system this would involve an inventory database of some kind.
// For our purposes we just split orders arbitrarily.
func (a *Activities) ReserveItems(_ context.Context, input *orderv1.ReserveItemsInput) (*orderv1.ReserveItemsResult, error) {
	if len(input.GetItems()) < 1 {
		return &orderv1.ReserveItemsResult{}, nil
	}

	var fulfillments []*omsv1.Fulfillment
	var unavailableItems []*omsv1.Item
	var availableItems []*omsv1.Item

	for _, item := range input.GetItems() {
		if strings.Contains(item.GetSku(), "Adidas") {
			unavailableItems = append(unavailableItems, item)
		} else {
			availableItems = append(availableItems, item)
		}
	}

	var fulfillmentId int

	if len(unavailableItems) > 0 {
		fulfillmentId++
		fulfillments = append(
			fulfillments,
			&omsv1.Fulfillment{
				OrderId:    input.GetOrderId(),
				CustomerId: input.GetCustomerId(),
				Id:         FulfillmentID(input.GetOrderId(), fulfillmentId),
				Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_UNAVAILABLE,
				Items:      unavailableItems,
			},
		)
	}

	// First item from one warehouse
	fulfillmentId++
	fulfillments = append(
		fulfillments,
		&omsv1.Fulfillment{
			OrderId:    input.GetOrderId(),
			CustomerId: input.GetCustomerId(),
			Id:         FulfillmentID(input.GetOrderId(), fulfillmentId),
			Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_PENDING,
			Location:   "Warehouse A",
			Items:      availableItems[0:1],
		},
	)

	if len(availableItems) > 1 {
		fulfillmentId++
		// Second fulfillment with all other items
		fulfillments = append(
			fulfillments,
			&omsv1.Fulfillment{
				OrderId:    input.GetOrderId(),
				CustomerId: input.GetCustomerId(),
				Id:         FulfillmentID(input.GetOrderId(), fulfillmentId),
				Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_PENDING,
				Location:   "Warehouse B",
				Items:      availableItems[1:],
			},
		)
	}

	return &orderv1.ReserveItemsResult{
		Fulfillments: fulfillments,
	}, nil
}

func FulfillmentID(orderID string, i int) string {
	return fmt.Sprintf("%s:%d", orderID, i)
}
