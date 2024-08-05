// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	users "be-empower-hr/features/Users"

	mock "github.com/stretchr/testify/mock"
)

// ServiceUserInterface is an autogenerated mock type for the ServiceUserInterface type
type ServiceUserInterface struct {
	mock.Mock
}

// CreateNewEmployee provides a mock function with given fields: addPersonal, addEmployment, addPayroll, addLeaves
func (_m *ServiceUserInterface) CreateNewEmployee(addPersonal users.PersonalDataEntity, addEmployment users.EmploymentDataEntity, addPayroll users.PayrollDataEntity, addLeaves users.LeavesDataEntity) error {
	ret := _m.Called(addPersonal, addEmployment, addPayroll, addLeaves)

	if len(ret) == 0 {
		panic("no return value specified for CreateNewEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(users.PersonalDataEntity, users.EmploymentDataEntity, users.PayrollDataEntity, users.LeavesDataEntity) error); ok {
		r0 = rf(addPersonal, addEmployment, addPayroll, addLeaves)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dashboard provides a mock function with given fields: companyID
func (_m *ServiceUserInterface) Dashboard(companyID uint) (*users.DashboardStats, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for Dashboard")
	}

	var r0 *users.DashboardStats
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*users.DashboardStats, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) *users.DashboardStats); ok {
		r0 = rf(companyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.DashboardStats)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccountAdmin provides a mock function with given fields: userid
func (_m *ServiceUserInterface) DeleteAccountAdmin(userid uint) error {
	ret := _m.Called(userid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAccountAdmin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAccountEmployeeByAdmin provides a mock function with given fields: userid
func (_m *ServiceUserInterface) DeleteAccountEmployeeByAdmin(userid uint) error {
	ret := _m.Called(userid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAccountEmployeeByAdmin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllAccount provides a mock function with given fields: name, jobLevel, page, pageSize
func (_m *ServiceUserInterface) GetAllAccount(name string, jobLevel string, page int, pageSize int) ([]users.PersonalDataEntity, error) {
	ret := _m.Called(name, jobLevel, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAccount")
	}

	var r0 []users.PersonalDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]users.PersonalDataEntity, error)); ok {
		return rf(name, jobLevel, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []users.PersonalDataEntity); ok {
		r0 = rf(name, jobLevel, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(name, jobLevel, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfile provides a mock function with given fields: userid
func (_m *ServiceUserInterface) GetProfile(userid uint) (*users.PersonalDataEntity, error) {
	ret := _m.Called(userid)

	if len(ret) == 0 {
		panic("no return value specified for GetProfile")
	}

	var r0 *users.PersonalDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*users.PersonalDataEntity, error)); ok {
		return rf(userid)
	}
	if rf, ok := ret.Get(0).(func(uint) *users.PersonalDataEntity); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfileById provides a mock function with given fields: userid
func (_m *ServiceUserInterface) GetProfileById(userid uint) (*users.PersonalDataEntity, error) {
	ret := _m.Called(userid)

	if len(ret) == 0 {
		panic("no return value specified for GetProfileById")
	}

	var r0 *users.PersonalDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*users.PersonalDataEntity, error)); ok {
		return rf(userid)
	}
	if rf, ok := ret.Get(0).(func(uint) *users.PersonalDataEntity); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginAccount provides a mock function with given fields: email, password
func (_m *ServiceUserInterface) LoginAccount(email string, password string) (*users.PersonalDataEntity, string, error) {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for LoginAccount")
	}

	var r0 *users.PersonalDataEntity
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*users.PersonalDataEntity, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *users.PersonalDataEntity); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegistrasiAccountAdmin provides a mock function with given fields: accounts, companyName, department, jobPosition
func (_m *ServiceUserInterface) RegistrasiAccountAdmin(accounts users.PersonalDataEntity, companyName string, department string, jobPosition string) (uint, uint, error) {
	ret := _m.Called(accounts, companyName, department, jobPosition)

	if len(ret) == 0 {
		panic("no return value specified for RegistrasiAccountAdmin")
	}

	var r0 uint
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(users.PersonalDataEntity, string, string, string) (uint, uint, error)); ok {
		return rf(accounts, companyName, department, jobPosition)
	}
	if rf, ok := ret.Get(0).(func(users.PersonalDataEntity, string, string, string) uint); ok {
		r0 = rf(accounts, companyName, department, jobPosition)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(users.PersonalDataEntity, string, string, string) uint); ok {
		r1 = rf(accounts, companyName, department, jobPosition)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(users.PersonalDataEntity, string, string, string) error); ok {
		r2 = rf(accounts, companyName, department, jobPosition)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateEmploymentEmployee provides a mock function with given fields: ID, employeID, updateEmploymentEmployee
func (_m *ServiceUserInterface) UpdateEmploymentEmployee(ID uint, employeID uint, updateEmploymentEmployee users.EmploymentDataEntity) error {
	ret := _m.Called(ID, employeID, updateEmploymentEmployee)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmploymentEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, users.EmploymentDataEntity) error); ok {
		r0 = rf(ID, employeID, updateEmploymentEmployee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfileAdmins provides a mock function with given fields: userid, accounts
func (_m *ServiceUserInterface) UpdateProfileAdmins(userid uint, accounts users.PersonalDataEntity) error {
	ret := _m.Called(userid, accounts)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProfileAdmins")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.PersonalDataEntity) error); ok {
		r0 = rf(userid, accounts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfileEmployees provides a mock function with given fields: userid, accounts
func (_m *ServiceUserInterface) UpdateProfileEmployees(userid uint, accounts users.PersonalDataEntity) error {
	ret := _m.Called(userid, accounts)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProfileEmployees")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.PersonalDataEntity) error); ok {
		r0 = rf(userid, accounts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfileEmployments provides a mock function with given fields: userid, accounts
func (_m *ServiceUserInterface) UpdateProfileEmployments(userid uint, accounts users.EmploymentDataEntity) error {
	ret := _m.Called(userid, accounts)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProfileEmployments")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.EmploymentDataEntity) error); ok {
		r0 = rf(userid, accounts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServiceUserInterface creates a new instance of ServiceUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceUserInterface {
	mock := &ServiceUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
