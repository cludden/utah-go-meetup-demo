package shipment_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/temporalio/reference-app-orders-go/app/shipment"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	"github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1/shipmentv1connect"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	shipmentv1mocks "github.com/temporalio/reference-app-orders-go/mocks/gen/oms/shipment/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestShipmentUpdate(t *testing.T) {
	c := shipmentv1mocks.NewMockWorkerClient(t)
	c.EXPECT().
		UpdateShipmentStatus(mock.Anything, shipment.ShipmentWorkflowID("test"), "", mock.MatchedBy(func(req *shipmentv1.UpdateShipmentStatusInput) bool {
			return assert.Empty(t, cmp.Diff(&shipmentv1.UpdateShipmentStatusInput{
				ShipmentId: "test",
				Status:     omsv1.ShipmentStatus_SHIPMENT_STATUS_DISPATCHED,
			}, req, protocmp.Transform()))
		})).
		Return(nil)
	h := shipment.NewHandler(c, nil)

	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)

	cc := shipmentv1connect.NewApiClient(&http.Client{}, srv.URL)
	_, err := cc.UpdateShipmentStatus(context.Background(), connect.NewRequest(&shipmentv1.UpdateShipmentStatusInput{
		ShipmentId: "test",
		Status:     omsv1.ShipmentStatus_SHIPMENT_STATUS_DISPATCHED,
	}))
	require.NoError(t, err)
}
