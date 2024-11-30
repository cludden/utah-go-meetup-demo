package shipment

import (
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Workflows struct{}

type ShipmentWorkflow struct {
	*shipmentv1.ShipmentWorkflowInput
	logger    log.Logger
	status    omsv1.ShipmentStatus
	updatedAt *timestamppb.Timestamp
	m         workflow.Mutex
}

// Shipment implements the Shipment workflow.
func (w *Workflows) Shipment(ctx workflow.Context, input *shipmentv1.ShipmentWorkflowInput) (shipmentv1.ShipmentWorkflow, error) {
	return &ShipmentWorkflow{input, log.With(workflow.GetLogger(ctx), "shipmentId", input.Req.GetId()), omsv1.ShipmentStatus_SHIPMENT_STATUS_PENDING, nil, workflow.NewMutex(ctx)}, nil
}

func (wf *ShipmentWorkflow) Execute(ctx workflow.Context) (*shipmentv1.CreateShipmentResult, error) {
	result, err := shipmentv1.CreateShipment(ctx, wf.Req)
	if err != nil {
		return nil, err
	}

	if err := wf.UpdateShipmentStatus(ctx, &shipmentv1.UpdateShipmentStatusInput{
		Status:    omsv1.ShipmentStatus_SHIPMENT_STATUS_BOOKED,
		UpdatedAt: timestamppb.New(workflow.Now(ctx)),
	}); err != nil {
		wf.logger.Error("failure updating shipment status", "error", err)
	}

	if err := workflow.Await(ctx, func() bool {
		return wf.status == omsv1.ShipmentStatus_SHIPMENT_STATUS_DELIVERED && workflow.AllHandlersFinished(ctx)
	}); err != nil {
		return nil, err
	}

	return &shipmentv1.CreateShipmentResult{
		CourierReference: result.CourierReference,
	}, err
}

func (s *ShipmentWorkflow) GetStatus() (*shipmentv1.GetShipmentResult, error) {
	return &shipmentv1.GetShipmentResult{
		Shipment: &omsv1.Shipment{
			Id:        s.Req.GetId(),
			Status:    s.status,
			UpdatedAt: s.updatedAt,
			Items:     s.Req.GetItems(),
		},
	}, nil
}

func (s *ShipmentWorkflow) UpdateShipmentStatus(ctx workflow.Context, input *shipmentv1.UpdateShipmentStatusInput) error {
	if err := s.m.Lock(ctx); err != nil {
		return err
	}
	defer s.m.Unlock()

	s.logger.Info("received shipment status update", "status", input.GetStatus().String())
	input.ShipmentId = s.Req.GetId()
	s.status = input.GetStatus()
	if input.GetUpdatedAt().IsValid() {
		s.updatedAt = input.GetUpdatedAt()
	} else {
		s.updatedAt = timestamppb.New(workflow.Now(ctx))
	}

	if err := shipmentv1.UpdateShipmentStatusLocal(ctx, &shipmentv1.UpdateShipmentStatusInput{
		ShipmentId: s.Req.GetId(),
		Status:     s.status,
		UpdatedAt:  s.updatedAt,
	}); err != nil {
		return err
	}

	return shipmentv1.ShipmentStatusUpdatedExternal(ctx, s.Req.GetRequestorWid(), "", input)
}
