package order

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	"github.com/temporalio/reference-app-orders-go/gen/oms/order/v1/orderv1connect"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"github.com/temporalio/reference-app-orders-go/internal/service"
	"github.com/temporalio/reference-app-orders-go/internal/sliceutil"
	"go.temporal.io/api/serviceerror"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// OrderWorkflowId returns the workflow Id for an Order.
func OrderWorkflowId(id string) string {
	return "Order:" + id
}

type Handler struct {
	orderv1connect.UnimplementedApiHandler
	db        *gorm.DB
	logger    *slog.Logger
	workflows orderv1.WorkerClient
}

func RunServer(ctx context.Context, params *service.RunParams) error {
	if err := params.DB.AutoMigrate(&Order{}); err != nil {
		return fmt.Errorf("error initializing order db schema: %w", err)
	}
	return service.RunConnectServer(
		ctx,
		service.Must(params.Config.ServiceHostPort("order")),
		NewHandler(orderv1.NewWorkerClient(params.Temporal), params.DB, params.Logger),
		params.Logger,
	)
}

func NewHandler(workflows orderv1.WorkerClient, db *gorm.DB, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	path, handler := orderv1connect.NewApiHandler(&Handler{db: db, logger: logger, workflows: workflows})
	mux.Handle(path, handler)
	return mux
}

func (s *Handler) CreateOrder(ctx context.Context, req *connect.Request[orderv1.CreateOrderInput]) (*connect.Response[orderv1.CreateOrderResult], error) {
	order, err := s.workflows.OrderAsync(ctx, req.Msg)
	if err != nil {
		s.logger.Error("Failed to start order workflow", "error", err)
		return nil, err
	}

	status, err := order.GetStatus(ctx)
	if err != nil {
		s.logger.Error("Failed to query order workflow", "error", err)
		return nil, err
	}

	return connect.NewResponse(&orderv1.CreateOrderResult{Order: status.GetOrder()}), nil
}

func (s *Handler) CustomerAction(ctx context.Context, req *connect.Request[orderv1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error) {
	err := s.workflows.CustomerAction(ctx, OrderWorkflowId(req.Msg.GetId()), "", req.Msg)
	if err != nil {
		if _, ok := err.(*serviceerror.NotFound); ok {
			s.logger.Error("Failed to signal order workflow", "error", err)
			return nil, connect.NewError(connect.CodeNotFound, err)
		} else {
			s.logger.Error("Failed to signal order workflow", "error", err)
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *Handler) GetOrder(ctx context.Context, req *connect.Request[orderv1.GetOrderInput]) (*connect.Response[orderv1.GetOrderResult], error) {
	order, err := s.workflows.GetStatus(ctx, OrderWorkflowId(req.Msg.GetId()), "")
	if err != nil {
		if _, ok := err.(*serviceerror.NotFound); ok {
			s.logger.Error("Unable to query non-existent workflow", "error", err)
			return nil, connect.NewError(connect.CodeNotFound, err)
		} else {
			s.logger.Error("Failed to query order workflow", "error", err)
			return nil, err
		}
	}
	return connect.NewResponse(order), nil
}

func (s *Handler) ListOrders(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[orderv1.ListOrdersResult], error) {
	var orders []Order
	if err := s.db.WithContext(ctx).Find(&orders).Error; err != nil {
		s.logger.Error("Failed to list orders", "error", err)
		return nil, err
	}

	return connect.NewResponse(&orderv1.ListOrdersResult{
		Orders: sliceutil.Map(orders, func(o Order) *omsv1.Order {
			return &omsv1.Order{
				Id:         o.ID,
				CustomerId: o.CustomerID,
				Status:     omsv1.OrderStatus(omsv1.OrderStatus_value[o.Status]),
				ReceivedAt: timestamppb.New(o.ReceivedAt),
			}
		}),
	}), nil
}

func (s *Handler) UpdateOrderStatus(ctx context.Context, req *connect.Request[orderv1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse(&emptypb.Empty{}), s.workflows.UpdateOrderStatus(ctx, OrderWorkflowId(req.Msg.GetId()), "", req.Msg)
}
