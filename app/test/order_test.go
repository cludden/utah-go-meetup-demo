package test

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/temporalio/reference-app-orders-go/app/billing"
	"github.com/temporalio/reference-app-orders-go/app/fraud"
	"github.com/temporalio/reference-app-orders-go/app/order"
	"github.com/temporalio/reference-app-orders-go/app/shipment"
	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	"github.com/temporalio/reference-app-orders-go/gen/oms/order/v1/orderv1connect"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	"github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1/shipmentv1connect"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"github.com/temporalio/reference-app-orders-go/internal/config"
	"github.com/temporalio/reference-app-orders-go/internal/service"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/testsuite"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func Test_Order(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	s, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ClientOptions: &client.Options{
			Logger: log.NewStructuredLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
		},
		EnableUI:  false,
		ExtraArgs: []string{"--dynamic-config-value", "system.forceSearchAttributesCacheRefreshOnRead=true"},
		LogLevel:  "error",
	})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, s.Stop()) })

	c := s.Client()
	t.Cleanup(c.Close)

	_, err = c.OperatorService().AddSearchAttributes(ctx, &operatorservice.AddSearchAttributesRequest{
		Namespace: "default",
		SearchAttributes: map[string]enums.IndexedValueType{
			"CustomerId":    enums.INDEXED_VALUE_TYPE_KEYWORD,
			"FulfillmentId": enums.INDEXED_VALUE_TYPE_KEYWORD,
			"OrderId":       enums.INDEXED_VALUE_TYPE_KEYWORD,
		},
	})
	require.NoError(t, err)

	logger := slog.Default()

	fraudAPI := httptest.NewServer(fraud.Router(logger))
	t.Cleanup(fraudAPI.Close)
	billingAPI := httptest.NewServer(billing.NewHandler(billingv1.NewWorkerClient(c)))
	t.Cleanup(billingAPI.Close)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&order.Order{}, &shipment.Shipment{})
	require.NoError(t, err)

	orderAPI := httptest.NewServer(order.NewHandler(orderv1.NewWorkerClient(c), db, logger))
	t.Cleanup(orderAPI.Close)
	orderClient := orderv1connect.NewApiClient(&http.Client{}, orderAPI.URL)
	shipmentAPI := httptest.NewServer(shipment.NewHandler(shipmentv1.NewWorkerClient(c), db))
	t.Cleanup(shipmentAPI.Close)
	shipmentClient := shipmentv1connect.NewApiClient(&http.Client{}, shipmentAPI.URL)

	config := config.AppConfig{
		BillingURL:  billingAPI.URL,
		OrderURL:    orderAPI.URL,
		ShipmentURL: shipmentAPI.URL,
		FraudURL:    fraudAPI.URL,
	}

	g, ctx := errgroup.WithContext(ctx)

	params := &service.RunParams{
		Config:   config,
		DB:       db,
		Temporal: c,
	}

	g.Go(func() error {
		return billing.RunWorker(ctx, params)
	})
	g.Go(func() error {
		return shipment.RunWorker(ctx, params)
	})
	g.Go(func() error {
		return order.RunWorker(ctx, params)
	})

	_, err = orderClient.CreateOrder(ctx, connect.NewRequest(&orderv1.CreateOrderInput{
		Id:         "order123",
		CustomerId: "customer123",
		Items: []*omsv1.Item{
			{Sku: "Adidas Classic", Quantity: 1},
			{Sku: "Nike Air", Quantity: 2},
		},
	}))
	require.NoError(t, err)

	require.EventuallyWithT(t, func(c *assert.CollectT) {
		o, err := orderClient.GetOrder(ctx, connect.NewRequest(&orderv1.GetOrderInput{
			Id: "order123",
		}))
		require.NoError(t, err)

		assert.Equal(c, omsv1.OrderStatus_ORDER_STATUS_CUSTOMER_ACTION_REQUIRED.String(), o.Msg.GetOrder().GetStatus().String())
	}, 10*time.Second, 100*time.Millisecond)

	_, err = orderClient.CustomerAction(ctx, connect.NewRequest(&orderv1.CustomerActionInput{
		Id:     "order123",
		Action: omsv1.CustomerAction_CUSTOMER_ACTION_AMEND,
	}))
	require.NoError(t, err)

	require.EventuallyWithT(t, func(c *assert.CollectT) {
		o, err := orderClient.GetOrder(ctx, connect.NewRequest(&orderv1.GetOrderInput{
			Id: "order123",
		}))
		require.NoError(t, err)

		assert.Equal(c, omsv1.OrderStatus_ORDER_STATUS_PROCESSING, o.Msg.GetOrder().GetStatus())
		assert.Len(c, o.Msg.GetOrder().GetFulfillments(), 2)
		assert.NotEmpty(c, o.Msg.GetOrder().GetFulfillments()[1].GetPayment())
		assert.NotEmpty(c, o.Msg.GetOrder().GetFulfillments()[1].GetShipment())
		assert.NotEmpty(c, o.Msg.GetOrder().GetFulfillments()[1].GetShipment().GetId())
	}, 10*time.Second, 100*time.Millisecond)

	o, err := orderClient.GetOrder(ctx, connect.NewRequest(&orderv1.GetOrderInput{
		Id: "order123",
	}))
	require.NoError(t, err)

	time.Sleep(time.Second * 3)
	for _, f := range o.Msg.GetOrder().GetFulfillments() {
		if f.GetShipment() == nil {
			continue
		}

		_, err := shipmentClient.UpdateShipmentStatus(ctx, connect.NewRequest(&shipmentv1.UpdateShipmentStatusInput{
			ShipmentId: f.GetShipment().GetId(),
			Status:     omsv1.ShipmentStatus_SHIPMENT_STATUS_DELIVERED,
		}))
		require.NoError(t, err)
	}

	require.EventuallyWithT(t, func(c *assert.CollectT) {
		o, err := orderClient.GetOrder(context.Background(), connect.NewRequest(&orderv1.GetOrderInput{
			Id: "order123",
		}))
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(c, omsv1.OrderStatus_ORDER_STATUS_COMPLETED.String(), o.Msg.GetOrder().GetStatus().String())
	}, 3*time.Second, 100*time.Millisecond)

	cancel()

	err = g.Wait()
	require.NoError(t, err)
}
