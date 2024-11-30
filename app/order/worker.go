package order

import (
	"context"

	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	"github.com/temporalio/reference-app-orders-go/internal/service"
	"github.com/temporalio/reference-app-orders-go/internal/temporalutil"
	"go.temporal.io/sdk/worker"
)

// RunWorker runs a Workflow and Activity worker for the Order system.
func RunWorker(ctx context.Context, params *service.RunParams) error {
	w := worker.New(params.Temporal, orderv1.WorkerTaskQueue, worker.Options{})

	orderv1.RegisterWorkerWorkflows(w, &Workflows{})
	orderv1.RegisterWorkerActivities(w, &Activities{BillingURL: params.Config.BillingURL, OrderURL: params.Config.OrderURL, DB: params.DB})

	return w.Run(temporalutil.WorkerInterruptFromContext(ctx))
}
