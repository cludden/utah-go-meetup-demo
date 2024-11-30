package billing

import (
	"context"

	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	"github.com/temporalio/reference-app-orders-go/internal/service"
	"github.com/temporalio/reference-app-orders-go/internal/temporalutil"
	"go.temporal.io/sdk/worker"
)

// RunWorker runs a Workflow and Activity worker for the Billing system.
func RunWorker(ctx context.Context, params *service.RunParams) error {
	w := worker.New(params.Temporal, billingv1.WorkerTaskQueue, worker.Options{})

	billingv1.RegisterWorkerWorkflows(w, &Workflows{})
	billingv1.RegisterWorkerActivities(w, &Activities{FraudCheckURL: params.Config.FraudURL})

	return w.Run(temporalutil.WorkerInterruptFromContext(ctx))
}
