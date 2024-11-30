package shipment_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/temporalio/reference-app-orders-go/app/shipment"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"github.com/temporalio/reference-app-orders-go/internal/interceptors"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

type WorkflowSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env       *testsuite.TestWorkflowEnvironment
	shipments *shipmentv1.TestWorkerClient
}

func TestWorkflowSuite(t *testing.T) {
	suite.Run(t, new(WorkflowSuite))
}

func (s *WorkflowSuite) SetupTest() {
	s.T().Setenv("TEMPORAL_DEBUG", "true")
	s.env = s.NewTestWorkflowEnvironment()
	s.env.SetWorkerOptions(worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{interceptors.NewValidation()},
	})
	s.shipments = shipmentv1.NewTestWorkerClient(s.env, &shipment.Workflows{}, &shipment.Activities{})
}

func (s *WorkflowSuite) TestShipmentWorkflow() {
	r, ctx := s.Require(), context.Background()

	shipmentInput := &shipmentv1.CreateShipmentInput{
		RequestorWid:  "parentwid",
		CustomerId:    "abc",
		OrderId:       "123",
		FulfillmentId: "123:1",
		Id:            "test",
		Items: []*omsv1.Item{
			{Sku: "test1", Quantity: 1},
			{Sku: "test2", Quantity: 3},
		},
	}

	shipment, err := s.shipments.ShipmentAsync(ctx, shipmentInput)
	r.NoError(err)

	s.env.
		OnActivity(shipmentv1.CreateShipmentActivityName, mock.Anything, mock.Anything).
		Return(&shipmentv1.CreateShipmentResult{CourierReference: "test"}, nil)

	s.env.
		OnActivity(shipmentv1.UpdateShipmentStatusActivityName, mock.Anything, mock.Anything).
		Return(nil)

	var update1 shipmentv1.UpdateShipmentStatusHandle
	s.env.RegisterDelayedCallback(func() {
		update1, err = shipment.UpdateShipmentStatusAsync(ctx, &shipmentv1.UpdateShipmentStatusInput{
			ShipmentId: shipmentInput.Id,
			Status:     omsv1.ShipmentStatus_SHIPMENT_STATUS_DELIVERED,
		})
		r.NoError(err)
	}, time.Hour)

	s.env.OnSignalExternalWorkflow(mock.Anything,
		"parentwid", "",
		shipmentv1.ShipmentStatusUpdatedSignalName,
		mock.MatchedBy(func(arg *shipmentv1.UpdateShipmentStatusInput) bool {
			return arg.Status == omsv1.ShipmentStatus_SHIPMENT_STATUS_BOOKED
		}),
	).Return(nil).Once()

	s.env.OnSignalExternalWorkflow(mock.Anything,
		"parentwid", "",
		shipmentv1.ShipmentStatusUpdatedSignalName,
		mock.MatchedBy(func(arg *shipmentv1.UpdateShipmentStatusInput) bool {
			return arg.Status == omsv1.ShipmentStatus_SHIPMENT_STATUS_DELIVERED
		}),
	).Return(nil).Once()

	_, err = shipment.Get(ctx)
	r.NoError(err)
	r.NoError(update1.Get(ctx))
}
