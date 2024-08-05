// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	attendance "be-empower-hr/features/Attendance"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AQuery is an autogenerated mock type for the AQuery type
type AQuery struct {
	mock.Mock
}

// CountAttByIdPersonAndSearch provides a mock function with given fields: personID, search
func (_m *AQuery) CountAttByIdPersonAndSearch(personID uint, search string) (int64, error) {
	ret := _m.Called(personID, search)

	if len(ret) == 0 {
		panic("no return value specified for CountAttByIdPersonAndSearch")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string) (int64, error)); ok {
		return rf(personID, search)
	}
	if rf, ok := ret.Get(0).(func(uint, string) int64); ok {
		r0 = rf(personID, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(personID, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountAttBySearch provides a mock function with given fields: search
func (_m *AQuery) CountAttBySearch(search string) (int64, error) {
	ret := _m.Called(search)

	if len(ret) == 0 {
		panic("no return value specified for CountAttBySearch")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(search)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: newAtt
func (_m *AQuery) Create(newAtt attendance.Attandance) error {
	ret := _m.Called(newAtt)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(attendance.Attandance) error); ok {
		r0 = rf(newAtt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAttbyId provides a mock function with given fields: attId
func (_m *AQuery) DeleteAttbyId(attId uint) error {
	ret := _m.Called(attId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAttbyId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(attId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllAttDownload provides a mock function with given fields:
func (_m *AQuery) GetAllAttDownload() ([]attendance.Attandance, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllAttDownload")
	}

	var r0 []attendance.Attandance
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]attendance.Attandance, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []attendance.Attandance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.Attandance)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAttbyDate provides a mock function with given fields: date, limit, offset
func (_m *AQuery) GetAllAttbyDate(date int, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	ret := _m.Called(date, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAttbyDate")
	}

	var r0 []attendance.AttendanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, int) ([]attendance.AttendanceDetail, error)); ok {
		return rf(date, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(int, int, int) []attendance.AttendanceDetail); ok {
		r0 = rf(date, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.AttendanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(date, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAttbyDateandPerson provides a mock function with given fields: perseonID, date, limit, offset
func (_m *AQuery) GetAllAttbyDateandPerson(perseonID uint, date int, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	ret := _m.Called(perseonID, date, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAttbyDateandPerson")
	}

	var r0 []attendance.AttendanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, int, int, int) ([]attendance.AttendanceDetail, error)); ok {
		return rf(perseonID, date, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int, int) []attendance.AttendanceDetail); ok {
		r0 = rf(perseonID, date, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.AttendanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int, int) error); ok {
		r1 = rf(perseonID, date, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAttbyIdPersonAndStatus provides a mock function with given fields: id, status, limit, offset
func (_m *AQuery) GetAllAttbyIdPersonAndStatus(id uint, status string, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	ret := _m.Called(id, status, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAttbyIdPersonAndStatus")
	}

	var r0 []attendance.AttendanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string, int, int) ([]attendance.AttendanceDetail, error)); ok {
		return rf(id, status, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(uint, string, int, int) []attendance.AttendanceDetail); ok {
		r0 = rf(id, status, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.AttendanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string, int, int) error); ok {
		r1 = rf(id, status, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAttbyStatus provides a mock function with given fields: status, limit, offset
func (_m *AQuery) GetAllAttbyStatus(status string, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	ret := _m.Called(status, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAttbyStatus")
	}

	var r0 []attendance.AttendanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]attendance.AttendanceDetail, error)); ok {
		return rf(status, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []attendance.AttendanceDetail); ok {
		r0 = rf(status, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.AttendanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(status, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttByIdAtt provides a mock function with given fields: idAtt
func (_m *AQuery) GetAttByIdAtt(idAtt uint) ([]attendance.AttendanceDetail, error) {
	ret := _m.Called(idAtt)

	if len(ret) == 0 {
		panic("no return value specified for GetAttByIdAtt")
	}

	var r0 []attendance.AttendanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]attendance.AttendanceDetail, error)); ok {
		return rf(idAtt)
	}
	if rf, ok := ret.Get(0).(func(uint) []attendance.AttendanceDetail); ok {
		r0 = rf(idAtt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.AttendanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idAtt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttByPersonalID provides a mock function with given fields: personalID, term, limit, offset
func (_m *AQuery) GetAttByPersonalID(personalID uint, term string, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	ret := _m.Called(personalID, term, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAttByPersonalID")
	}

	var r0 []attendance.AttendanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string, int, int) ([]attendance.AttendanceDetail, error)); ok {
		return rf(personalID, term, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(uint, string, int, int) []attendance.AttendanceDetail); ok {
		r0 = rf(personalID, term, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.AttendanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string, int, int) error); ok {
		r1 = rf(personalID, term, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttendanceDetails provides a mock function with given fields: searchTerm, limit, offset
func (_m *AQuery) GetAttendanceDetails(searchTerm string, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	ret := _m.Called(searchTerm, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAttendanceDetails")
	}

	var r0 []attendance.AttendanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]attendance.AttendanceDetail, error)); ok {
		return rf(searchTerm, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []attendance.AttendanceDetail); ok {
		r0 = rf(searchTerm, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.AttendanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(searchTerm, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCompany provides a mock function with given fields: idPerson
func (_m *AQuery) GetCompany(idPerson uint) ([]attendance.CompanyDataEntity, error) {
	ret := _m.Called(idPerson)

	if len(ret) == 0 {
		panic("no return value specified for GetCompany")
	}

	var r0 []attendance.CompanyDataEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]attendance.CompanyDataEntity, error)); ok {
		return rf(idPerson)
	}
	if rf, ok := ret.Get(0).(func(uint) []attendance.CompanyDataEntity); ok {
		r0 = rf(idPerson)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.CompanyDataEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idPerson)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAttendancesCount provides a mock function with given fields:
func (_m *AQuery) GetTotalAttendancesCount() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotalAttendancesCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAttendancesCountByStatus provides a mock function with given fields: status
func (_m *AQuery) GetTotalAttendancesCountByStatus(status string) (int64, error) {
	ret := _m.Called(status)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalAttendancesCountByStatus")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAttendancesCountByStatusandPerson provides a mock function with given fields: status, personID
func (_m *AQuery) GetTotalAttendancesCountByStatusandPerson(status string, personID uint) (int64, error) {
	ret := _m.Called(status, personID)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalAttendancesCountByStatusandPerson")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint) (int64, error)); ok {
		return rf(status, personID)
	}
	if rf, ok := ret.Get(0).(func(string, uint) int64); ok {
		r0 = rf(status, personID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, uint) error); ok {
		r1 = rf(status, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAttendancesCountbyDate provides a mock function with given fields: date
func (_m *AQuery) GetTotalAttendancesCountbyDate(date int) (int64, error) {
	ret := _m.Called(date)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalAttendancesCountbyDate")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int64, error)); ok {
		return rf(date)
	}
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(date)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAttendancesCountbyDateandPerson provides a mock function with given fields: date, personID
func (_m *AQuery) GetTotalAttendancesCountbyDateandPerson(date int, personID uint) (int64, error) {
	ret := _m.Called(date, personID)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalAttendancesCountbyDateandPerson")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int, uint) (int64, error)); ok {
		return rf(date, personID)
	}
	if rf, ok := ret.Get(0).(func(int, uint) int64); ok {
		r0 = rf(date, personID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int, uint) error); ok {
		r1 = rf(date, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAttendancesCountbyPerson provides a mock function with given fields: personID
func (_m *AQuery) GetTotalAttendancesCountbyPerson(personID uint) (int64, error) {
	ret := _m.Called(personID)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalAttendancesCountbyPerson")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(personID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(personID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsDateExists provides a mock function with given fields: personalID, date
func (_m *AQuery) IsDateExists(personalID uint, date time.Time) (bool, error) {
	ret := _m.Called(personalID, date)

	if len(ret) == 0 {
		panic("no return value specified for IsDateExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, time.Time) (bool, error)); ok {
		return rf(personalID, date)
	}
	if rf, ok := ret.Get(0).(func(uint, time.Time) bool); ok {
		r0 = rf(personalID, date)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(uint, time.Time) error); ok {
		r1 = rf(personalID, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updatedAtt
func (_m *AQuery) Update(id uint, updatedAtt attendance.Attandance) error {
	ret := _m.Called(id, updatedAtt)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, attendance.Attandance) error); ok {
		r0 = rf(id, updatedAtt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAQuery creates a new instance of AQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *AQuery {
	mock := &AQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
