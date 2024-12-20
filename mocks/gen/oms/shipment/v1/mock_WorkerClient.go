// Code generated by mockery v2.49.1. DO NOT EDIT.

package shipmentv1mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	internal "go.temporal.io/sdk/client"

	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
)

// MockWorkerClient is an autogenerated mock type for the WorkerClient type
type MockWorkerClient struct {
	mock.Mock
}

type MockWorkerClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorkerClient) EXPECT() *MockWorkerClient_Expecter {
	return &MockWorkerClient_Expecter{mock: &_m.Mock}
}

// CancelWorkflow provides a mock function with given fields: ctx, workflowID, runID
func (_m *MockWorkerClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	ret := _m.Called(ctx, workflowID, runID)

	if len(ret) == 0 {
		panic("no return value specified for CancelWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, workflowID, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkerClient_CancelWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelWorkflow'
type MockWorkerClient_CancelWorkflow_Call struct {
	*mock.Call
}

// CancelWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
func (_e *MockWorkerClient_Expecter) CancelWorkflow(ctx interface{}, workflowID interface{}, runID interface{}) *MockWorkerClient_CancelWorkflow_Call {
	return &MockWorkerClient_CancelWorkflow_Call{Call: _e.mock.On("CancelWorkflow", ctx, workflowID, runID)}
}

func (_c *MockWorkerClient_CancelWorkflow_Call) Run(run func(ctx context.Context, workflowID string, runID string)) *MockWorkerClient_CancelWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockWorkerClient_CancelWorkflow_Call) Return(_a0 error) *MockWorkerClient_CancelWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerClient_CancelWorkflow_Call) RunAndReturn(run func(context.Context, string, string) error) *MockWorkerClient_CancelWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// GetShipment provides a mock function with given fields: ctx, workflowID, runID
func (_m *MockWorkerClient) GetShipment(ctx context.Context, workflowID string, runID string) shipmentv1.ShipmentRun {
	ret := _m.Called(ctx, workflowID, runID)

	if len(ret) == 0 {
		panic("no return value specified for GetShipment")
	}

	var r0 shipmentv1.ShipmentRun
	if rf, ok := ret.Get(0).(func(context.Context, string, string) shipmentv1.ShipmentRun); ok {
		r0 = rf(ctx, workflowID, runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shipmentv1.ShipmentRun)
		}
	}

	return r0
}

// MockWorkerClient_GetShipment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShipment'
type MockWorkerClient_GetShipment_Call struct {
	*mock.Call
}

// GetShipment is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
func (_e *MockWorkerClient_Expecter) GetShipment(ctx interface{}, workflowID interface{}, runID interface{}) *MockWorkerClient_GetShipment_Call {
	return &MockWorkerClient_GetShipment_Call{Call: _e.mock.On("GetShipment", ctx, workflowID, runID)}
}

func (_c *MockWorkerClient_GetShipment_Call) Run(run func(ctx context.Context, workflowID string, runID string)) *MockWorkerClient_GetShipment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockWorkerClient_GetShipment_Call) Return(_a0 shipmentv1.ShipmentRun) *MockWorkerClient_GetShipment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerClient_GetShipment_Call) RunAndReturn(run func(context.Context, string, string) shipmentv1.ShipmentRun) *MockWorkerClient_GetShipment_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatus provides a mock function with given fields: ctx, workflowID, runID
func (_m *MockWorkerClient) GetStatus(ctx context.Context, workflowID string, runID string) (*shipmentv1.GetShipmentResult, error) {
	ret := _m.Called(ctx, workflowID, runID)

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 *shipmentv1.GetShipmentResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*shipmentv1.GetShipmentResult, error)); ok {
		return rf(ctx, workflowID, runID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *shipmentv1.GetShipmentResult); ok {
		r0 = rf(ctx, workflowID, runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shipmentv1.GetShipmentResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, workflowID, runID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkerClient_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type MockWorkerClient_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
func (_e *MockWorkerClient_Expecter) GetStatus(ctx interface{}, workflowID interface{}, runID interface{}) *MockWorkerClient_GetStatus_Call {
	return &MockWorkerClient_GetStatus_Call{Call: _e.mock.On("GetStatus", ctx, workflowID, runID)}
}

func (_c *MockWorkerClient_GetStatus_Call) Run(run func(ctx context.Context, workflowID string, runID string)) *MockWorkerClient_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockWorkerClient_GetStatus_Call) Return(_a0 *shipmentv1.GetShipmentResult, _a1 error) *MockWorkerClient_GetStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkerClient_GetStatus_Call) RunAndReturn(run func(context.Context, string, string) (*shipmentv1.GetShipmentResult, error)) *MockWorkerClient_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetUpdateShipmentStatus provides a mock function with given fields: ctx, req
func (_m *MockWorkerClient) GetUpdateShipmentStatus(ctx context.Context, req internal.GetWorkflowUpdateHandleOptions) (shipmentv1.UpdateShipmentStatusHandle, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetUpdateShipmentStatus")
	}

	var r0 shipmentv1.UpdateShipmentStatusHandle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, internal.GetWorkflowUpdateHandleOptions) (shipmentv1.UpdateShipmentStatusHandle, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, internal.GetWorkflowUpdateHandleOptions) shipmentv1.UpdateShipmentStatusHandle); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shipmentv1.UpdateShipmentStatusHandle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, internal.GetWorkflowUpdateHandleOptions) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkerClient_GetUpdateShipmentStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUpdateShipmentStatus'
type MockWorkerClient_GetUpdateShipmentStatus_Call struct {
	*mock.Call
}

// GetUpdateShipmentStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - req internal.GetWorkflowUpdateHandleOptions
func (_e *MockWorkerClient_Expecter) GetUpdateShipmentStatus(ctx interface{}, req interface{}) *MockWorkerClient_GetUpdateShipmentStatus_Call {
	return &MockWorkerClient_GetUpdateShipmentStatus_Call{Call: _e.mock.On("GetUpdateShipmentStatus", ctx, req)}
}

func (_c *MockWorkerClient_GetUpdateShipmentStatus_Call) Run(run func(ctx context.Context, req internal.GetWorkflowUpdateHandleOptions)) *MockWorkerClient_GetUpdateShipmentStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(internal.GetWorkflowUpdateHandleOptions))
	})
	return _c
}

func (_c *MockWorkerClient_GetUpdateShipmentStatus_Call) Return(_a0 shipmentv1.UpdateShipmentStatusHandle, _a1 error) *MockWorkerClient_GetUpdateShipmentStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkerClient_GetUpdateShipmentStatus_Call) RunAndReturn(run func(context.Context, internal.GetWorkflowUpdateHandleOptions) (shipmentv1.UpdateShipmentStatusHandle, error)) *MockWorkerClient_GetUpdateShipmentStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Shipment provides a mock function with given fields: ctx, req, opts
func (_m *MockWorkerClient) Shipment(ctx context.Context, req *shipmentv1.CreateShipmentInput, opts ...*shipmentv1.ShipmentOptions) (*shipmentv1.CreateShipmentResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Shipment")
	}

	var r0 *shipmentv1.CreateShipmentResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) (*shipmentv1.CreateShipmentResult, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) *shipmentv1.CreateShipmentResult); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shipmentv1.CreateShipmentResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkerClient_Shipment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shipment'
type MockWorkerClient_Shipment_Call struct {
	*mock.Call
}

// Shipment is a helper method to define mock.On call
//   - ctx context.Context
//   - req *shipmentv1.CreateShipmentInput
//   - opts ...*shipmentv1.ShipmentOptions
func (_e *MockWorkerClient_Expecter) Shipment(ctx interface{}, req interface{}, opts ...interface{}) *MockWorkerClient_Shipment_Call {
	return &MockWorkerClient_Shipment_Call{Call: _e.mock.On("Shipment",
		append([]interface{}{ctx, req}, opts...)...)}
}

func (_c *MockWorkerClient_Shipment_Call) Run(run func(ctx context.Context, req *shipmentv1.CreateShipmentInput, opts ...*shipmentv1.ShipmentOptions)) *MockWorkerClient_Shipment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*shipmentv1.ShipmentOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*shipmentv1.ShipmentOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*shipmentv1.CreateShipmentInput), variadicArgs...)
	})
	return _c
}

func (_c *MockWorkerClient_Shipment_Call) Return(_a0 *shipmentv1.CreateShipmentResult, _a1 error) *MockWorkerClient_Shipment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkerClient_Shipment_Call) RunAndReturn(run func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) (*shipmentv1.CreateShipmentResult, error)) *MockWorkerClient_Shipment_Call {
	_c.Call.Return(run)
	return _c
}

// ShipmentAsync provides a mock function with given fields: ctx, req, opts
func (_m *MockWorkerClient) ShipmentAsync(ctx context.Context, req *shipmentv1.CreateShipmentInput, opts ...*shipmentv1.ShipmentOptions) (shipmentv1.ShipmentRun, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ShipmentAsync")
	}

	var r0 shipmentv1.ShipmentRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) (shipmentv1.ShipmentRun, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) shipmentv1.ShipmentRun); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shipmentv1.ShipmentRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkerClient_ShipmentAsync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShipmentAsync'
type MockWorkerClient_ShipmentAsync_Call struct {
	*mock.Call
}

// ShipmentAsync is a helper method to define mock.On call
//   - ctx context.Context
//   - req *shipmentv1.CreateShipmentInput
//   - opts ...*shipmentv1.ShipmentOptions
func (_e *MockWorkerClient_Expecter) ShipmentAsync(ctx interface{}, req interface{}, opts ...interface{}) *MockWorkerClient_ShipmentAsync_Call {
	return &MockWorkerClient_ShipmentAsync_Call{Call: _e.mock.On("ShipmentAsync",
		append([]interface{}{ctx, req}, opts...)...)}
}

func (_c *MockWorkerClient_ShipmentAsync_Call) Run(run func(ctx context.Context, req *shipmentv1.CreateShipmentInput, opts ...*shipmentv1.ShipmentOptions)) *MockWorkerClient_ShipmentAsync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*shipmentv1.ShipmentOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*shipmentv1.ShipmentOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*shipmentv1.CreateShipmentInput), variadicArgs...)
	})
	return _c
}

func (_c *MockWorkerClient_ShipmentAsync_Call) Return(_a0 shipmentv1.ShipmentRun, _a1 error) *MockWorkerClient_ShipmentAsync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkerClient_ShipmentAsync_Call) RunAndReturn(run func(context.Context, *shipmentv1.CreateShipmentInput, ...*shipmentv1.ShipmentOptions) (shipmentv1.ShipmentRun, error)) *MockWorkerClient_ShipmentAsync_Call {
	_c.Call.Return(run)
	return _c
}

// ShipmentStatusUpdated provides a mock function with given fields: ctx, workflowID, runID, signal
func (_m *MockWorkerClient) ShipmentStatusUpdated(ctx context.Context, workflowID string, runID string, signal *shipmentv1.UpdateShipmentStatusInput) error {
	ret := _m.Called(ctx, workflowID, runID, signal)

	if len(ret) == 0 {
		panic("no return value specified for ShipmentStatusUpdated")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput) error); ok {
		r0 = rf(ctx, workflowID, runID, signal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkerClient_ShipmentStatusUpdated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShipmentStatusUpdated'
type MockWorkerClient_ShipmentStatusUpdated_Call struct {
	*mock.Call
}

// ShipmentStatusUpdated is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
//   - signal *shipmentv1.UpdateShipmentStatusInput
func (_e *MockWorkerClient_Expecter) ShipmentStatusUpdated(ctx interface{}, workflowID interface{}, runID interface{}, signal interface{}) *MockWorkerClient_ShipmentStatusUpdated_Call {
	return &MockWorkerClient_ShipmentStatusUpdated_Call{Call: _e.mock.On("ShipmentStatusUpdated", ctx, workflowID, runID, signal)}
}

func (_c *MockWorkerClient_ShipmentStatusUpdated_Call) Run(run func(ctx context.Context, workflowID string, runID string, signal *shipmentv1.UpdateShipmentStatusInput)) *MockWorkerClient_ShipmentStatusUpdated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*shipmentv1.UpdateShipmentStatusInput))
	})
	return _c
}

func (_c *MockWorkerClient_ShipmentStatusUpdated_Call) Return(_a0 error) *MockWorkerClient_ShipmentStatusUpdated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerClient_ShipmentStatusUpdated_Call) RunAndReturn(run func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput) error) *MockWorkerClient_ShipmentStatusUpdated_Call {
	_c.Call.Return(run)
	return _c
}

// TerminateWorkflow provides a mock function with given fields: ctx, workflowID, runID, reason, details
func (_m *MockWorkerClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, workflowID, runID, reason)
	_ca = append(_ca, details...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TerminateWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, ...interface{}) error); ok {
		r0 = rf(ctx, workflowID, runID, reason, details...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkerClient_TerminateWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TerminateWorkflow'
type MockWorkerClient_TerminateWorkflow_Call struct {
	*mock.Call
}

// TerminateWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
//   - reason string
//   - details ...interface{}
func (_e *MockWorkerClient_Expecter) TerminateWorkflow(ctx interface{}, workflowID interface{}, runID interface{}, reason interface{}, details ...interface{}) *MockWorkerClient_TerminateWorkflow_Call {
	return &MockWorkerClient_TerminateWorkflow_Call{Call: _e.mock.On("TerminateWorkflow",
		append([]interface{}{ctx, workflowID, runID, reason}, details...)...)}
}

func (_c *MockWorkerClient_TerminateWorkflow_Call) Run(run func(ctx context.Context, workflowID string, runID string, reason string, details ...interface{})) *MockWorkerClient_TerminateWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockWorkerClient_TerminateWorkflow_Call) Return(_a0 error) *MockWorkerClient_TerminateWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerClient_TerminateWorkflow_Call) RunAndReturn(run func(context.Context, string, string, string, ...interface{}) error) *MockWorkerClient_TerminateWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateShipmentStatus provides a mock function with given fields: ctx, workflowID, runID, req, opts
func (_m *MockWorkerClient) UpdateShipmentStatus(ctx context.Context, workflowID string, runID string, req *shipmentv1.UpdateShipmentStatusInput, opts ...*shipmentv1.UpdateShipmentStatusOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, workflowID, runID, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateShipmentStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput, ...*shipmentv1.UpdateShipmentStatusOptions) error); ok {
		r0 = rf(ctx, workflowID, runID, req, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkerClient_UpdateShipmentStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateShipmentStatus'
type MockWorkerClient_UpdateShipmentStatus_Call struct {
	*mock.Call
}

// UpdateShipmentStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
//   - req *shipmentv1.UpdateShipmentStatusInput
//   - opts ...*shipmentv1.UpdateShipmentStatusOptions
func (_e *MockWorkerClient_Expecter) UpdateShipmentStatus(ctx interface{}, workflowID interface{}, runID interface{}, req interface{}, opts ...interface{}) *MockWorkerClient_UpdateShipmentStatus_Call {
	return &MockWorkerClient_UpdateShipmentStatus_Call{Call: _e.mock.On("UpdateShipmentStatus",
		append([]interface{}{ctx, workflowID, runID, req}, opts...)...)}
}

func (_c *MockWorkerClient_UpdateShipmentStatus_Call) Run(run func(ctx context.Context, workflowID string, runID string, req *shipmentv1.UpdateShipmentStatusInput, opts ...*shipmentv1.UpdateShipmentStatusOptions)) *MockWorkerClient_UpdateShipmentStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*shipmentv1.UpdateShipmentStatusOptions, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(*shipmentv1.UpdateShipmentStatusOptions)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*shipmentv1.UpdateShipmentStatusInput), variadicArgs...)
	})
	return _c
}

func (_c *MockWorkerClient_UpdateShipmentStatus_Call) Return(_a0 error) *MockWorkerClient_UpdateShipmentStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkerClient_UpdateShipmentStatus_Call) RunAndReturn(run func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput, ...*shipmentv1.UpdateShipmentStatusOptions) error) *MockWorkerClient_UpdateShipmentStatus_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateShipmentStatusAsync provides a mock function with given fields: ctx, workflowID, runID, req, opts
func (_m *MockWorkerClient) UpdateShipmentStatusAsync(ctx context.Context, workflowID string, runID string, req *shipmentv1.UpdateShipmentStatusInput, opts ...*shipmentv1.UpdateShipmentStatusOptions) (shipmentv1.UpdateShipmentStatusHandle, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, workflowID, runID, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateShipmentStatusAsync")
	}

	var r0 shipmentv1.UpdateShipmentStatusHandle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput, ...*shipmentv1.UpdateShipmentStatusOptions) (shipmentv1.UpdateShipmentStatusHandle, error)); ok {
		return rf(ctx, workflowID, runID, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput, ...*shipmentv1.UpdateShipmentStatusOptions) shipmentv1.UpdateShipmentStatusHandle); ok {
		r0 = rf(ctx, workflowID, runID, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shipmentv1.UpdateShipmentStatusHandle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput, ...*shipmentv1.UpdateShipmentStatusOptions) error); ok {
		r1 = rf(ctx, workflowID, runID, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkerClient_UpdateShipmentStatusAsync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateShipmentStatusAsync'
type MockWorkerClient_UpdateShipmentStatusAsync_Call struct {
	*mock.Call
}

// UpdateShipmentStatusAsync is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
//   - req *shipmentv1.UpdateShipmentStatusInput
//   - opts ...*shipmentv1.UpdateShipmentStatusOptions
func (_e *MockWorkerClient_Expecter) UpdateShipmentStatusAsync(ctx interface{}, workflowID interface{}, runID interface{}, req interface{}, opts ...interface{}) *MockWorkerClient_UpdateShipmentStatusAsync_Call {
	return &MockWorkerClient_UpdateShipmentStatusAsync_Call{Call: _e.mock.On("UpdateShipmentStatusAsync",
		append([]interface{}{ctx, workflowID, runID, req}, opts...)...)}
}

func (_c *MockWorkerClient_UpdateShipmentStatusAsync_Call) Run(run func(ctx context.Context, workflowID string, runID string, req *shipmentv1.UpdateShipmentStatusInput, opts ...*shipmentv1.UpdateShipmentStatusOptions)) *MockWorkerClient_UpdateShipmentStatusAsync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*shipmentv1.UpdateShipmentStatusOptions, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(*shipmentv1.UpdateShipmentStatusOptions)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*shipmentv1.UpdateShipmentStatusInput), variadicArgs...)
	})
	return _c
}

func (_c *MockWorkerClient_UpdateShipmentStatusAsync_Call) Return(_a0 shipmentv1.UpdateShipmentStatusHandle, _a1 error) *MockWorkerClient_UpdateShipmentStatusAsync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkerClient_UpdateShipmentStatusAsync_Call) RunAndReturn(run func(context.Context, string, string, *shipmentv1.UpdateShipmentStatusInput, ...*shipmentv1.UpdateShipmentStatusOptions) (shipmentv1.UpdateShipmentStatusHandle, error)) *MockWorkerClient_UpdateShipmentStatusAsync_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorkerClient creates a new instance of MockWorkerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorkerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkerClient {
	mock := &MockWorkerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
