// Code generated by mockery. DO NOT EDIT.

package domain

import (
	context "context"
	domain "todo_list/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockConnectionProvider is an autogenerated mock type for the ConnectionProvider type
type MockConnectionProvider struct {
	mock.Mock
}

type MockConnectionProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnectionProvider) EXPECT() *MockConnectionProvider_Expecter {
	return &MockConnectionProvider_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with no fields
func (_m *MockConnectionProvider) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionProvider_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockConnectionProvider_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockConnectionProvider_Expecter) Close() *MockConnectionProvider_Close_Call {
	return &MockConnectionProvider_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockConnectionProvider_Close_Call) Run(run func()) *MockConnectionProvider_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConnectionProvider_Close_Call) Return(_a0 error) *MockConnectionProvider_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionProvider_Close_Call) RunAndReturn(run func() error) *MockConnectionProvider_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionProvider) Execute(_a0 context.Context, _a1 func(context.Context, domain.Connection) error) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context, domain.Connection) error) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionProvider_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockConnectionProvider_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 func(context.Context , domain.Connection) error
func (_e *MockConnectionProvider_Expecter) Execute(_a0 interface{}, _a1 interface{}) *MockConnectionProvider_Execute_Call {
	return &MockConnectionProvider_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1)}
}

func (_c *MockConnectionProvider_Execute_Call) Run(run func(_a0 context.Context, _a1 func(context.Context, domain.Connection) error)) *MockConnectionProvider_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context, domain.Connection) error))
	})
	return _c
}

func (_c *MockConnectionProvider_Execute_Call) Return(_a0 error) *MockConnectionProvider_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionProvider_Execute_Call) RunAndReturn(run func(context.Context, func(context.Context, domain.Connection) error) error) *MockConnectionProvider_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteTx provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionProvider) ExecuteTx(_a0 context.Context, _a1 func(context.Context, domain.Connection) error) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context, domain.Connection) error) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionProvider_ExecuteTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteTx'
type MockConnectionProvider_ExecuteTx_Call struct {
	*mock.Call
}

// ExecuteTx is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 func(context.Context , domain.Connection) error
func (_e *MockConnectionProvider_Expecter) ExecuteTx(_a0 interface{}, _a1 interface{}) *MockConnectionProvider_ExecuteTx_Call {
	return &MockConnectionProvider_ExecuteTx_Call{Call: _e.mock.On("ExecuteTx", _a0, _a1)}
}

func (_c *MockConnectionProvider_ExecuteTx_Call) Run(run func(_a0 context.Context, _a1 func(context.Context, domain.Connection) error)) *MockConnectionProvider_ExecuteTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context, domain.Connection) error))
	})
	return _c
}

func (_c *MockConnectionProvider_ExecuteTx_Call) Return(_a0 error) *MockConnectionProvider_ExecuteTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionProvider_ExecuteTx_Call) RunAndReturn(run func(context.Context, func(context.Context, domain.Connection) error) error) *MockConnectionProvider_ExecuteTx_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectionProvider creates a new instance of MockConnectionProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectionProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectionProvider {
	mock := &MockConnectionProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
