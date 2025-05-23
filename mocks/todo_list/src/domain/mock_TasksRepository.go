// Code generated by mockery. DO NOT EDIT.

package domain

import (
	context "context"
	domain "todo_list/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockTasksRepository is an autogenerated mock type for the TasksRepository type
type MockTasksRepository struct {
	mock.Mock
}

type MockTasksRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTasksRepository) EXPECT() *MockTasksRepository_Expecter {
	return &MockTasksRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockTasksRepository) Create(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.Task) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID, domain.Task) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTasksRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockTasksRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.UserID
//   - _a3 domain.Task
func (_e *MockTasksRepository_Expecter) Create(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *MockTasksRepository_Create_Call {
	return &MockTasksRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1, _a2, _a3)}
}

func (_c *MockTasksRepository_Create_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.Task)) *MockTasksRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.UserID), args[3].(domain.Task))
	})
	return _c
}

func (_c *MockTasksRepository_Create_Call) Return(_a0 error) *MockTasksRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTasksRepository_Create_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.UserID, domain.Task) error) *MockTasksRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockTasksRepository) Delete(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.TaskID) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID, domain.TaskID) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTasksRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockTasksRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.UserID
//   - _a3 domain.TaskID
func (_e *MockTasksRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *MockTasksRepository_Delete_Call {
	return &MockTasksRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1, _a2, _a3)}
}

func (_c *MockTasksRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.TaskID)) *MockTasksRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.UserID), args[3].(domain.TaskID))
	})
	return _c
}

func (_c *MockTasksRepository_Delete_Call) Return(_a0 error) *MockTasksRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTasksRepository_Delete_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.UserID, domain.TaskID) error) *MockTasksRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTasks provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockTasksRepository) GetAllTasks(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 []domain.ListID) ([]domain.Task, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTasks")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID, []domain.ListID) ([]domain.Task, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID, []domain.ListID) []domain.Task); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Connection, domain.UserID, []domain.ListID) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTasksRepository_GetAllTasks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTasks'
type MockTasksRepository_GetAllTasks_Call struct {
	*mock.Call
}

// GetAllTasks is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.UserID
//   - _a3 []domain.ListID
func (_e *MockTasksRepository_Expecter) GetAllTasks(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *MockTasksRepository_GetAllTasks_Call {
	return &MockTasksRepository_GetAllTasks_Call{Call: _e.mock.On("GetAllTasks", _a0, _a1, _a2, _a3)}
}

func (_c *MockTasksRepository_GetAllTasks_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 []domain.ListID)) *MockTasksRepository_GetAllTasks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.UserID), args[3].([]domain.ListID))
	})
	return _c
}

func (_c *MockTasksRepository_GetAllTasks_Call) Return(_a0 []domain.Task, _a1 error) *MockTasksRepository_GetAllTasks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTasksRepository_GetAllTasks_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.UserID, []domain.ListID) ([]domain.Task, error)) *MockTasksRepository_GetAllTasks_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockTasksRepository) Read(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.TaskID) (domain.Task, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID, domain.TaskID) (domain.Task, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID, domain.TaskID) domain.Task); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Connection, domain.UserID, domain.TaskID) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTasksRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockTasksRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.UserID
//   - _a3 domain.TaskID
func (_e *MockTasksRepository_Expecter) Read(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *MockTasksRepository_Read_Call {
	return &MockTasksRepository_Read_Call{Call: _e.mock.On("Read", _a0, _a1, _a2, _a3)}
}

func (_c *MockTasksRepository_Read_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.TaskID)) *MockTasksRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.UserID), args[3].(domain.TaskID))
	})
	return _c
}

func (_c *MockTasksRepository_Read_Call) Return(_a0 domain.Task, _a1 error) *MockTasksRepository_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTasksRepository_Read_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.UserID, domain.TaskID) (domain.Task, error)) *MockTasksRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockTasksRepository) Update(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.Task) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Connection, domain.UserID, domain.Task) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTasksRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockTasksRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 domain.Connection
//   - _a2 domain.UserID
//   - _a3 domain.Task
func (_e *MockTasksRepository_Expecter) Update(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *MockTasksRepository_Update_Call {
	return &MockTasksRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1, _a2, _a3)}
}

func (_c *MockTasksRepository_Update_Call) Run(run func(_a0 context.Context, _a1 domain.Connection, _a2 domain.UserID, _a3 domain.Task)) *MockTasksRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Connection), args[2].(domain.UserID), args[3].(domain.Task))
	})
	return _c
}

func (_c *MockTasksRepository_Update_Call) Return(_a0 error) *MockTasksRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTasksRepository_Update_Call) RunAndReturn(run func(context.Context, domain.Connection, domain.UserID, domain.Task) error) *MockTasksRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTasksRepository creates a new instance of MockTasksRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTasksRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTasksRepository {
	mock := &MockTasksRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
