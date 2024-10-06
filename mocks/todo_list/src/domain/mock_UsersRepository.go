// Code generated by mockery. DO NOT EDIT.

package domain

import (
	context "context"
	domain "todo_list/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockUsersRepository is an autogenerated mock type for the UsersRepository type
type MockUsersRepository struct {
	mock.Mock
}

type MockUsersRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUsersRepository) EXPECT() *MockUsersRepository_Expecter {
	return &MockUsersRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockUsersRepository) Create(_a0 context.Context, _a1 domain.Connection, _a2 domain.User) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.User) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUsersRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.User
func (_e *MockUsersRepository_Expecter) Create(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockUsersRepository_Create_Call {
	return &MockUsersRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1, _a2)}
}

func (_c *MockUsersRepository_Create_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.User)) *MockUsersRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.User))
	})
	return _c
}

func (_c *MockUsersRepository_Create_Call) Return(_a0 error) *MockUsersRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersRepository_Create_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.User) error) *MockUsersRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockUsersRepository) Delete(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUsersRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.UserID
func (_e *MockUsersRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockUsersRepository_Delete_Call {
	return &MockUsersRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1, _a2)}
}

func (_c *MockUsersRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID)) *MockUsersRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.UserID))
	})
	return _c
}

func (_c *MockUsersRepository_Delete_Call) Return(_a0 error) *MockUsersRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersRepository_Delete_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.UserID) error) *MockUsersRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockUsersRepository) Read(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID) (domain.User, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID) (domain.User, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID) domain.User); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Connection, domain.UserID) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockUsersRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.UserID
func (_e *MockUsersRepository_Expecter) Read(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockUsersRepository_Read_Call {
	return &MockUsersRepository_Read_Call{Call: _e.mock.On("Read", _a0, _a1, _a2)}
}

func (_c *MockUsersRepository_Read_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID)) *MockUsersRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.UserID))
	})
	return _c
}

func (_c *MockUsersRepository_Read_Call) Return(_a0 domain.User, _a1 error) *MockUsersRepository_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersRepository_Read_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.UserID) (domain.User, error)) *MockUsersRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockUsersRepository) Update(_a0 context.Context, _a1 domain.Connection, _a2 domain.User) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.User) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUsersRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.User
func (_e *MockUsersRepository_Expecter) Update(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockUsersRepository_Update_Call {
	return &MockUsersRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1, _a2)}
}

func (_c *MockUsersRepository_Update_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.User)) *MockUsersRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.User))
	})
	return _c
}

func (_c *MockUsersRepository_Update_Call) Return(_a0 error) *MockUsersRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersRepository_Update_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.User) error) *MockUsersRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUsersRepository creates a new instance of MockUsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUsersRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUsersRepository {
	mock := &MockUsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
