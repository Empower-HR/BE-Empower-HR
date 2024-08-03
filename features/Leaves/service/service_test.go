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
}
