// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	leaves "be-empower-hr/features/Leaves"

	mock "github.com/stretchr/testify/mock"
)

// ServiceLeavesInterface is an autogenerated mock type for the ServiceLeavesInterface type
type ServiceLeavesInterface struct {
	mock.Mock
}

// Dashboard provides a mock function with given fields: companyID
func (_m *ServiceLeavesInterface) Dashboard(companyID uint) (*leaves.DashboardLeavesStats, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for Dashboard")
	}

	var r0 *leaves.DashboardLeavesStats
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*leaves.DashboardLeavesStats, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) *leaves.DashboardLeavesStats); ok {
		r0 = rf(companyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*leaves.DashboardLeavesStats)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DashboardEmployees provides a mock function with given fields: companyID, page, pageSize
func (_m *ServiceLeavesInterface) DashboardEmployees(companyID uint, page int, pageSize int) (*leaves.DashboardStats, error) {
	ret := _m.Called(companyID, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for DashboardEmployees")
	}

	var r0 *leaves.DashboardStats
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, int, int) (*leaves.DashboardStats, error)); ok {
		return rf(companyID, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int) *leaves.DashboardStats); ok {
		r0 = rf(companyID, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*leaves.DashboardStats)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int) error); ok {
		r1 = rf(companyID, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeavesByID provides a mock function with given fields: leaveID
func (_m *ServiceLeavesInterface) GetLeavesByID(leaveID uint) (*leaves.LeavesDataEntity, error) {
	ret := _m.Called(leaveID)

	if len(ret) == 0 {
		panic("no return value specified for GetLeavesByID")
	}

	var r0 *leaves.LeavesDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*leaves.LeavesDataEntity, error)); ok {
		return rf(leaveID)
	}
	if rf, ok := ret.Get(0).(func(uint) *leaves.LeavesDataEntity); ok {
		r0 = rf(leaveID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*leaves.LeavesDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(leaveID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestLeave provides a mock function with given fields: userID, leave
func (_m *ServiceLeavesInterface) RequestLeave(userID uint, leave leaves.LeavesDataEntity) error {
	ret := _m.Called(userID, leave)

	if len(ret) == 0 {
		panic("no return value specified for RequestLeave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, leaves.LeavesDataEntity) error); ok {
		r0 = rf(userID, leave)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLeaveStatus provides a mock function with given fields: userID, leaveID, updatesleaves
func (_m *ServiceLeavesInterface) UpdateLeaveStatus(userID uint, leaveID uint, updatesleaves leaves.LeavesDataEntity) error {
	ret := _m.Called(userID, leaveID, updatesleaves)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLeaveStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, leaves.LeavesDataEntity) error); ok {
		r0 = rf(userID, leaveID, updatesleaves)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ViewLeaveHistory provides a mock function with given fields: companyID, personalDataID, page, pageSize, status, startDate, endDate
func (_m *ServiceLeavesInterface) ViewLeaveHistory(companyID uint, personalDataID uint, page int, pageSize int, status string, startDate string, endDate string) ([]leaves.LeavesDataEntity, error) {
	ret := _m.Called(companyID, personalDataID, page, pageSize, status, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for ViewLeaveHistory")
	}

	var r0 []leaves.LeavesDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, int, int, string, string, string) ([]leaves.LeavesDataEntity, error)); ok {
		return rf(companyID, personalDataID, page, pageSize, status, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, int, int, string, string, string) []leaves.LeavesDataEntity); ok {
		r0 = rf(companyID, personalDataID, page, pageSize, status, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]leaves.LeavesDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, int, int, string, string, string) error); ok {
		r1 = rf(companyID, personalDataID, page, pageSize, status, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ViewLeaveHistoryEmployee provides a mock function with given fields: personalDataID, page, pageSize, status, startDate, endDate
func (_m *ServiceLeavesInterface) ViewLeaveHistoryEmployee(personalDataID uint, page int, pageSize int, status string, startDate string, endDate string) ([]leaves.LeavesDataEntity, error) {
	ret := _m.Called(personalDataID, page, pageSize, status, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for ViewLeaveHistoryEmployee")
	}

	var r0 []leaves.LeavesDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, int, int, string, string, string) ([]leaves.LeavesDataEntity, error)); ok {
		return rf(personalDataID, page, pageSize, status, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int, string, string, string) []leaves.LeavesDataEntity); ok {
		r0 = rf(personalDataID, page, pageSize, status, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]leaves.LeavesDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int, string, string, string) error); ok {
		r1 = rf(personalDataID, page, pageSize, status, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceLeavesInterface creates a new instance of ServiceLeavesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceLeavesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceLeavesInterface {
	mock := &ServiceLeavesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
