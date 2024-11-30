// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: oms/billing/v1/billing.proto

package billingv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
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
	ApiName = "oms.billing.v1.Api"
	// WorkerName is the fully-qualified name of the Worker service.
	WorkerName = "oms.billing.v1.Worker"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ApiChargeProcedure is the fully-qualified name of the Api's Charge RPC.
	ApiChargeProcedure = "/oms.billing.v1.Api/Charge"
	// WorkerChargeProcedure is the fully-qualified name of the Worker's Charge RPC.
	WorkerChargeProcedure = "/oms.billing.v1.Worker/Charge"
	// WorkerCheckFraudProcedure is the fully-qualified name of the Worker's CheckFraud RPC.
	WorkerCheckFraudProcedure = "/oms.billing.v1.Worker/CheckFraud"
	// WorkerGenerateInvoiceProcedure is the fully-qualified name of the Worker's GenerateInvoice RPC.
	WorkerGenerateInvoiceProcedure = "/oms.billing.v1.Worker/GenerateInvoice"
	// WorkerProcessPaymentProcedure is the fully-qualified name of the Worker's ProcessPayment RPC.
	WorkerProcessPaymentProcedure = "/oms.billing.v1.Worker/ProcessPayment"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	apiServiceDescriptor                  = v1.File_oms_billing_v1_billing_proto.Services().ByName("Api")
	apiChargeMethodDescriptor             = apiServiceDescriptor.Methods().ByName("Charge")
	workerServiceDescriptor               = v1.File_oms_billing_v1_billing_proto.Services().ByName("Worker")
	workerChargeMethodDescriptor          = workerServiceDescriptor.Methods().ByName("Charge")
	workerCheckFraudMethodDescriptor      = workerServiceDescriptor.Methods().ByName("CheckFraud")
	workerGenerateInvoiceMethodDescriptor = workerServiceDescriptor.Methods().ByName("GenerateInvoice")
	workerProcessPaymentMethodDescriptor  = workerServiceDescriptor.Methods().ByName("ProcessPayment")
)

// ApiClient is a client for the oms.billing.v1.Api service.
type ApiClient interface {
	// process an order from a customer
	Charge(context.Context, *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error)
}

// NewApiClient constructs a client for the oms.billing.v1.Api service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewApiClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ApiClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &apiClient{
		charge: connect.NewClient[v1.ChargeInput, v1.ChargeResult](
			httpClient,
			baseURL+ApiChargeProcedure,
			connect.WithSchema(apiChargeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// apiClient implements ApiClient.
type apiClient struct {
	charge *connect.Client[v1.ChargeInput, v1.ChargeResult]
}

// Charge calls oms.billing.v1.Api.Charge.
func (c *apiClient) Charge(ctx context.Context, req *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error) {
	return c.charge.CallUnary(ctx, req)
}

// ApiHandler is an implementation of the oms.billing.v1.Api service.
type ApiHandler interface {
	// process an order from a customer
	Charge(context.Context, *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error)
}

// NewApiHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewApiHandler(svc ApiHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	apiChargeHandler := connect.NewUnaryHandler(
		ApiChargeProcedure,
		svc.Charge,
		connect.WithSchema(apiChargeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/oms.billing.v1.Api/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ApiChargeProcedure:
			apiChargeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedApiHandler returns CodeUnimplemented from all methods.
type UnimplementedApiHandler struct{}

func (UnimplementedApiHandler) Charge(context.Context, *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.billing.v1.Api.Charge is not implemented"))
}

// WorkerClient is a client for the oms.billing.v1.Worker service.
type WorkerClient interface {
	// durably process an order from a customer
	Charge(context.Context, *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error)
	// determines whether the charge is fraudulent
	CheckFraud(context.Context, *connect.Request[v1.CheckFraudInput]) (*connect.Response[v1.CheckFraudResult], error)
	// generates an invoice file
	GenerateInvoice(context.Context, *connect.Request[v1.GenerateInvoiceInput]) (*connect.Response[v1.GenerateInvoiceResult], error)
	// processes the customer payment
	ProcessPayment(context.Context, *connect.Request[v1.ProcessPaymentInput]) (*connect.Response[v1.ProcessPaymentResult], error)
}

// NewWorkerClient constructs a client for the oms.billing.v1.Worker service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkerClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WorkerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workerClient{
		charge: connect.NewClient[v1.ChargeInput, v1.ChargeResult](
			httpClient,
			baseURL+WorkerChargeProcedure,
			connect.WithSchema(workerChargeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		checkFraud: connect.NewClient[v1.CheckFraudInput, v1.CheckFraudResult](
			httpClient,
			baseURL+WorkerCheckFraudProcedure,
			connect.WithSchema(workerCheckFraudMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		generateInvoice: connect.NewClient[v1.GenerateInvoiceInput, v1.GenerateInvoiceResult](
			httpClient,
			baseURL+WorkerGenerateInvoiceProcedure,
			connect.WithSchema(workerGenerateInvoiceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		processPayment: connect.NewClient[v1.ProcessPaymentInput, v1.ProcessPaymentResult](
			httpClient,
			baseURL+WorkerProcessPaymentProcedure,
			connect.WithSchema(workerProcessPaymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// workerClient implements WorkerClient.
type workerClient struct {
	charge          *connect.Client[v1.ChargeInput, v1.ChargeResult]
	checkFraud      *connect.Client[v1.CheckFraudInput, v1.CheckFraudResult]
	generateInvoice *connect.Client[v1.GenerateInvoiceInput, v1.GenerateInvoiceResult]
	processPayment  *connect.Client[v1.ProcessPaymentInput, v1.ProcessPaymentResult]
}

// Charge calls oms.billing.v1.Worker.Charge.
func (c *workerClient) Charge(ctx context.Context, req *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error) {
	return c.charge.CallUnary(ctx, req)
}

// CheckFraud calls oms.billing.v1.Worker.CheckFraud.
func (c *workerClient) CheckFraud(ctx context.Context, req *connect.Request[v1.CheckFraudInput]) (*connect.Response[v1.CheckFraudResult], error) {
	return c.checkFraud.CallUnary(ctx, req)
}

// GenerateInvoice calls oms.billing.v1.Worker.GenerateInvoice.
func (c *workerClient) GenerateInvoice(ctx context.Context, req *connect.Request[v1.GenerateInvoiceInput]) (*connect.Response[v1.GenerateInvoiceResult], error) {
	return c.generateInvoice.CallUnary(ctx, req)
}

// ProcessPayment calls oms.billing.v1.Worker.ProcessPayment.
func (c *workerClient) ProcessPayment(ctx context.Context, req *connect.Request[v1.ProcessPaymentInput]) (*connect.Response[v1.ProcessPaymentResult], error) {
	return c.processPayment.CallUnary(ctx, req)
}

// WorkerHandler is an implementation of the oms.billing.v1.Worker service.
type WorkerHandler interface {
	// durably process an order from a customer
	Charge(context.Context, *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error)
	// determines whether the charge is fraudulent
	CheckFraud(context.Context, *connect.Request[v1.CheckFraudInput]) (*connect.Response[v1.CheckFraudResult], error)
	// generates an invoice file
	GenerateInvoice(context.Context, *connect.Request[v1.GenerateInvoiceInput]) (*connect.Response[v1.GenerateInvoiceResult], error)
	// processes the customer payment
	ProcessPayment(context.Context, *connect.Request[v1.ProcessPaymentInput]) (*connect.Response[v1.ProcessPaymentResult], error)
}

// NewWorkerHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkerHandler(svc WorkerHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	workerChargeHandler := connect.NewUnaryHandler(
		WorkerChargeProcedure,
		svc.Charge,
		connect.WithSchema(workerChargeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerCheckFraudHandler := connect.NewUnaryHandler(
		WorkerCheckFraudProcedure,
		svc.CheckFraud,
		connect.WithSchema(workerCheckFraudMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerGenerateInvoiceHandler := connect.NewUnaryHandler(
		WorkerGenerateInvoiceProcedure,
		svc.GenerateInvoice,
		connect.WithSchema(workerGenerateInvoiceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workerProcessPaymentHandler := connect.NewUnaryHandler(
		WorkerProcessPaymentProcedure,
		svc.ProcessPayment,
		connect.WithSchema(workerProcessPaymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/oms.billing.v1.Worker/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkerChargeProcedure:
			workerChargeHandler.ServeHTTP(w, r)
		case WorkerCheckFraudProcedure:
			workerCheckFraudHandler.ServeHTTP(w, r)
		case WorkerGenerateInvoiceProcedure:
			workerGenerateInvoiceHandler.ServeHTTP(w, r)
		case WorkerProcessPaymentProcedure:
			workerProcessPaymentHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkerHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkerHandler struct{}

func (UnimplementedWorkerHandler) Charge(context.Context, *connect.Request[v1.ChargeInput]) (*connect.Response[v1.ChargeResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.billing.v1.Worker.Charge is not implemented"))
}

func (UnimplementedWorkerHandler) CheckFraud(context.Context, *connect.Request[v1.CheckFraudInput]) (*connect.Response[v1.CheckFraudResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.billing.v1.Worker.CheckFraud is not implemented"))
}

func (UnimplementedWorkerHandler) GenerateInvoice(context.Context, *connect.Request[v1.GenerateInvoiceInput]) (*connect.Response[v1.GenerateInvoiceResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.billing.v1.Worker.GenerateInvoice is not implemented"))
}

func (UnimplementedWorkerHandler) ProcessPayment(context.Context, *connect.Request[v1.ProcessPaymentInput]) (*connect.Response[v1.ProcessPaymentResult], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("oms.billing.v1.Worker.ProcessPayment is not implemented"))
}