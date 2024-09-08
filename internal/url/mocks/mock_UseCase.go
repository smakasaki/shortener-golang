// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	url "github.com/smakasaki/shortener/internal/url"
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

// Create provides a mock function with given fields: ctx, userID, originalURL
func (_m *MockUseCase) Create(ctx context.Context, userID *uuid.UUID, originalURL string) (*url.URL, error) {
	ret := _m.Called(ctx, userID, originalURL)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, string) (*url.URL, error)); ok {
		return rf(ctx, userID, originalURL)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, string) *url.URL); ok {
		r0 = rf(ctx, userID, originalURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uuid.UUID, string) error); ok {
		r1 = rf(ctx, userID, originalURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCase_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUseCase_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - userID *uuid.UUID
//   - originalURL string
func (_e *MockUseCase_Expecter) Create(ctx interface{}, userID interface{}, originalURL interface{}) *MockUseCase_Create_Call {
	return &MockUseCase_Create_Call{Call: _e.mock.On("Create", ctx, userID, originalURL)}
}

func (_c *MockUseCase_Create_Call) Run(run func(ctx context.Context, userID *uuid.UUID, originalURL string)) *MockUseCase_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *MockUseCase_Create_Call) Return(_a0 *url.URL, _a1 error) *MockUseCase_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCase_Create_Call) RunAndReturn(run func(context.Context, *uuid.UUID, string) (*url.URL, error)) *MockUseCase_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, shortCode, userID
func (_m *MockUseCase) Delete(ctx context.Context, shortCode string, userID *uuid.UUID) error {
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

// MockUseCase_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUseCase_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
//   - userID *uuid.UUID
func (_e *MockUseCase_Expecter) Delete(ctx interface{}, shortCode interface{}, userID interface{}) *MockUseCase_Delete_Call {
	return &MockUseCase_Delete_Call{Call: _e.mock.On("Delete", ctx, shortCode, userID)}
}

func (_c *MockUseCase_Delete_Call) Run(run func(ctx context.Context, shortCode string, userID *uuid.UUID)) *MockUseCase_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*uuid.UUID))
	})
	return _c
}

func (_c *MockUseCase_Delete_Call) Return(_a0 error) *MockUseCase_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCase_Delete_Call) RunAndReturn(run func(context.Context, string, *uuid.UUID) error) *MockUseCase_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx, userID, limit, offset
func (_m *MockUseCase) GetAll(ctx context.Context, userID *uuid.UUID, limit int, offset int) ([]*url.URL, error) {
	ret := _m.Called(ctx, userID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
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

// MockUseCase_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockUseCase_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
//   - userID *uuid.UUID
//   - limit int
//   - offset int
func (_e *MockUseCase_Expecter) GetAll(ctx interface{}, userID interface{}, limit interface{}, offset interface{}) *MockUseCase_GetAll_Call {
	return &MockUseCase_GetAll_Call{Call: _e.mock.On("GetAll", ctx, userID, limit, offset)}
}

func (_c *MockUseCase_GetAll_Call) Run(run func(ctx context.Context, userID *uuid.UUID, limit int, offset int)) *MockUseCase_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*uuid.UUID), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *MockUseCase_GetAll_Call) Return(_a0 []*url.URL, _a1 error) *MockUseCase_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCase_GetAll_Call) RunAndReturn(run func(context.Context, *uuid.UUID, int, int) ([]*url.URL, error)) *MockUseCase_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByShortCode provides a mock function with given fields: ctx, shortCode
func (_m *MockUseCase) GetByShortCode(ctx context.Context, shortCode string) (*url.URL, error) {
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

// MockUseCase_GetByShortCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByShortCode'
type MockUseCase_GetByShortCode_Call struct {
	*mock.Call
}

// GetByShortCode is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
func (_e *MockUseCase_Expecter) GetByShortCode(ctx interface{}, shortCode interface{}) *MockUseCase_GetByShortCode_Call {
	return &MockUseCase_GetByShortCode_Call{Call: _e.mock.On("GetByShortCode", ctx, shortCode)}
}

func (_c *MockUseCase_GetByShortCode_Call) Run(run func(ctx context.Context, shortCode string)) *MockUseCase_GetByShortCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUseCase_GetByShortCode_Call) Return(_a0 *url.URL, _a1 error) *MockUseCase_GetByShortCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCase_GetByShortCode_Call) RunAndReturn(run func(context.Context, string) (*url.URL, error)) *MockUseCase_GetByShortCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetStats provides a mock function with given fields: ctx, shortCode, userID
func (_m *MockUseCase) GetStats(ctx context.Context, shortCode string, userID *uuid.UUID) (*url.URLStats, error) {
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

// MockUseCase_GetStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStats'
type MockUseCase_GetStats_Call struct {
	*mock.Call
}

// GetStats is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
//   - userID *uuid.UUID
func (_e *MockUseCase_Expecter) GetStats(ctx interface{}, shortCode interface{}, userID interface{}) *MockUseCase_GetStats_Call {
	return &MockUseCase_GetStats_Call{Call: _e.mock.On("GetStats", ctx, shortCode, userID)}
}

func (_c *MockUseCase_GetStats_Call) Run(run func(ctx context.Context, shortCode string, userID *uuid.UUID)) *MockUseCase_GetStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*uuid.UUID))
	})
	return _c
}

func (_c *MockUseCase_GetStats_Call) Return(_a0 *url.URLStats, _a1 error) *MockUseCase_GetStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCase_GetStats_Call) RunAndReturn(run func(context.Context, string, *uuid.UUID) (*url.URLStats, error)) *MockUseCase_GetStats_Call {
	_c.Call.Return(run)
	return _c
}

// IncrementClickCount provides a mock function with given fields: ctx, urlID, ipAddress, userAgent, referer
func (_m *MockUseCase) IncrementClickCount(ctx context.Context, urlID int, ipAddress string, userAgent string, referer string) error {
	ret := _m.Called(ctx, urlID, ipAddress, userAgent, referer)

	if len(ret) == 0 {
		panic("no return value specified for IncrementClickCount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) error); ok {
		r0 = rf(ctx, urlID, ipAddress, userAgent, referer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUseCase_IncrementClickCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IncrementClickCount'
type MockUseCase_IncrementClickCount_Call struct {
	*mock.Call
}

// IncrementClickCount is a helper method to define mock.On call
//   - ctx context.Context
//   - urlID int
//   - ipAddress string
//   - userAgent string
//   - referer string
func (_e *MockUseCase_Expecter) IncrementClickCount(ctx interface{}, urlID interface{}, ipAddress interface{}, userAgent interface{}, referer interface{}) *MockUseCase_IncrementClickCount_Call {
	return &MockUseCase_IncrementClickCount_Call{Call: _e.mock.On("IncrementClickCount", ctx, urlID, ipAddress, userAgent, referer)}
}

func (_c *MockUseCase_IncrementClickCount_Call) Run(run func(ctx context.Context, urlID int, ipAddress string, userAgent string, referer string)) *MockUseCase_IncrementClickCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockUseCase_IncrementClickCount_Call) Return(_a0 error) *MockUseCase_IncrementClickCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCase_IncrementClickCount_Call) RunAndReturn(run func(context.Context, int, string, string, string) error) *MockUseCase_IncrementClickCount_Call {
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
