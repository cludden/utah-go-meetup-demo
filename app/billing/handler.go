package billing

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	"github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1/billingv1connect"
	"github.com/temporalio/reference-app-orders-go/internal/service"
)

type Handler struct {
	workflows billingv1.WorkerClient
}

func RunServer(ctx context.Context, params *service.RunParams) error {
	return service.RunConnectServer(
		ctx,
		service.Must(params.Config.ServiceHostPort("billing")),
		NewHandler(billingv1.NewWorkerClient(params.Temporal)),
		params.Logger,
	)
}

func NewHandler(workflows billingv1.WorkerClient) http.Handler {
	mux := http.NewServeMux()
	path, handler := billingv1connect.NewApiHandler(&Handler{workflows})
	mux.Handle(path, handler)
	return mux
}

func (h *Handler) Charge(ctx context.Context, req *connect.Request[billingv1.ChargeInput]) (*connect.Response[billingv1.ChargeResult], error) {
	result, err := h.workflows.Charge(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(result), nil
}
