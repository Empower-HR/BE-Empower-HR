package service_test

import (
	leaves "be-empower-hr/features/Leaves"
	service "be-empower-hr/features/Leaves/service"
	"be-empower-hr/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateLeaveStatus(t *testing.T) {
	qry := mocks.NewDataLeavesInterface(t)
	srv := service.New(qry)

	t.Run("Success Update Leave Status", func(t *testing.T) {
		leaveId := uint(1)
		updateLeaves := leaves.LeavesDataEntity{
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
		qry.On("GetPersonalDataByID", uint(1), mock.AnythingOfType("*leaves.PersonalDataEntity")).Return(nil).Once()
		qry.On("GetLeavesDetail", leaveId).Return(updateLeaves, nil).Once()
		leavesData := leaves.LeavesDataEntity{
			LeavesID:   leaveId,
			TotalLeave: 20,
		}
		qry.On("GetLeavesDataByID", leaveId, mock.AnythingOfType("*leaves.LeavesDataEntity")).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*leaves.LeavesDataEntity)
			*arg = leavesData
		}).Return(nil).Once()

		qry.On("UpdateLeaveData", mock.AnythingOfType("leaves.LeavesDataEntity")).Return(nil).Once()
		qry.On("UpdateLeaveStatus", leaveId, updateLeaves).Return(nil).Once()
		err := srv.UpdateLeaveStatus(leaveId, leaveId, updateLeaves)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
}
