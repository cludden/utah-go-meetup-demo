package shipment

import (
	"strings"
	"time"
)

// ShipmentWorkflowID returns the workflow ID for a Shipment.
func ShipmentWorkflowID(id string) string {
	return "Shipment:" + id
}

// ShipmentIDFromWorkflowID returns the ID for a Shipment from a WorkflowID.
func ShipmentIDFromWorkflowID(id string) string {
	return strings.TrimPrefix(id, "Shipment:")
}

type Shipment struct {
	ID       string `gorm:"primaryKey"`
	Status   string
	BookedAt time.Time `gorm:"autoCreateTime"`
}
