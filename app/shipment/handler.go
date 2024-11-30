package shipment

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	"github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1/shipmentv1connect"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"github.com/temporalio/reference-app-orders-go/internal/service"
	"github.com/temporalio/reference-app-orders-go/internal/sliceutil"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Handler struct {
	db        *gorm.DB
	workflows shipmentv1.WorkerClient
}

func RunServer(ctx context.Context, params *service.RunParams) error {
	if err := params.DB.AutoMigrate(&Shipment{}); err != nil {
		return fmt.Errorf("error initializing shipment db schema: %w", err)
	}
	return service.RunConnectServer(
		ctx,
		service.Must(params.Config.ServiceHostPort("shipment")),
		NewHandler(shipmentv1.NewWorkerClient(params.Temporal), params.DB),
		params.Logger,
	)
}

func NewHandler(workflows shipmentv1.WorkerClient, db *gorm.DB) http.Handler {
	mux := http.NewServeMux()
	path, handler := shipmentv1connect.NewApiHandler(&Handler{db, workflows})
	mux.Handle(path, handler)
	return mux
}

func (h *Handler) GetShipment(ctx context.Context, req *connect.Request[shipmentv1.GetShipmentInput]) (*connect.Response[shipmentv1.GetShipmentResult], error) {
	shipment, err := h.workflows.GetStatus(ctx, ShipmentWorkflowID(req.Msg.GetId()), "")
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(shipment), nil
}

func (h *Handler) ListShipments(ctx context.Context, req *connect.Request[shipmentv1.ListShipmentsInput]) (*connect.Response[shipmentv1.ListShipmentsResult], error) {
	var shipments []Shipment
	if err := h.db.WithContext(ctx).Find(&shipments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}
	return connect.NewResponse(&shipmentv1.ListShipmentsResult{
		Shipments: sliceutil.Map(shipments, func(s Shipment) *omsv1.Shipment {
			return &omsv1.Shipment{
				Id:        s.ID,
				Status:    omsv1.ShipmentStatus(omsv1.ShipmentStatus_value[s.Status]),
				UpdatedAt: timestamppb.New(s.BookedAt),
			}
		}),
	}), nil
}

func (h *Handler) UpdateShipmentStatus(ctx context.Context, req *connect.Request[shipmentv1.UpdateShipmentStatusInput]) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse(&emptypb.Empty{}), h.workflows.UpdateShipmentStatus(ctx, ShipmentWorkflowID(req.Msg.GetShipmentId()), "", req.Msg)
}
