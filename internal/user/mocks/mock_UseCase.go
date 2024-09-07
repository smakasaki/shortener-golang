// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/smakasaki/shortener/internal/user"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUseCase is an autogenerated mock type for the UseCase type
type MockUseCase struct {
	mock.Mock
}

type MockUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUseCase) EXPECT() *MockUseCase_Expecter {
	return &MockUseCase_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, _a1
func (_m *MockUseCase) CreateUser(ctx context.Context, _a1 *user.User) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUseCase_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockUseCase_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *user.User
func (_e *MockUseCase_Expecter) CreateUser(ctx interface{}, _a1 interface{}) *MockUseCase_CreateUser_Call {
	return &MockUseCase_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, _a1)}
}

func (_c *MockUseCase_CreateUser_Call) Run(run func(ctx context.Context, _a1 *user.User)) *MockUseCase_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*user.User))
	})
	return _c
}

func (_c *MockUseCase_CreateUser_Call) Return(_a0 error) *MockUseCase_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCase_CreateUser_Call) RunAndReturn(run func(context.Context, *user.User) error) *MockUseCase_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *MockUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*user.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *user.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCase_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type MockUseCase_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockUseCase_Expecter) GetUserByID(ctx interface{}, id interface{}) *MockUseCase_GetUserByID_Call {
	return &MockUseCase_GetUserByID_Call{Call: _e.mock.On("GetUserByID", ctx, id)}
}

func (_c *MockUseCase_GetUserByID_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockUseCase_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockUseCase_GetUserByID_Call) Return(_a0 *user.User, _a1 error) *MockUseCase_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCase_GetUserByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*user.User, error)) *MockUseCase_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUseCase creates a new instance of MockUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUseCase {
	mock := &MockUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
