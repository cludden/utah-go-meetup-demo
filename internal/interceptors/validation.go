package interceptors

import (
	"context"
	"errors"

	"github.com/bufbuild/protovalidate-go"
	"github.com/temporalio/reference-app-orders-go/internal/temporalutil"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
)

type (
	validationInterceptor struct {
		interceptor.InterceptorBase
		v *protovalidate.Validator
	}

	validationInterceptorClientOutbound struct {
		interceptor.ClientOutboundInterceptorBase
		v *protovalidate.Validator
	}

	validationInterceptorActivityInbound struct {
		interceptor.ActivityInboundInterceptorBase
		v *protovalidate.Validator
	}

	validationInterceptorWorklowInbound struct {
		interceptor.WorkflowInboundInterceptorBase
		v *protovalidate.Validator
	}

	validationInterceptorWorklowOutbound struct {
		interceptor.WorkflowOutboundInterceptorBase
		v *protovalidate.Validator
	}
)

func NewValidation() interceptor.Interceptor {
	v, _ := protovalidate.New()
	return &validationInterceptor{v: v}
}

func (i *validationInterceptor) InterceptClient(next interceptor.ClientOutboundInterceptor) interceptor.ClientOutboundInterceptor {
	n := &validationInterceptorClientOutbound{v: i.v}
	n.Next = next
	return n
}

func (i *validationInterceptor) InterceptActivity(ctx context.Context, next interceptor.ActivityInboundInterceptor) interceptor.ActivityInboundInterceptor {
	n := &validationInterceptorActivityInbound{v: i.v}
	n.Next = next
	return n
}

func (i *validationInterceptor) InterceptWorkflow(ctx workflow.Context, next interceptor.WorkflowInboundInterceptor) interceptor.WorkflowInboundInterceptor {
	n := &validationInterceptorWorklowInbound{v: i.v}
	n.Next = next
	return n
}

// =============================================================================

func (i *validationInterceptorClientOutbound) ExecuteWorkflow(ctx context.Context, input *interceptor.ClientExecuteWorkflowInput) (client.WorkflowRun, error) {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return nil, err
	}
	return i.Next.ExecuteWorkflow(ctx, input)
}

func (i *validationInterceptorClientOutbound) QueryWorkflow(ctx context.Context, input *interceptor.ClientQueryWorkflowInput) (converter.EncodedValue, error) {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return nil, err
	}
	return i.Next.QueryWorkflow(ctx, input)
}

func (i *validationInterceptorClientOutbound) SignalWithStartWorkflow(ctx context.Context, input *interceptor.ClientSignalWithStartWorkflowInput) (client.WorkflowRun, error) {
	if err := validateArgs(i.v, append([]any{input.SignalArg}, input.Args...)); err != nil {
		return nil, err
	}
	return i.Next.SignalWithStartWorkflow(ctx, input)
}

func (i *validationInterceptorClientOutbound) SignalWorkflow(ctx context.Context, input *interceptor.ClientSignalWorkflowInput) error {
	if err := validateArgs(i.v, input.Arg); err != nil {
		return err
	}
	return i.Next.SignalWorkflow(ctx, input)
}

func (i *validationInterceptorClientOutbound) UpdateWorkflow(ctx context.Context, input *interceptor.ClientUpdateWorkflowInput) (client.WorkflowUpdateHandle, error) {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return nil, err
	}
	return i.Next.UpdateWorkflow(ctx, input)
}

// =============================================================================

func (i *validationInterceptorActivityInbound) ExecuteActivity(ctx context.Context, input *interceptor.ExecuteActivityInput) (any, error) {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return nil, err
	}
	return i.Next.ExecuteActivity(ctx, input)
}

// =============================================================================

func (i *validationInterceptorWorklowInbound) ExecuteUpdate(ctx workflow.Context, input *interceptor.UpdateInput) (any, error) {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return nil, err
	}
	return i.Next.ExecuteUpdate(ctx, input)
}

func (i *validationInterceptorWorklowInbound) ExecuteWorkflow(ctx workflow.Context, input *interceptor.ExecuteWorkflowInput) (any, error) {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return nil, err
	}
	return i.Next.ExecuteWorkflow(ctx, input)
}

func (i *validationInterceptorWorklowInbound) HandleQuery(ctx workflow.Context, input *interceptor.HandleQueryInput) (any, error) {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return nil, err
	}
	return i.Next.HandleQuery(ctx, input)
}

func (i *validationInterceptorWorklowInbound) HandleSignal(ctx workflow.Context, input *interceptor.HandleSignalInput) error {
	if err := validateArgs(i.v, input.Arg); err != nil {
		return err
	}
	return i.Next.HandleSignal(ctx, input)
}

func (i *validationInterceptorWorklowInbound) Init(next interceptor.WorkflowOutboundInterceptor) error {
	n := &validationInterceptorWorklowOutbound{v: i.v}
	n.Next = next
	return i.Next.Init(n)
}

func (i *validationInterceptorWorklowInbound) ValidateUpdate(ctx workflow.Context, input *interceptor.UpdateInput) error {
	if err := validateArgs(i.v, input.Args...); err != nil {
		return err
	}
	return i.Next.ValidateUpdate(ctx, input)
}

// =============================================================================

func (i *validationInterceptorWorklowOutbound) ExecuteActivity(ctx workflow.Context, activityType string, args ...any) workflow.Future {
	if err := validateArgs(i.v, args...); err != nil {
		return temporalutil.NewFutureFromError(ctx, err)
	}
	return i.Next.ExecuteActivity(ctx, activityType, args...)
}

func (i *validationInterceptorWorklowOutbound) ExecuteChildWorkflow(ctx workflow.Context, childWorkflowType string, args ...any) workflow.ChildWorkflowFuture {
	if err := validateArgs(i.v, args...); err != nil {
		return temporalutil.NewChildWorkflowFutureFromError(ctx, err)
	}
	return i.Next.ExecuteChildWorkflow(ctx, childWorkflowType, args...)
}

// =============================================================================

func validateArgs(v *protovalidate.Validator, args ...any) (err error) {
	for _, arg := range args {
		if pb, ok := arg.(proto.Message); ok && arg != nil {
			err = errors.Join(err, v.Validate(pb))
		}
	}
	return
}
