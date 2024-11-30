package temporalutil

import "go.temporal.io/sdk/workflow"

type childWorkflowErrorFuture struct {
	f   workflow.Future
	err error
}

func NewChildWorkflowFutureFromError(ctx workflow.Context, err error) workflow.ChildWorkflowFuture {
	return &childWorkflowErrorFuture{NewFutureFromError(ctx, err), err}
}

func (f *childWorkflowErrorFuture) Get(workflow.Context, any) error { return f.err }

func (f *childWorkflowErrorFuture) GetChildWorkflowExecution() workflow.Future {
	return f.f
}

func (f *childWorkflowErrorFuture) IsReady() bool { return true }

func (f *childWorkflowErrorFuture) SignalChildWorkflow(workflow.Context, string, any) workflow.Future {
	return f.f
}

func NewFutureFromError(ctx workflow.Context, err error) workflow.Future {
	f, s := workflow.NewFuture(ctx)
	s.SetError(err)
	return f
}
