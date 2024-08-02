// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	companies "be-empower-hr/features/Companies"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetCompany provides a mock function with given fields: ID
func (_m *Service) GetCompany(ID uint) (companies.CompanyDataEntity, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for GetCompany")
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
func (_m *Service) UpdateCompany(ID uint, updateCompany companies.CompanyDataEntity) error {
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

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
