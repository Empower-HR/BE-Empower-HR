// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// MiddlewaresInterface is an autogenerated mock type for the MiddlewaresInterface type
type MiddlewaresInterface struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields: userId, companyId
func (_m *MiddlewaresInterface) CreateToken(userId int, companyId int) (string, error) {
	ret := _m.Called(userId, companyId)

	if len(ret) == 0 {
		panic("no return value specified for CreateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (string, error)); ok {
		return rf(userId, companyId)
	}
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(userId, companyId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userId, companyId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtractCompanyID provides a mock function with given fields: e
func (_m *MiddlewaresInterface) ExtractCompanyID(e echo.Context) (uint, error) {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for ExtractCompanyID")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context) (uint, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(echo.Context) uint); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(echo.Context) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtractTokenUserId provides a mock function with given fields: e
func (_m *MiddlewaresInterface) ExtractTokenUserId(e echo.Context) int {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for ExtractTokenUserId")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(echo.Context) int); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ExtractTokenUserRole provides a mock function with given fields: e
func (_m *MiddlewaresInterface) ExtractTokenUserRole(e echo.Context) string {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for ExtractTokenUserRole")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(echo.Context) string); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InvalidateToken provides a mock function with given fields: token
func (_m *MiddlewaresInterface) InvalidateToken(token string) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for InvalidateToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsTokenInvalidated provides a mock function with given fields: token
func (_m *MiddlewaresInterface) IsTokenInvalidated(token string) bool {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for IsTokenInvalidated")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JWTMiddleware provides a mock function with given fields:
func (_m *MiddlewaresInterface) JWTMiddleware() echo.MiddlewareFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JWTMiddleware")
	}

	var r0 echo.MiddlewareFunc
	if rf, ok := ret.Get(0).(func() echo.MiddlewareFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.MiddlewareFunc)
		}
	}

	return r0
}

// NewMiddlewaresInterface creates a new instance of MiddlewaresInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMiddlewaresInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MiddlewaresInterface {
	mock := &MiddlewaresInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
