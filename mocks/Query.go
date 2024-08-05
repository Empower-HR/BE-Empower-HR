// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	companies "be-empower-hr/features/Companies"

	mock "github.com/stretchr/testify/mock"
)

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

// GetCompany provides a mock function with given fields:
func (_m *Query) GetCompany() (companies.CompanyDataEntity, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCompany")
	}

	var r0 companies.CompanyDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func() (companies.CompanyDataEntity, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() companies.CompanyDataEntity); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(companies.CompanyDataEntity)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCompanyID provides a mock function with given fields: ID
func (_m *Query) GetCompanyID(ID uint) (companies.CompanyDataEntity, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for GetCompanyID")
	}

	var r0 companies.CompanyDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (companies.CompanyDataEntity, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(uint) companies.CompanyDataEntity); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(companies.CompanyDataEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCompany provides a mock function with given fields: ID, updateCompany
func (_m *Query) UpdateCompany(ID uint, updateCompany companies.CompanyDataEntity) error {
	ret := _m.Called(ID, updateCompany)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCompany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, companies.CompanyDataEntity) error); ok {
		r0 = rf(ID, updateCompany)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQuery creates a new instance of Query. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *Query {
	mock := &Query{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
