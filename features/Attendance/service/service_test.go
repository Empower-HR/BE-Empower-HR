package service_test

import (
	attendance "be-empower-hr/features/Attendance"
	service "be-empower-hr/features/Attendance/service"
	"be-empower-hr/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddAtt(t *testing.T) {
	qry := mocks.NewAQuery(t)
	hash := mocks.NewHashInterface(t)
	mi := mocks.NewMiddlewaresInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	pu := mocks.NewPdfUtilityInterface(t)
	mu := mocks.NewMapsUtilityInterface(t)
	srv := service.New(qry, hash, mi, au, pu, mu)

	t.Run("Success Add Attendance", func(t *testing.T) {
		newAtt := attendance.Attandance{
			PersonalDataID: 1,
			Clock_in:       "08:00",
			Clock_out:      "17:00",
			Status:         "present",
			Date:           time.Now(),
			Long:           "106.816666",
			Lat:            "-6.200000",
			Notes:          "On time",
		}

		qry.On("IsDateExists", newAtt.PersonalDataID, newAtt.Date).Return(false, nil).Once()
		qry.On("GetCompany", newAtt.PersonalDataID).Return([]attendance.CompanyDataEntity{{ID: 1, CompanyAddress: "Jakarta"}}, nil).Once()
		mu.On("GeoCode", "Jakarta").Return(-6.200000, 106.816666, nil).Once()
		mu.On("Haversine", -6.200000, 106.816666, -6.200000, 106.816666).Return(50.0).Once()
		qry.On("Create", newAtt).Return(nil).Once()

		err := srv.AddAtt(newAtt)

		qry.AssertExpectations(t)
		mu.AssertExpectations(t)

		assert.NoError(t, err)
	})

	t.Run("Failed Add Attendance - Date Exists", func(t *testing.T) {
		newAtt := attendance.Attandance{
			PersonalDataID: 1,
			Clock_in:       "08:00",
			Clock_out:      "17:00",
			Status:         "present",
			Date:           time.Now(),
			Long:           "106.816666",
			Lat:            "-6.200000",
			Notes:          "On time",
		}

		qry.On("IsDateExists", newAtt.PersonalDataID, newAtt.Date).Return(true, nil).Once()

		err := srv.AddAtt(newAtt)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, "attendance record already exists for this date", err.Error())
	})

	// Add more tests for other scenarios like invalid latitude/longitude, server errors, etc.
}

func TestUpdateAtt(t *testing.T) {
	qry := mocks.NewAQuery(t)
	hash := mocks.NewHashInterface(t)
	mi := mocks.NewMiddlewaresInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	pu := mocks.NewPdfUtilityInterface(t)
	mu := mocks.NewMapsUtilityInterface(t)
	srv := service.New(qry, hash, mi, au, pu, mu)

	t.Run("Success Update Attendance", func(t *testing.T) {
		updateAtt := attendance.Attandance{
			PersonalDataID: 1,
			Clock_in:       "08:00",
			Clock_out:      "17:00",
			Status:         "present",
			Date:           time.Now(),
			Long:           "106.816666",
			Lat:            "-6.200000",
			Notes:          "Updated note",
		}

		qry.On("GetCompany", updateAtt.PersonalDataID).Return([]attendance.CompanyDataEntity{{ID: 1, CompanyAddress: "Jakarta"}}, nil).Once()
		mu.On("GeoCode", "Jakarta").Return(-6.200000, 106.816666, nil).Once()
		mu.On("Haversine", -6.200000, 106.816666, -6.200000, 106.816666).Return(50.0).Once()
		qry.On("Update", uint(1), updateAtt).Return(nil).Once()

		err := srv.UpdateAtt(1, updateAtt)

		qry.AssertExpectations(t)
		mu.AssertExpectations(t)

		assert.NoError(t, err)
	})

	// Add more tests for other scenarios like invalid latitude/longitude, server errors, etc.
}

func TestDeleteAttByID(t *testing.T) {
	qry := mocks.NewAQuery(t)
	hash := mocks.NewHashInterface(t)
	mi := mocks.NewMiddlewaresInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	pu := mocks.NewPdfUtilityInterface(t)
	mu := mocks.NewMapsUtilityInterface(t)
	srv := service.New(qry, hash, mi, au, pu, mu)

	t.Run("Success Delete Attendance", func(t *testing.T) {
		attID := uint(1)

		qry.On("DeleteAttbyId", attID).Return(nil).Once()

		err := srv.DeleteAttByID(attID)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
	})

	// Add more tests for error scenarios
}

func TestGetAttByPersonalID(t *testing.T) {
	qry := mocks.NewAQuery(t)
	hash := mocks.NewHashInterface(t)
	mi := mocks.NewMiddlewaresInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	pu := mocks.NewPdfUtilityInterface(t)
	mu := mocks.NewMapsUtilityInterface(t)
	srv := service.New(qry, hash, mi, au, pu, mu)

	t.Run("Success Get Attendance by Personal ID", func(t *testing.T) {
		personalID := uint(1)
		attDetail := []attendance.AttendanceDetail{
			{Name: "Test User", PersonalDataID: 1, ScheduleIn: "08:00", ScheduleOut: "17:00", ClockIn: "08:05", ClockOut: "17:10", Date: "23-02-2023"},
		}

		qry.On("GetAttByPersonalID", personalID, "", 10, 0).Return(attDetail, nil).Once()
		qry.On("GetTotalAttendancesCountbyPerson", personalID).Return(int64(1), nil).Once()

		result, count, err := srv.GetAttByPersonalID(personalID, "", 10, 0)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, attDetail, result)
		assert.Equal(t, int64(1), count)
	})

	// Add more tests for error scenarios and searchBox usage
}

func TestDownloadAllAtt(t *testing.T) {
	qry := mocks.NewAQuery(t)
	hash := mocks.NewHashInterface(t)
	mi := mocks.NewMiddlewaresInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	pu := mocks.NewPdfUtilityInterface(t)
	mu := mocks.NewMapsUtilityInterface(t)
	srv := service.New(qry, hash, mi, au, pu, mu)

	t.Run("Success Download All Attendance", func(t *testing.T) {
		attendanceData := []attendance.Attandance{
			{PersonalDataID: 1, Clock_in: "08:05", Clock_out: "17:10", Date: time.Now()},
		}

		qry.On("GetAllAttDownload").Return(attendanceData, nil).Once()
		pu.On("DownloadPdf", attendanceData, "Attendance.pdf").Return(nil).Once()

		err := srv.DownloadAllAtt()

		qry.AssertExpectations(t)
		pu.AssertExpectations(t)

		assert.NoError(t, err)
	})

	// Add more tests for error scenarios
}
