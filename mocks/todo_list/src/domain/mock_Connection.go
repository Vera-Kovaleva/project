// Code generated by mockery. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockConnection is an autogenerated mock type for the Connection type
type MockConnection struct {
	mock.Mock
}

type MockConnection_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnection) EXPECT() *MockConnection_Expecter {
	return &MockConnection_Expecter{mock: &_m.Mock}
}

// ExecContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockConnection) ExecContext(_a0 context.Context, _a1 string, _a2 ...any) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...any) (int64, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...any) int64); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...any) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnection_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type MockConnection_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 ...any
func (_e *MockConnection_Expecter) ExecContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *MockConnection_ExecContext_Call {
	return &MockConnection_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *MockConnection_ExecContext_Call) Run(run func(_a0 context.Context, _a1 string, _a2 ...any)) *MockConnection_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]any, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(any)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockConnection_ExecContext_Call) Return(_a0 int64, _a1 error) *MockConnection_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnection_ExecContext_Call) RunAndReturn(run func(context.Context, string, ...any) (int64, error)) *MockConnection_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockConnection) GetContext(_a0 context.Context, _a1 any, _a2 string, _a3 ...any) error {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _a3...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, any, string, ...any) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnection_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockConnection_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 any
//   - _a2 string
//   - _a3 ...any
func (_e *MockConnection_Expecter) GetContext(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 ...interface{}) *MockConnection_GetContext_Call {
	return &MockConnection_GetContext_Call{Call: _e.mock.On("GetContext",
		append([]interface{}{_a0, _a1, _a2}, _a3...)...)}
}

func (_c *MockConnection_GetContext_Call) Run(run func(_a0 context.Context, _a1 any, _a2 string, _a3 ...any)) *MockConnection_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]any, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(any)
			}
		}
		run(args[0].(context.Context), args[1].(any), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockConnection_GetContext_Call) Return(_a0 error) *MockConnection_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnection_GetContext_Call) RunAndReturn(run func(context.Context, any, string, ...any) error) *MockConnection_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// SelectContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockConnection) SelectContext(_a0 context.Context, _a1 any, _a2 string, _a3 ...any) error {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _a3...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SelectContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, any, string, ...any) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnection_SelectContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectContext'
type MockConnection_SelectContext_Call struct {
	*mock.Call
}

// SelectContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 any
//   - _a2 string
//   - _a3 ...any
func (_e *MockConnection_Expecter) SelectContext(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 ...interface{}) *MockConnection_SelectContext_Call {
	return &MockConnection_SelectContext_Call{Call: _e.mock.On("SelectContext",
		append([]interface{}{_a0, _a1, _a2}, _a3...)...)}
}

func (_c *MockConnection_SelectContext_Call) Run(run func(_a0 context.Context, _a1 any, _a2 string, _a3 ...any)) *MockConnection_SelectContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]any, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(any)
			}
		}
		run(args[0].(context.Context), args[1].(any), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockConnection_SelectContext_Call) Return(_a0 error) *MockConnection_SelectContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnection_SelectContext_Call) RunAndReturn(run func(context.Context, any, string, ...any) error) *MockConnection_SelectContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnection creates a new instance of MockConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnection {
	mock := &MockConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
