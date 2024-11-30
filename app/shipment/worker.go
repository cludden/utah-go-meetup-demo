package shipment

import (
	"context"

	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	"github.com/temporalio/reference-app-orders-go/internal/service"
	"github.com/temporalio/reference-app-orders-go/internal/temporalutil"
	"go.temporal.io/sdk/worker"
)

// Config is the configuration for the Order system.
type Config struct {
	ShipmentURL string
}

// RunWorker runs a Workflow and Activity worker for the Shipment system.
func RunWorker(ctx context.Context, params *service.RunParams) error {
	w := worker.New(params.Temporal, shipmentv1.WorkerTaskQueue, worker.Options{})

	shipmentv1.RegisterWorkerWorkflows(w, &Workflows{})
	shipmentv1.RegisterWorkerActivities(w, &Activities{ShipmentURL: params.Config.ShipmentURL, DB: params.DB})

	return w.Run(temporalutil.WorkerInterruptFromContext(ctx))
}
