package shipment

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"connectrpc.com/connect"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"go.temporal.io/sdk/temporal"
	"gorm.io/gorm"
)

// Activities implements the shipment package's Activities.
// Any state shared by the worker among the activities is stored here.
type Activities struct {
	ShipmentURL string
	DB          *gorm.DB
}

// CreateShipment engages a courier who can deliver the shipment to the customer
func (a *Activities) CreateShipment(ctx context.Context, input *shipmentv1.CreateShipmentInput) (*shipmentv1.CreateShipmentResult, error) {
	if v := os.Getenv("UNKNOWN_ITEMS"); v != "" {
		panicbrands := strings.Split(v, ",")
		for _, item := range input.GetItems() {
			for _, brand := range panicbrands {
				if strings.Contains(item.GetSku(), brand) {
					panic(fmt.Errorf("unknown brand %s", brand))
				}
			}
		}
	}

	m := &Shipment{
		ID:     input.GetId(),
		Status: omsv1.ShipmentStatus_SHIPMENT_STATUS_PENDING.String(),
	}

	if err := a.DB.WithContext(ctx).Create(m).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, temporal.NewNonRetryableApplicationError("failure booking shipment", connect.CodeAlreadyExists.String(), err)
		}
		return nil, err
	}

	return &shipmentv1.CreateShipmentResult{
		CourierReference: input.GetId() + ":1234",
	}, nil
}

// UpdateShipmentStatus stores the Order status to the database.
func (a *Activities) UpdateShipmentStatus(ctx context.Context, input *shipmentv1.UpdateShipmentStatusInput) error {
	if err := a.DB.WithContext(ctx).Model(&Shipment{}).Where("id = ?", input.GetShipmentId()).Updates(&Shipment{
		Status: input.GetStatus().String(),
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return temporal.NewNonRetryableApplicationError("failure updating shipment status", connect.CodeNotFound.String(), err)
		}
		return err
	}
	return nil
}
