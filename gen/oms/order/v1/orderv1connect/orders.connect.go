// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: oms/order/v1/orders.proto

package orderv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ApiName is the fully-qualified name of the Api service.
	ApiName = "oms.order.v1.Api"
	// WorkerName is the fully-qualified name of the Worker service.
	WorkerName = "oms.order.v1.Worker"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ApiCreateOrderProcedure is the fully-qualified name of the Api's CreateOrder RPC.
	ApiCreateOrderProcedure = "/oms.order.v1.Api/CreateOrder"
	// ApiGetOrderProcedure is the fully-qualified name of the Api's GetOrder RPC.
	ApiGetOrderProcedure = "/oms.order.v1.Api/GetOrder"
	// ApiListOrdersProcedure is the fully-qualified name of the Api's ListOrders RPC.
	ApiListOrdersProcedure = "/oms.order.v1.Api/ListOrders"
	// ApiUpdateOrderStatusProcedure is the fully-qualified name of the Api's UpdateOrderStatus RPC.
	ApiUpdateOrderStatusProcedure = "/oms.order.v1.Api/UpdateOrderStatus"
	// ApiCustomerActionProcedure is the fully-qualified name of the Api's CustomerAction RPC.
	ApiCustomerActionProcedure = "/oms.order.v1.Api/CustomerAction"
	// WorkerCreateOrderProcedure is the fully-qualified name of the Worker's CreateOrder RPC.
	WorkerCreateOrderProcedure = "/oms.order.v1.Worker/CreateOrder"
	// WorkerCustomerActionProcedure is the fully-qualified name of the Worker's CustomerAction RPC.
	WorkerCustomerActionProcedure = "/oms.order.v1.Worker/CustomerAction"
	// WorkerGetStatusProcedure is the fully-qualified name of the Worker's GetStatus RPC.
	WorkerGetStatusProcedure = "/oms.order.v1.Worker/GetStatus"
	// WorkerOrderProcedure is the fully-qualified name of the Worker's Order RPC.
	WorkerOrderProcedure = "/oms.order.v1.Worker/Order"
	// WorkerReserveItemsProcedure is the fully-qualified name of the Worker's ReserveItems RPC.
	WorkerReserveItemsProcedure = "/oms.order.v1.Worker/ReserveItems"
	// WorkerUpdateOrderStatusProcedure is the fully-qualified name of the Worker's UpdateOrderStatus
	// RPC.
	WorkerUpdateOrderStatusProcedure = "/oms.order.v1.Worker/UpdateOrderStatus"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	apiServiceDescriptor                    = v1.File_oms_order_v1_orders_proto.Services().ByName("Api")
	apiCreateOrderMethodDescriptor          = apiServiceDescriptor.Methods().ByName("CreateOrder")
	apiGetOrderMethodDescriptor             = apiServiceDescriptor.Methods().ByName("GetOrder")
	apiListOrdersMethodDescriptor           = apiServiceDescriptor.Methods().ByName("ListOrders")
	apiUpdateOrderStatusMethodDescriptor    = apiServiceDescriptor.Methods().ByName("UpdateOrderStatus")
	apiCustomerActionMethodDescriptor       = apiServiceDescriptor.Methods().ByName("CustomerAction")
	workerServiceDescriptor                 = v1.File_oms_order_v1_orders_proto.Services().ByName("Worker")
	workerCreateOrderMethodDescriptor       = workerServiceDescriptor.Methods().ByName("CreateOrder")
	workerCustomerActionMethodDescriptor    = workerServiceDescriptor.Methods().ByName("CustomerAction")
	workerGetStatusMethodDescriptor         = workerServiceDescriptor.Methods().ByName("GetStatus")
	workerOrderMethodDescriptor             = workerServiceDescriptor.Methods().ByName("Order")
	workerReserveItemsMethodDescriptor      = workerServiceDescriptor.Methods().ByName("ReserveItems")
	workerUpdateOrderStatusMethodDescriptor = workerServiceDescriptor.Methods().ByName("UpdateOrderStatus")
)

// ApiClient is a client for the oms.order.v1.Api service.
type ApiClient interface {
	// submit a new order
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error)
	// returns information about an existing order
	GetOrder(context.Context, *connect.Request[v1.GetOrderInput]) (*connect.Response[v1.GetOrderResult], error)
	// returns a list of existing orders
	ListOrders(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListOrdersResult], error)
	// updates the status of an exsiting order
	UpdateOrderStatus(context.Context, *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error)
	// process a customer action
	CustomerAction(context.Context, *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error)
}

// NewApiClient constructs a client for the oms.order.v1.Api service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewApiClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ApiClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &apiClient{
		createOrder: connect.NewClient[v1.CreateOrderInput, v1.CreateOrderResult](
			httpClient,
			baseURL+ApiCreateOrderProcedure,
			connect.WithSchema(apiCreateOrderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOrder: connect.NewClient[v1.GetOrderInput, v1.GetOrderResult](
			httpClient,
			baseURL+ApiGetOrderProcedure,
			connect.WithSchema(apiGetOrderMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listOrders: connect.NewClient[emptypb.Empty, v1.ListOrdersResult](
			httpClient,
			baseURL+ApiListOrdersProcedure,
			connect.WithSchema(apiListOrdersMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		updateOrderStatus: connect.NewClient[v1.UpdateOrderStatusInput, emptypb.Empty](
			httpClient,
			baseURL+ApiUpdateOrderStatusProcedure,
			connect.WithSchema(apiUpdateOrderStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		customerAction: connect.NewClient[v1.CustomerActionInput, emptypb.Empty](
			httpClient,
			baseURL+ApiCustomerActionProcedure,
			connect.WithSchema(apiCustomerActionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// apiClient implements ApiClient.
type apiClient struct {
	createOrder       *connect.Client[v1.CreateOrderInput, v1.CreateOrderResult]
	getOrder          *connect.Client[v1.GetOrderInput, v1.GetOrderResult]
	listOrders        *connect.Client[emptypb.Empty, v1.ListOrdersResult]
	updateOrderStatus *connect.Client[v1.UpdateOrderStatusInput, emptypb.Empty]
	customerAction    *connect.Client[v1.CustomerActionInput, emptypb.Empty]
}

// CreateOrder calls oms.order.v1.Api.CreateOrder.
func (c *apiClient) CreateOrder(ctx context.Context, req *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error) {
	return c.createOrder.CallUnary(ctx, req)
}

// GetOrder calls oms.order.v1.Api.GetOrder.
func (c *apiClient) GetOrder(ctx context.Context, req *connect.Request[v1.GetOrderInput]) (*connect.Response[v1.GetOrderResult], error) {
	return c.getOrder.CallUnary(ctx, req)
}

// ListOrders calls oms.order.v1.Api.ListOrders.
func (c *apiClient) ListOrders(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListOrdersResult], error) {
	return c.listOrders.CallUnary(ctx, req)
}

// UpdateOrderStatus calls oms.order.v1.Api.UpdateOrderStatus.
func (c *apiClient) UpdateOrderStatus(ctx context.Context, req *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error) {
	return c.updateOrderStatus.CallUnary(ctx, req)
}

// CustomerAction calls oms.order.v1.Api.CustomerAction.
func (c *apiClient) CustomerAction(ctx context.Context, req *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error) {
	return c.customerAction.CallUnary(ctx, req)
}

// ApiHandler is an implementation of the oms.order.v1.Api service.
type ApiHandler interface {
	// submit a new order
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error)
	// returns information about an existing order
	GetOrder(context.Context, *connect.Request[v1.GetOrderInput]) (*connect.Response[v1.GetOrderResult], error)
	// returns a list of existing orders
	ListOrders(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListOrdersResult], error)
	// updates the status of an exsiting order
	UpdateOrderStatus(context.Context, *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error)
	// process a customer action
	CustomerAction(context.Context, *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error)
}

// NewApiHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewApiHandler(svc ApiHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	apiCreateOrderHandler := connect.NewUnaryHandler(
		ApiCreateOrderProcedure,
		svc.CreateOrder,
		connect.WithSchema(apiCreateOrderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiGetOrderHandler := connect.NewUnaryHandler(
		ApiGetOrderProcedure,
		svc.GetOrder,
		connect.WithSchema(apiGetOrderMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	apiListOrdersHandler := connect.NewUnaryHandler(
		ApiListOrdersProcedure,
		svc.ListOrders,
		connect.WithSchema(apiListOrdersMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	apiUpdateOrderStatusHandler := connect.NewUnaryHandler(
		ApiUpdateOrderStatusProcedure,
		svc.UpdateOrderStatus,
		connect.WithSchema(apiUpdateOrderStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiCustomerActionHandler := connect.NewUnaryHandler(
		ApiCustomerActionProcedure,
		svc.CustomerAction,
		connect.WithSchema(apiCustomerActionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/oms.order.v1.Api/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ApiCreateOrderProcedure:
			apiCreateOrderHandler.ServeHTTP(w, r)
		case ApiGetOrderProcedure:
			apiGetOrderHandler.ServeHTTP(w, r)
		case ApiListOrdersProcedure:
			apiListOrdersHandler.ServeHTTP(w, r)
		case ApiUpdateOrderStatusProcedure:
			apiUpdateOrderStatusHandler.ServeHTTP(w, r)
		case ApiCustomerActionProcedure:
			apiCustomerActionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedApiHandler returns CodeUnimplemented from all methods.
type UnimplementedApiHandler struct{}

func (UnimplementedApiHandler) CreateOrder(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Api.CreateOrder is not implemented"))
}

func (UnimplementedApiHandler) GetOrder(context.Context, *connect.Request[v1.GetOrderInput]) (*connect.Response[v1.GetOrderResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Api.GetOrder is not implemented"))
}

func (UnimplementedApiHandler) ListOrders(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListOrdersResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Api.ListOrders is not implemented"))
}

func (UnimplementedApiHandler) UpdateOrderStatus(context.Context, *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Api.UpdateOrderStatus is not implemented"))
}

func (UnimplementedApiHandler) CustomerAction(context.Context, *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Api.CustomerAction is not implemented"))
}

// WorkerClient is a client for the oms.order.v1.Worker service.
type WorkerClient interface {
	// initialize a new order
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error)
	// process a customer action
	CustomerAction(context.Context, *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error)
	// returns information about the order
	GetStatus(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.GetOrderResult], error)
	// manage the lifecycle of an order
	Order(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error)
	// reserves items to satisfy an order and returns a list of reservations for the items
	// Any unavailable items will be returned in a Reservation with Available set to false.
	// In a real system this would involve an inventory database of some kind.
	// For our purposes we just split orders arbitrarily.
	ReserveItems(context.Context, *connect.Request[v1.ReserveItemsInput]) (*connect.Response[v1.ReserveItemsResult], error)
	// updates the order status in the database
	UpdateOrderStatus(context.Context, *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error)
}

// NewWorkerClient constructs a client for the oms.order.v1.Worker service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkerClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WorkerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workerClient{
		createOrder: connect.NewClient[v1.CreateOrderInput, v1.CreateOrderResult](
			httpClient,
			baseURL+WorkerCreateOrderProcedure,
			connect.WithSchema(workerCreateOrderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		customerAction: connect.NewClient[v1.CustomerActionInput, emptypb.Empty](
			httpClient,
			baseURL+WorkerCustomerActionProcedure,
			connect.WithSchema(workerCustomerActionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getStatus: connect.NewClient[emptypb.Empty, v1.GetOrderResult](
			httpClient,
			baseURL+WorkerGetStatusProcedure,
			connect.WithSchema(workerGetStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		order: connect.NewClient[v1.CreateOrderInput, v1.CreateOrderResult](
			httpClient,
			baseURL+WorkerOrderProcedure,
			connect.WithSchema(workerOrderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		reserveItems: connect.NewClient[v1.ReserveItemsInput, v1.ReserveItemsResult](
			httpClient,
			baseURL+WorkerReserveItemsProcedure,
			connect.WithSchema(workerReserveItemsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateOrderStatus: connect.NewClient[v1.UpdateOrderStatusInput, emptypb.Empty](
			httpClient,
			baseURL+WorkerUpdateOrderStatusProcedure,
			connect.WithSchema(workerUpdateOrderStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// workerClient implements WorkerClient.
type workerClient struct {
	createOrder       *connect.Client[v1.CreateOrderInput, v1.CreateOrderResult]
	customerAction    *connect.Client[v1.CustomerActionInput, emptypb.Empty]
	getStatus         *connect.Client[emptypb.Empty, v1.GetOrderResult]
	order             *connect.Client[v1.CreateOrderInput, v1.CreateOrderResult]
	reserveItems      *connect.Client[v1.ReserveItemsInput, v1.ReserveItemsResult]
	updateOrderStatus *connect.Client[v1.UpdateOrderStatusInput, emptypb.Empty]
}

// CreateOrder calls oms.order.v1.Worker.CreateOrder.
func (c *workerClient) CreateOrder(ctx context.Context, req *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error) {
	return c.createOrder.CallUnary(ctx, req)
}

// CustomerAction calls oms.order.v1.Worker.CustomerAction.
func (c *workerClient) CustomerAction(ctx context.Context, req *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error) {
	return c.customerAction.CallUnary(ctx, req)
}

// GetStatus calls oms.order.v1.Worker.GetStatus.
func (c *workerClient) GetStatus(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.GetOrderResult], error) {
	return c.getStatus.CallUnary(ctx, req)
}

// Order calls oms.order.v1.Worker.Order.
func (c *workerClient) Order(ctx context.Context, req *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error) {
	return c.order.CallUnary(ctx, req)
}

// ReserveItems calls oms.order.v1.Worker.ReserveItems.
func (c *workerClient) ReserveItems(ctx context.Context, req *connect.Request[v1.ReserveItemsInput]) (*connect.Response[v1.ReserveItemsResult], error) {
	return c.reserveItems.CallUnary(ctx, req)
}

// UpdateOrderStatus calls oms.order.v1.Worker.UpdateOrderStatus.
func (c *workerClient) UpdateOrderStatus(ctx context.Context, req *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error) {
	return c.updateOrderStatus.CallUnary(ctx, req)
}

// WorkerHandler is an implementation of the oms.order.v1.Worker service.
type WorkerHandler interface {
	// initialize a new order
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error)
	// process a customer action
	CustomerAction(context.Context, *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error)
	// returns information about the order
	GetStatus(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.GetOrderResult], error)
	// manage the lifecycle of an order
	Order(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error)
	// reserves items to satisfy an order and returns a list of reservations for the items
	// Any unavailable items will be returned in a Reservation with Available set to false.
	// In a real system this would involve an inventory database of some kind.
	// For our purposes we just split orders arbitrarily.
	ReserveItems(context.Context, *connect.Request[v1.ReserveItemsInput]) (*connect.Response[v1.ReserveItemsResult], error)
	// updates the order status in the database
	UpdateOrderStatus(context.Context, *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error)
}

// NewWorkerHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkerHandler(svc WorkerHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	workerCreateOrderHandler := connect.NewUnaryHandler(
		WorkerCreateOrderProcedure,
		svc.CreateOrder,
		connect.WithSchema(workerCreateOrderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerCustomerActionHandler := connect.NewUnaryHandler(
		WorkerCustomerActionProcedure,
		svc.CustomerAction,
		connect.WithSchema(workerCustomerActionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerGetStatusHandler := connect.NewUnaryHandler(
		WorkerGetStatusProcedure,
		svc.GetStatus,
		connect.WithSchema(workerGetStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerOrderHandler := connect.NewUnaryHandler(
		WorkerOrderProcedure,
		svc.Order,
		connect.WithSchema(workerOrderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerReserveItemsHandler := connect.NewUnaryHandler(
		WorkerReserveItemsProcedure,
		svc.ReserveItems,
		connect.WithSchema(workerReserveItemsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerUpdateOrderStatusHandler := connect.NewUnaryHandler(
		WorkerUpdateOrderStatusProcedure,
		svc.UpdateOrderStatus,
		connect.WithSchema(workerUpdateOrderStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/oms.order.v1.Worker/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkerCreateOrderProcedure:
			workerCreateOrderHandler.ServeHTTP(w, r)
		case WorkerCustomerActionProcedure:
			workerCustomerActionHandler.ServeHTTP(w, r)
		case WorkerGetStatusProcedure:
			workerGetStatusHandler.ServeHTTP(w, r)
		case WorkerOrderProcedure:
			workerOrderHandler.ServeHTTP(w, r)
		case WorkerReserveItemsProcedure:
			workerReserveItemsHandler.ServeHTTP(w, r)
		case WorkerUpdateOrderStatusProcedure:
			workerUpdateOrderStatusHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkerHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkerHandler struct{}

func (UnimplementedWorkerHandler) CreateOrder(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Worker.CreateOrder is not implemented"))
}

func (UnimplementedWorkerHandler) CustomerAction(context.Context, *connect.Request[v1.CustomerActionInput]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Worker.CustomerAction is not implemented"))
}

func (UnimplementedWorkerHandler) GetStatus(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.GetOrderResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Worker.GetStatus is not implemented"))
}

func (UnimplementedWorkerHandler) Order(context.Context, *connect.Request[v1.CreateOrderInput]) (*connect.Response[v1.CreateOrderResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Worker.Order is not implemented"))
}

func (UnimplementedWorkerHandler) ReserveItems(context.Context, *connect.Request[v1.ReserveItemsInput]) (*connect.Response[v1.ReserveItemsResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Worker.ReserveItems is not implemented"))
}

func (UnimplementedWorkerHandler) UpdateOrderStatus(context.Context, *connect.Request[v1.UpdateOrderStatusInput]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.order.v1.Worker.UpdateOrderStatus is not implemented"))
}
