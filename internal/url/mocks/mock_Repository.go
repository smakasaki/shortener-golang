// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	url "github.com/smakasaki/shortener/internal/url"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *MockRepository) Create(ctx context.Context, _a1 *url.URL) (*url.URL, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL) (*url.URL, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL) *url.URL); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *url.URL) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *url.URL
func (_e *MockRepository_Expecter) Create(ctx interface{}, _a1 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", ctx, _a1)}
}

func (_c *MockRepository_Create_Call) Run(run func(ctx context.Context, _a1 *url.URL)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*url.URL))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 *url.URL, _a1 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, *url.URL) (*url.URL, error)) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateClick provides a mock function with given fields: ctx, urlID, ipAddress, userAgent, referer
func (_m *MockRepository) CreateClick(ctx context.Context, urlID int, ipAddress string, userAgent string, referer string) error {
	ret := _m.Called(ctx, urlID, ipAddress, userAgent, referer)

	if len(ret) == 0 {
		panic("no return value specified for CreateClick")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) error); ok {
		r0 = rf(ctx, urlID, ipAddress, userAgent, referer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_CreateClick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateClick'
type MockRepository_CreateClick_Call struct {
	*mock.Call
}

// CreateClick is a helper method to define mock.On call
//   - ctx context.Context
//   - urlID int
//   - ipAddress string
//   - userAgent string
//   - referer string
func (_e *MockRepository_Expecter) CreateClick(ctx interface{}, urlID interface{}, ipAddress interface{}, userAgent interface{}, referer interface{}) *MockRepository_CreateClick_Call {
	return &MockRepository_CreateClick_Call{Call: _e.mock.On("CreateClick", ctx, urlID, ipAddress, userAgent, referer)}
}

func (_c *MockRepository_CreateClick_Call) Run(run func(ctx context.Context, urlID int, ipAddress string, userAgent string, referer string)) *MockRepository_CreateClick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockRepository_CreateClick_Call) Return(_a0 error) *MockRepository_CreateClick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_CreateClick_Call) RunAndReturn(run func(context.Context, int, string, string, string) error) *MockRepository_CreateClick_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, shortCode, userID
func (_m *MockRepository) Delete(ctx context.Context, shortCode string, userID *uuid.UUID) error {
	ret := _m.Called(ctx, shortCode, userID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) error); ok {
		r0 = rf(ctx, shortCode, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
//   - userID *uuid.UUID
func (_e *MockRepository_Expecter) Delete(ctx interface{}, shortCode interface{}, userID interface{}) *MockRepository_Delete_Call {
	return &MockRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, shortCode, userID)}
}

func (_c *MockRepository_Delete_Call) Run(run func(ctx context.Context, shortCode string, userID *uuid.UUID)) *MockRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*uuid.UUID))
	})
	return _c
}

func (_c *MockRepository_Delete_Call) Return(_a0 error) *MockRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Delete_Call) RunAndReturn(run func(context.Context, string, *uuid.UUID) error) *MockRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByUser provides a mock function with given fields: ctx, userID, limit, offset
func (_m *MockRepository) GetAllByUser(ctx context.Context, userID *uuid.UUID, limit int, offset int) ([]*url.URL, error) {
	ret := _m.Called(ctx, userID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByUser")
	}

	var r0 []*url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, int, int) ([]*url.URL, error)); ok {
		return rf(ctx, userID, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, int, int) []*url.URL); ok {
		r0 = rf(ctx, userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uuid.UUID, int, int) error); ok {
		r1 = rf(ctx, userID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetAllByUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByUser'
type MockRepository_GetAllByUser_Call struct {
	*mock.Call
}

// GetAllByUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userID *uuid.UUID
//   - limit int
//   - offset int
func (_e *MockRepository_Expecter) GetAllByUser(ctx interface{}, userID interface{}, limit interface{}, offset interface{}) *MockRepository_GetAllByUser_Call {
	return &MockRepository_GetAllByUser_Call{Call: _e.mock.On("GetAllByUser", ctx, userID, limit, offset)}
}

func (_c *MockRepository_GetAllByUser_Call) Run(run func(ctx context.Context, userID *uuid.UUID, limit int, offset int)) *MockRepository_GetAllByUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*uuid.UUID), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *MockRepository_GetAllByUser_Call) Return(_a0 []*url.URL, _a1 error) *MockRepository_GetAllByUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetAllByUser_Call) RunAndReturn(run func(context.Context, *uuid.UUID, int, int) ([]*url.URL, error)) *MockRepository_GetAllByUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetByShortCode provides a mock function with given fields: ctx, shortCode
func (_m *MockRepository) GetByShortCode(ctx context.Context, shortCode string) (*url.URL, error) {
	ret := _m.Called(ctx, shortCode)

	if len(ret) == 0 {
		panic("no return value specified for GetByShortCode")
	}

	var r0 *url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*url.URL, error)); ok {
		return rf(ctx, shortCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *url.URL); ok {
		r0 = rf(ctx, shortCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetByShortCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByShortCode'
type MockRepository_GetByShortCode_Call struct {
	*mock.Call
}

// GetByShortCode is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
func (_e *MockRepository_Expecter) GetByShortCode(ctx interface{}, shortCode interface{}) *MockRepository_GetByShortCode_Call {
	return &MockRepository_GetByShortCode_Call{Call: _e.mock.On("GetByShortCode", ctx, shortCode)}
}

func (_c *MockRepository_GetByShortCode_Call) Run(run func(ctx context.Context, shortCode string)) *MockRepository_GetByShortCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetByShortCode_Call) Return(_a0 *url.URL, _a1 error) *MockRepository_GetByShortCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetByShortCode_Call) RunAndReturn(run func(context.Context, string) (*url.URL, error)) *MockRepository_GetByShortCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetStats provides a mock function with given fields: ctx, shortCode, userID
func (_m *MockRepository) GetStats(ctx context.Context, shortCode string, userID *uuid.UUID) (*url.URLStats, error) {
	ret := _m.Called(ctx, shortCode, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetStats")
	}

	var r0 *url.URLStats
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) (*url.URLStats, error)); ok {
		return rf(ctx, shortCode, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) *url.URLStats); ok {
		r0 = rf(ctx, shortCode, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URLStats)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *uuid.UUID) error); ok {
		r1 = rf(ctx, shortCode, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStats'
type MockRepository_GetStats_Call struct {
	*mock.Call
}

// GetStats is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
//   - userID *uuid.UUID
func (_e *MockRepository_Expecter) GetStats(ctx interface{}, shortCode interface{}, userID interface{}) *MockRepository_GetStats_Call {
	return &MockRepository_GetStats_Call{Call: _e.mock.On("GetStats", ctx, shortCode, userID)}
}

func (_c *MockRepository_GetStats_Call) Run(run func(ctx context.Context, shortCode string, userID *uuid.UUID)) *MockRepository_GetStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*uuid.UUID))
	})
	return _c
}

func (_c *MockRepository_GetStats_Call) Return(_a0 *url.URLStats, _a1 error) *MockRepository_GetStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetStats_Call) RunAndReturn(run func(context.Context, string, *uuid.UUID) (*url.URLStats, error)) *MockRepository_GetStats_Call {
	_c.Call.Return(run)
	return _c
}

// IncrementClick provides a mock function with given fields: ctx, urlID
func (_m *MockRepository) IncrementClick(ctx context.Context, urlID int) error {
	ret := _m.Called(ctx, urlID)

	if len(ret) == 0 {
		panic("no return value specified for IncrementClick")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, urlID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_IncrementClick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IncrementClick'
type MockRepository_IncrementClick_Call struct {
	*mock.Call
}

// IncrementClick is a helper method to define mock.On call
//   - ctx context.Context
//   - urlID int
func (_e *MockRepository_Expecter) IncrementClick(ctx interface{}, urlID interface{}) *MockRepository_IncrementClick_Call {
	return &MockRepository_IncrementClick_Call{Call: _e.mock.On("IncrementClick", ctx, urlID)}
}

func (_c *MockRepository_IncrementClick_Call) Run(run func(ctx context.Context, urlID int)) *MockRepository_IncrementClick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepository_IncrementClick_Call) Return(_a0 error) *MockRepository_IncrementClick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_IncrementClick_Call) RunAndReturn(run func(context.Context, int) error) *MockRepository_IncrementClick_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
