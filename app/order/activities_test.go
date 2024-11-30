package order_test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"
	"github.com/temporalio/reference-app-orders-go/app/order"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"github.com/temporalio/reference-app-orders-go/internal/interceptors"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"google.golang.org/protobuf/testing/protocmp"
	"gorm.io/gorm"
)

type ActivitiesSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env        *testsuite.TestActivityEnvironment
	activities *order.Activities
}

func TestActivities(t *testing.T) {
	suite.Run(t, new(ActivitiesSuite))
}

func (s *ActivitiesSuite) SetupTest() {
	s.T().Setenv("TEMPORAL_DEBUG", "true")
	s.env = s.NewTestActivityEnvironment()
	s.env.SetWorkerOptions(worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{interceptors.NewValidation()},
	})
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.Require().NoError(err)
	s.activities = &order.Activities{DB: db}
	orderv1.RegisterWorkerActivities(s.env, s.activities)
}

func (s *ActivitiesSuite) TestFulfillOrderZeroItems() {
	input := orderv1.ReserveItemsInput{
		OrderId:    "1234",
		CustomerId: "abc",
		Items:      []*omsv1.Item{},
	}

	future, err := s.env.ExecuteActivity(orderv1.ReserveItemsActivityName, &input)
	s.Require().NoError(err)

	var result orderv1.ReserveItemsResult
	s.Require().NoError(future.Get(&result))

	expected := orderv1.ReserveItemsResult{}

	s.Require().Empty(cmp.Diff(&expected, &result, protocmp.Transform()))
}

func (s *ActivitiesSuite) TestFulfillOrderOneItem() {
	input := orderv1.ReserveItemsInput{
		OrderId:    "test",
		CustomerId: "abc",
		Items: []*omsv1.Item{
			{Sku: "Hiking Boots", Quantity: 2},
		},
	}

	future, err := s.env.ExecuteActivity(orderv1.ReserveItemsActivityName, &input)
	s.Require().NoError(err)

	var result orderv1.ReserveItemsResult
	s.Require().NoError(future.Get(&result))

	expected := orderv1.ReserveItemsResult{
		Fulfillments: []*omsv1.Fulfillment{
			{
				OrderId:    input.GetOrderId(),
				CustomerId: "abc",
				Id:         order.FulfillmentID(input.GetOrderId(), 1),
				Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_PENDING,
				Location:   "Warehouse A",
				Items: []*omsv1.Item{
					{Sku: "Hiking Boots", Quantity: 2},
				},
			},
		},
	}

	s.Require().Empty(cmp.Diff(&expected, &result, protocmp.Transform()))
}

func (s *ActivitiesSuite) TestFulfillOrderTwoItems() {
	input := orderv1.ReserveItemsInput{
		OrderId:    "test",
		CustomerId: "abc",
		Items: []*omsv1.Item{
			{Sku: "Hiking Boots", Quantity: 2},
			{Sku: "Tennis Shoes", Quantity: 1},
		},
	}

	future, err := s.env.ExecuteActivity(orderv1.ReserveItemsActivityName, &input)
	s.Require().NoError(err)

	var result orderv1.ReserveItemsResult
	s.Require().NoError(future.Get(&result))

	expected := orderv1.ReserveItemsResult{
		Fulfillments: []*omsv1.Fulfillment{
			{
				Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_PENDING,
				OrderId:    input.GetOrderId(),
				CustomerId: "abc",
				Id:         order.FulfillmentID(input.GetOrderId(), 1),
				Location:   "Warehouse A",
				Items: []*omsv1.Item{
					{Sku: "Hiking Boots", Quantity: 2},
				},
			},
			{
				Status:     omsv1.FulfillmentStatus_FULFILLMENT_STATUS_PENDING,
				OrderId:    input.GetOrderId(),
				CustomerId: "abc",
				Id:         order.FulfillmentID(input.GetOrderId(), 2),
				Location:   "Warehouse B",
				Items: []*omsv1.Item{
					{Sku: "Tennis Shoes", Quantity: 1},
				},
			},
		},
	}

	s.Require().Empty(cmp.Diff(&expected, &result, protocmp.Transform()))
}
