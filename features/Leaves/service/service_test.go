package service_test

import (
	leaves "be-empower-hr/features/Leaves"
	service "be-empower-hr/features/Leaves/service"
	"be-empower-hr/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestLeave(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	t.Run("Success Request Leave", func(t *testing.T) {
		leave := leaves.LeavesDataEntity{
			LeavesID:       1,
			Name:           "Test",
			JobPosition:    "Software Engineer",
			StartDate:      "20 Agustus 2021",
			EndDate:        "30 Agustus 2021",
			Reason:         "Flu",
			Status:         "approved",
			TotalLeave:     12,
			PersonalDataID: uint(1),
		}
		userId := uint(1)

		qry.On("RequestLeave", leave).Return(nil).Once()
		err := srv.RequestLeave(userId, leave)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("Failed Request Leave", func(t *testing.T) {
		leave := leaves.LeavesDataEntity{
			LeavesID:       1,
			Name:           "Test",
			JobPosition:    "Software Engineer",
			StartDate:      "20 Agustus 2021",
			EndDate:        "30 Agustus 2021",
			Reason:         "Flu",
			Status:         "approved",
			TotalLeave:     12,
			PersonalDataID: 1,
		}
		userId := uint(1)
		expectedError := errors.New("failed to request leave")

		qry.On("RequestLeave", leave).Return(expectedError).Once()
		err := srv.RequestLeave(userId, leave)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetLeavesById(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	t.Run("Success Get Leaves By Id", func(t *testing.T) {
		leaveId := uint(1)
		leave := &leaves.LeavesDataEntity{
			LeavesID:       1,
			Name:           "Test",
			JobPosition:    "Software Engineer",
			StartDate:      "20 Agustus 2021",
			EndDate:        "30 Agustus 2021",
			Reason:         "Flu",
			Status:         "approved",
			TotalLeave:     12,
			PersonalDataID: uint(1),
		}

		qry.On("GetLeavesDetail", leaveId).Return(leave, nil).Once()
		data, err := srv.GetLeavesByID(leaveId)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, leave, data)
	})

	t.Run("Error From Validate", func(t *testing.T) {
		leaveId := uint(1)
		expectedError := errors.New("Error getting leave detail:")

		qry.On("GetLeavesDetail", leaveId).Return(nil, expectedError).Once()
		_, err := srv.GetLeavesByID(leaveId)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Error getting leave detail:")
	})
}

func TestDashboardEmployees(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	t.Run("Success DashboardEmployees", func(t *testing.T) {
		companyId := uint(1)
		page := 1
		pageSize := 5
		leave := &leaves.DashboardStats{
			Quota:             2,
			Used:              2,
			PersonalDataNames: "Rian",
		}

		qry.On("DashboardEmployees", companyId, page, pageSize).Return(leave, nil).Once()
		stats, err := srv.DashboardEmployees(companyId, page, pageSize)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, leave, stats)
	})

	t.Run("Error From Validate", func(t *testing.T) {
		companyId := uint(1)
		page := 1
		pageSize := 5
		expectedError := errors.New("Error retrieving dashboard stats")

		qry.On("DashboardEmployees", companyId, page, pageSize).Return(nil, expectedError).Once()
		_, err := srv.DashboardEmployees(companyId, page, pageSize)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Error retrieving dashboard stats")
	})

	t.Run("Error Stats", func(t *testing.T) {
		companyId := uint(1)
		page := 1
		pageSize := 5
		expectedError := errors.New("failed to retrieve dashboard statistics")

		qry.On("DashboardEmployees", companyId, page, pageSize).Return(nil, expectedError).Once()
		_, err := srv.DashboardEmployees(companyId, page, pageSize)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve dashboard statistics")
	})

	t.Run("Invalid Parameters", func(t *testing.T) {
		companyId := uint(0) // Invalid ID
		page := -1           // Invalid page number
		pageSize := 0        // Invalid page size
		expectedError := errors.New("invalid parameters")

		qry.On("DashboardEmployees", companyId, page, pageSize).Return(nil, expectedError).Once()
		_, err := srv.DashboardEmployees(companyId, page, pageSize)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid parameters")
	})

	t.Run("No Data", func(t *testing.T) {
		companyId := uint(1)
		page := 1
		pageSize := 5

		qry.On("DashboardEmployees", companyId, page, pageSize).Return(nil, nil).Once()
		stats, err := srv.DashboardEmployees(companyId, page, pageSize)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Nil(t, stats)
	})
}

func TestViewLeaveHistory(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	t.Run("Success_View_Leave_History_By_Status", func(t *testing.T) {
		companyID := uint(1)
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := "approved"
		startDate := ""
		endDate := ""
		expectedLeaveData := []leaves.LeavesDataEntity{
			{LeavesID: 1, PersonalDataID: personalDataID, Status: status, Reason: "Vacation"},
			{LeavesID: 2, PersonalDataID: personalDataID, Status: status, Reason: "Medical"},
		}

		qry.On("GetLeavesByStatus", personalDataID, status).Return(expectedLeaveData, nil).Once()

		leaveHistory, err := srv.ViewLeaveHistory(companyID, personalDataID, page, pageSize, status, startDate, endDate)

		assert.NoError(t, err)
		assert.NotNil(t, leaveHistory)
		assert.Equal(t, 2, len(leaveHistory))
		qry.AssertExpectations(t)
	})

	t.Run("Success_View_Leave_History_By_Date_Range", func(t *testing.T) {
		companyID := uint(1)
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := "2024-08-01"
		endDate := "2024-08-10"
		expectedLeaveData := []leaves.LeavesDataEntity{
			{LeavesID: 1, PersonalDataID: personalDataID, Status: "approved", Reason: "Vacation"},
			{LeavesID: 2, PersonalDataID: personalDataID, Status: "approved", Reason: "Medical"},
		}

		qry.On("GetLeavesByDateRange", personalDataID, startDate, endDate).Return(expectedLeaveData, nil).Once()

		leaveHistory, err := srv.ViewLeaveHistory(companyID, personalDataID, page, pageSize, status, startDate, endDate)

		assert.NoError(t, err)
		assert.NotNil(t, leaveHistory)
		assert.Equal(t, 2, len(leaveHistory))
		qry.AssertExpectations(t)
	})

	t.Run("Success_View_Leave_History_Default", func(t *testing.T) {
		companyID := uint(1)
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := ""
		endDate := ""
		expectedLeaveData := []leaves.LeavesDataEntity{
			{LeavesID: 1, PersonalDataID: personalDataID, Status: "approved", Reason: "Vacation"},
			{LeavesID: 2, PersonalDataID: personalDataID, Status: "approved", Reason: "Medical"},
		}

		qry.On("GetLeaveHistory", companyID, personalDataID, page, pageSize).Return(expectedLeaveData, nil).Once()

		leaveHistory, err := srv.ViewLeaveHistory(companyID, personalDataID, page, pageSize, status, startDate, endDate)

		assert.NoError(t, err)
		assert.NotNil(t, leaveHistory)
		assert.Equal(t, 2, len(leaveHistory))
		qry.AssertExpectations(t)
	})

	t.Run("Error_GetLeavesByStatus", func(t *testing.T) {
		companyID := uint(1)
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := "approved"

		qry.On("GetLeavesByStatus", personalDataID, status).Return(nil, errors.New("no records found")).Once()

		leaveHistory, err := srv.ViewLeaveHistory(companyID, personalDataID, page, pageSize, status, "", "")

		assert.Error(t, err)
		assert.Nil(t, leaveHistory)
		assert.Equal(t, "no records found", err.Error())
		qry.AssertExpectations(t)
	})

	t.Run("Error_GetLeavesByDateRange", func(t *testing.T) {
		companyID := uint(1)
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := "2024-08-01"
		endDate := "2024-08-10"

		qry.On("GetLeavesByDateRange", personalDataID, startDate, endDate).Return(nil, errors.New("no records found")).Once()

		leaveHistory, err := srv.ViewLeaveHistory(companyID, personalDataID, page, pageSize, status, startDate, endDate)

		assert.Error(t, err)
		assert.Nil(t, leaveHistory)
		assert.Equal(t, "no records found", err.Error())
		qry.AssertExpectations(t)
	})

	t.Run("Error_GetLeaveHistory", func(t *testing.T) {
		companyID := uint(1)
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := ""
		endDate := ""

		qry.On("GetLeaveHistory", companyID, personalDataID, page, pageSize).Return(nil, errors.New("no records found")).Once()

		leaveHistory, err := srv.ViewLeaveHistory(companyID, personalDataID, page, pageSize, status, startDate, endDate)

		assert.Error(t, err)
		assert.Nil(t, leaveHistory)
		assert.Equal(t, "no records found", err.Error())
		qry.AssertExpectations(t)
	})
}

func TestViewLeaveHistoryEmployee(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	t.Run("Success_View_Leave_History_By_Status", func(t *testing.T) {
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := "approved"
		startDate := ""
		endDate := ""
		expectedLeaveData := []leaves.LeavesDataEntity{
			{LeavesID: 1, PersonalDataID: personalDataID, Status: status, Reason: "Vacation"},
			{LeavesID: 2, PersonalDataID: personalDataID, Status: status, Reason: "Medical"},
		}

		qry.On("GetLeavesByStatus", personalDataID, status).Return(expectedLeaveData, nil).Once()

		leaveHistory, err := srv.ViewLeaveHistoryEmployee(personalDataID, page, pageSize, status, startDate, endDate)

		assert.NoError(t, err)
		assert.NotNil(t, leaveHistory)
		assert.Equal(t, 2, len(leaveHistory))
		qry.AssertExpectations(t)
	})

	t.Run("Success_View_Leave_History_By_Date_Range", func(t *testing.T) {
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := "2024-08-01"
		endDate := "2024-08-10"
		expectedLeaveData := []leaves.LeavesDataEntity{
			{LeavesID: 1, PersonalDataID: personalDataID, Status: "approved", Reason: "Vacation"},
			{LeavesID: 2, PersonalDataID: personalDataID, Status: "approved", Reason: "Medical"},
		}

		qry.On("GetLeavesByDateRange", personalDataID, startDate, endDate).Return(expectedLeaveData, nil).Once()

		leaveHistory, err := srv.ViewLeaveHistoryEmployee(personalDataID, page, pageSize, status, startDate, endDate)

		assert.NoError(t, err)
		assert.NotNil(t, leaveHistory)
		assert.Equal(t, 2, len(leaveHistory))
		qry.AssertExpectations(t)
	})

	t.Run("Success_View_Leave_History_Default", func(t *testing.T) {
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := ""
		endDate := ""
		expectedLeaveData := []leaves.LeavesDataEntity{
			{LeavesID: 1, PersonalDataID: personalDataID, Status: "approved", Reason: "Vacation"},
			{LeavesID: 2, PersonalDataID: personalDataID, Status: "approved", Reason: "Medical"},
		}

		qry.On("GetLeaveHistoryEmployee", personalDataID, page, pageSize).Return(expectedLeaveData, nil).Once()

		leaveHistory, err := srv.ViewLeaveHistoryEmployee(personalDataID, page, pageSize, status, startDate, endDate)

		assert.NoError(t, err)
		assert.NotNil(t, leaveHistory)
		assert.Equal(t, 2, len(leaveHistory))
		qry.AssertExpectations(t)
	})

	t.Run("Error_GetLeavesByStatus", func(t *testing.T) {
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := "approved"
		startDate := "2024-08-01"
		endDate := "2024-08-10"

		qry.On("GetLeavesByStatus", personalDataID, status).Return(nil, errors.New("no records found")).Once()

		leaveHistory, err := srv.ViewLeaveHistoryEmployee(personalDataID, page, pageSize, status, startDate, endDate)

		assert.Error(t, err)
		assert.Nil(t, leaveHistory)
		assert.Equal(t, "no records found", err.Error())
		qry.AssertExpectations(t)
	})

	t.Run("Error_GetLeavesByDateRange", func(t *testing.T) {
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := "2024-08-01"
		endDate := "2024-08-10"

		qry.On("GetLeavesByDateRange", personalDataID, startDate, endDate).Return(nil, errors.New("no records found")).Once()

		leaveHistory, err := srv.ViewLeaveHistoryEmployee(personalDataID, page, pageSize, status, startDate, endDate)

		assert.Error(t, err)
		assert.Nil(t, leaveHistory)
		assert.Equal(t, "no records found", err.Error())
		qry.AssertExpectations(t)
	})

	t.Run("Error_GetLeaveHistoryEmployee", func(t *testing.T) {
		personalDataID := uint(1)
		page := 1
		pageSize := 10
		status := ""
		startDate := ""
		endDate := ""

		qry.On("GetLeaveHistoryEmployee", personalDataID, page, pageSize).Return(nil, errors.New("no records found")).Once()

		leaveHistory, err := srv.ViewLeaveHistoryEmployee(personalDataID, page, pageSize, status, startDate, endDate)

		assert.Error(t, err)
		assert.Nil(t, leaveHistory)
		assert.Equal(t, "no records found", err.Error())
		qry.AssertExpectations(t)
	})
}

func TestDashboard(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	t.Run("Success_Dashboard", func(t *testing.T) {
		companyID := uint(1)
		totalUsers := int64(100)
		pendingLeaves := int64(25)

		// Set up expectations
		qry.On("CountTotalUsers", companyID).Return(totalUsers, nil).Once()
		qry.On("CountPendingLeaves", companyID).Return(pendingLeaves, nil).Once()

		// Call the method
		stats, err := srv.Dashboard(companyID)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, stats)
		assert.Equal(t, totalUsers, stats.TotalUsers)
		assert.Equal(t, pendingLeaves, stats.LeavesPending)
		qry.AssertExpectations(t)
	})

	t.Run("Error_CountTotalUsers", func(t *testing.T) {
		companyID := uint(1)

		// Set up expectations
		qry.On("CountTotalUsers", companyID).Return(int64(0), errors.New("error counting users")).Once()

		// Call the method
		stats, err := srv.Dashboard(companyID)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, stats)
		assert.Equal(t, "error counting users", err.Error())
		qry.AssertExpectations(t)
	})

	t.Run("Error_CountPendingLeaves", func(t *testing.T) {
		companyID := uint(1)
		totalUsers := int64(100)

		// Set up expectations
		qry.On("CountTotalUsers", companyID).Return(totalUsers, nil).Once()
		qry.On("CountPendingLeaves", companyID).Return(int64(0), errors.New("error counting pending leaves")).Once()

		// Call the method
		stats, err := srv.Dashboard(companyID)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, stats)
		assert.Equal(t, "error counting pending leaves", err.Error())
		qry.AssertExpectations(t)
	})
}

func TestUpdateLeaveStatus(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	updates := leaves.LeavesDataEntity{
		Reason: "Leave",
		Status: "Approved",
	}

	t.Run("Success_Update_Leaves", func(t *testing.T) {
		qry.On("UpdateLeaveStatus", uint(1), uint(1), updates).Return(nil).Once()

		err := srv.UpdateLeaveStatus(uint(1), uint(1), updates)
		assert.NoError(t, err)
		qry.AssertExpectations(t)
	})

	t.Run("Failure_Update_Leaves", func(t *testing.T) {
		expectedErr := assert.AnError
		qry.On("UpdateLeaveStatus", uint(1), uint(1), updates).Return(expectedErr).Once()
		err := srv.UpdateLeaveStatus(uint(1), uint(1), updates)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		qry.AssertExpectations(t)
	})
}

func TestConvertIndonesiaMonthToEnglish(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		monthID := "02 Januari 2006"
		expected := "02 January 2006"
		result, err := service.ConvertIndonesiaMonthToEnglish(monthID)
		assert.NoError(t, err, "format tanggal tidak valid")
		assert.Equal(t, expected, result, "expected %s but got %s", expected, result)
	})

	t.Run("Error", func(t *testing.T) {
		monthID := "InvalidMonth"
		result, err := service.ConvertIndonesiaMonthToEnglish(monthID)
		assert.Error(t, err, "expected an error but got none")
		assert.Empty(t, result, "expected an empty result but got %s", result)
	})
}

func TestCalculateLeaveDays(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		startDate := "10 Agustus 2020"
		endDate := "10 Agustus 2020"
		expectedDays := 1

		days, err := service.CalculateLeaveDays(startDate, endDate)

		assert.NoError(t, err)
		assert.Equal(t, expectedDays, days)
	})

	t.Run("Error startDate", func(t *testing.T) {
		startDate := ""
		endDate := "2024-08-10"

		days, err := service.CalculateLeaveDays(startDate, endDate)

		assert.Error(t, err)
		assert.Equal(t, 0, days)
	})

	t.Run("Error endDate", func(t *testing.T) {
		startDate := "10 Agustus 2020"
		endDate := ""

		days, err := service.CalculateLeaveDays(startDate, endDate)

		assert.Error(t, err)
		assert.Equal(t, 0, days)
	})

	t.Run("Error invalid date format", func(t *testing.T) {
		startDate := "10 August 2020" // Incorrect month
		endDate := "15 August 2020"

		days, err := service.CalculateLeaveDays(startDate, endDate)

		assert.Error(t, err)
		assert.Equal(t, 0, days)
	})

	t.Run("End date before start date", func(t *testing.T) {
		startDate := "15 Agustus 2020"
		endDate := "10 Agustus 2020"

		days, err := service.CalculateLeaveDays(startDate, endDate)

		assert.Error(t, err)
		assert.Equal(t, 0, days)
	})
}
