// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	users "be-empower-hr/features/Users"

	mock "github.com/stretchr/testify/mock"
)

// DataUserInterface is an autogenerated mock type for the DataUserInterface type
type DataUserInterface struct {
	mock.Mock
}

// AccountByEmail provides a mock function with given fields: email
func (_m *DataUserInterface) AccountByEmail(email string) (*users.PersonalDataEntity, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for AccountByEmail")
	}

	var r0 *users.PersonalDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*users.PersonalDataEntity, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *users.PersonalDataEntity); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountById provides a mock function with given fields: userid
func (_m *DataUserInterface) AccountById(userid uint) (*users.PersonalDataEntity, error) {
	ret := _m.Called(userid)

	if len(ret) == 0 {
		panic("no return value specified for AccountById")
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

// CountContractUsers provides a mock function with given fields: companyID
func (_m *DataUserInterface) CountContractUsers(companyID uint) (int64, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for CountContractUsers")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(companyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountFemaleUsers provides a mock function with given fields: companyID
func (_m *DataUserInterface) CountFemaleUsers(companyID uint) (int64, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for CountFemaleUsers")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(companyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountMaleUsers provides a mock function with given fields: companyID
func (_m *DataUserInterface) CountMaleUsers(companyID uint) (int64, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for CountMaleUsers")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(companyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountPayrollUsers provides a mock function with given fields: companyID
func (_m *DataUserInterface) CountPayrollUsers(companyID uint) (int64, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for CountPayrollUsers")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(companyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountPermanentUsers provides a mock function with given fields: companyID
func (_m *DataUserInterface) CountPermanentUsers(companyID uint) (int64, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for CountPermanentUsers")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(companyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountTotalUsers provides a mock function with given fields: companyID
func (_m *DataUserInterface) CountTotalUsers(companyID uint) (int64, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for CountTotalUsers")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(companyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccountAdmin provides a mock function with given fields: account, companyName, department, jobPosition
func (_m *DataUserInterface) CreateAccountAdmin(account users.PersonalDataEntity, companyName string, department string, jobPosition string) (uint, uint, error) {
	ret := _m.Called(account, companyName, department, jobPosition)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccountAdmin")
	}

	var r0 uint
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(users.PersonalDataEntity, string, string, string) (uint, uint, error)); ok {
		return rf(account, companyName, department, jobPosition)
	}
	if rf, ok := ret.Get(0).(func(users.PersonalDataEntity, string, string, string) uint); ok {
		r0 = rf(account, companyName, department, jobPosition)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(users.PersonalDataEntity, string, string, string) uint); ok {
		r1 = rf(account, companyName, department, jobPosition)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(users.PersonalDataEntity, string, string, string) error); ok {
		r2 = rf(account, companyName, department, jobPosition)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateEmployment provides a mock function with given fields: personalID, addEmployment
func (_m *DataUserInterface) CreateEmployment(personalID uint, addEmployment users.EmploymentDataEntity) (uint, error) {
	ret := _m.Called(personalID, addEmployment)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmployment")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, users.EmploymentDataEntity) (uint, error)); ok {
		return rf(personalID, addEmployment)
	}
	if rf, ok := ret.Get(0).(func(uint, users.EmploymentDataEntity) uint); ok {
		r0 = rf(personalID, addEmployment)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(uint, users.EmploymentDataEntity) error); ok {
		r1 = rf(personalID, addEmployment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLeaves provides a mock function with given fields: PersonalID, addLeaves
func (_m *DataUserInterface) CreateLeaves(PersonalID uint, addLeaves users.LeavesDataEntity) (uint, error) {
	ret := _m.Called(PersonalID, addLeaves)

	if len(ret) == 0 {
		panic("no return value specified for CreateLeaves")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, users.LeavesDataEntity) (uint, error)); ok {
		return rf(PersonalID, addLeaves)
	}
	if rf, ok := ret.Get(0).(func(uint, users.LeavesDataEntity) uint); ok {
		r0 = rf(PersonalID, addLeaves)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(uint, users.LeavesDataEntity) error); ok {
		r1 = rf(PersonalID, addLeaves)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePayroll provides a mock function with given fields: employmentID, addPayroll
func (_m *DataUserInterface) CreatePayroll(employmentID uint, addPayroll users.PayrollDataEntity) error {
	ret := _m.Called(employmentID, addPayroll)

	if len(ret) == 0 {
		panic("no return value specified for CreatePayroll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.PayrollDataEntity) error); ok {
		r0 = rf(employmentID, addPayroll)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePersonal provides a mock function with given fields: CompanyID, addPersonal
func (_m *DataUserInterface) CreatePersonal(CompanyID uint, addPersonal users.PersonalDataEntity) (uint, error) {
	ret := _m.Called(CompanyID, addPersonal)

	if len(ret) == 0 {
		panic("no return value specified for CreatePersonal")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, users.PersonalDataEntity) (uint, error)); ok {
		return rf(CompanyID, addPersonal)
	}
	if rf, ok := ret.Get(0).(func(uint, users.PersonalDataEntity) uint); ok {
		r0 = rf(CompanyID, addPersonal)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(uint, users.PersonalDataEntity) error); ok {
		r1 = rf(CompanyID, addPersonal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Dashboard provides a mock function with given fields: userID, companyID
func (_m *DataUserInterface) Dashboard(userID uint, companyID uint) (*users.DashboardStats, error) {
	ret := _m.Called(userID, companyID)

	if len(ret) == 0 {
		panic("no return value specified for Dashboard")
	}

	var r0 *users.DashboardStats
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (*users.DashboardStats, error)); ok {
		return rf(userID, companyID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) *users.DashboardStats); ok {
		r0 = rf(userID, companyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.DashboardStats)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccountAdmin provides a mock function with given fields: userid
func (_m *DataUserInterface) DeleteAccountAdmin(userid uint) error {
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
func (_m *DataUserInterface) DeleteAccountEmployeeByAdmin(userid uint) error {
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

// GetAccountByJobLevel provides a mock function with given fields: jobLevel
func (_m *DataUserInterface) GetAccountByJobLevel(jobLevel string) ([]users.PersonalDataEntity, error) {
	ret := _m.Called(jobLevel)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountByJobLevel")
	}

	var r0 []users.PersonalDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]users.PersonalDataEntity, error)); ok {
		return rf(jobLevel)
	}
	if rf, ok := ret.Get(0).(func(string) []users.PersonalDataEntity); ok {
		r0 = rf(jobLevel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jobLevel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountByName provides a mock function with given fields: accountName
func (_m *DataUserInterface) GetAccountByName(accountName string) ([]users.PersonalDataEntity, error) {
	ret := _m.Called(accountName)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountByName")
	}

	var r0 []users.PersonalDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]users.PersonalDataEntity, error)); ok {
		return rf(accountName)
	}
	if rf, ok := ret.Get(0).(func(string) []users.PersonalDataEntity); ok {
		r0 = rf(accountName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: page, pageSize, companyID
func (_m *DataUserInterface) GetAll(page int, pageSize int, companyID uint) ([]users.PersonalDataEntity, error) {
	ret := _m.Called(page, pageSize, companyID)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []users.PersonalDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, uint) ([]users.PersonalDataEntity, error)); ok {
		return rf(page, pageSize, companyID)
	}
	if rf, ok := ret.Get(0).(func(int, int, uint) []users.PersonalDataEntity); ok {
		r0 = rf(page, pageSize, companyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.PersonalDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, uint) error); ok {
		r1 = rf(page, pageSize, companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCompanyIDByName provides a mock function with given fields: companyName
func (_m *DataUserInterface) GetCompanyIDByName(companyName string) (uint, error) {
	ret := _m.Called(companyName)

	if len(ret) == 0 {
		panic("no return value specified for GetCompanyIDByName")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (uint, error)); ok {
		return rf(companyName)
	}
	if rf, ok := ret.Get(0).(func(string) uint); ok {
		r0 = rf(companyName)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(companyName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountAdmins provides a mock function with given fields: userid, account
func (_m *DataUserInterface) UpdateAccountAdmins(userid uint, account users.PersonalDataEntity) error {
	ret := _m.Called(userid, account)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountAdmins")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.PersonalDataEntity) error); ok {
		r0 = rf(userid, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAccountEmployees provides a mock function with given fields: userid, account
func (_m *DataUserInterface) UpdateAccountEmployees(userid uint, account users.PersonalDataEntity) error {
	ret := _m.Called(userid, account)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountEmployees")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.PersonalDataEntity) error); ok {
		r0 = rf(userid, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEmploymentEmployee provides a mock function with given fields: ID, employeID, updateEmploymentEmployee
func (_m *DataUserInterface) UpdateEmploymentEmployee(ID uint, employeID uint, updateEmploymentEmployee users.EmploymentDataEntity) error {
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

// UpdateProfileEmployments provides a mock function with given fields: userid, accounts
func (_m *DataUserInterface) UpdateProfileEmployments(userid uint, accounts users.EmploymentDataEntity) error {
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

// NewDataUserInterface creates a new instance of DataUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataUserInterface {
	mock := &DataUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
