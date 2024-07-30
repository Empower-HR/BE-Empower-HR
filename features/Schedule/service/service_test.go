package service_test

import (
	schedule "be-empower-hr/features/Schedule"
	service "be-empower-hr/features/Schedule/service"
	"be-empower-hr/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSchedule(t *testing.T) {
	qry := mocks.NewDataScheduleInterface(t)
	aui := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, aui)
	input := schedule.ScheduleDataEntity{
		CompanyID:     uint(1),
		Name:          "Ali",
		EffectiveDate: time.Now(),
		ScheduleIn:    "08:00",
		ScheduleOut:   "16:00",
		BreakStart:    "12:00",
		BreakEnd:      "13:00",
		Days:          29,
		Description:   "test description",
	}

	t.Run("Error From Validate", func(t *testing.T) {
		data := schedule.ScheduleDataEntity{
			CompanyID:     uint(2),
			Name:          "Ammar",
			EffectiveDate: time.Now(),
			ScheduleIn:    "08:00",
			ScheduleOut:   "16:00",
			BreakStart:    "12:00",
			BreakEnd:      "13:00",
			Days:          0,
			Description:   "test description",
		}
		_, err := srv.CreateSchedule(data)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "days must be greater than 0")
	})

	t.Run("Error Name", func(t *testing.T) {
		data := schedule.ScheduleDataEntity{
			CompanyID:     uint(2),
			Name:          "",
			EffectiveDate: time.Now(),
			ScheduleIn:    "08:00",
			ScheduleOut:   "16:00",
			BreakStart:    "12:00",
			BreakEnd:      "13:00",
			Days:          29,
			Description:   "test description",
		}
		_, err := srv.CreateSchedule(data)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "name, effective_date, schedule_in, and schedule_out cannot be empty")
	})

	t.Run("Success Create Schedule", func(t *testing.T) {
		qry.On("CreateSchedule", input).Return(uint(1), nil).Once()

		data, err := srv.CreateSchedule(input)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), data)
	})
}

func TestDeleteSchedule(t *testing.T) {
	qry := mocks.NewDataScheduleInterface(t)
	aui := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, aui)

	t.Run("Success Delete Schedule", func(t *testing.T) {
		id := uint(1)
		qry.On("DeleteSchedule", id).Return(nil).Once()
		err := srv.DeleteSchedule(id)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestGetAllSchedule(t *testing.T) {
	qry := mocks.NewDataScheduleInterface(t)
	aui := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, aui)
	result := []schedule.ScheduleDataEntity{
		{
			CompanyID:     uint(1),
			Name:          "Ali",
			EffectiveDate: time.Now(),
			ScheduleIn:    "08:00",
			ScheduleOut:   "16:00",
			BreakStart:    "12:00",
			BreakEnd:      "13:00",
			Days:          29,
			Description:   "test description",
		},
	}

	t.Run("Success Get All Schedule", func(t *testing.T) {
		qry.On("GetAllSchedule").Return(result, nil).Once()
		data, err := srv.GetAllSchedule()

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, result, data)
	})
}

func TestGetScheduleById(t *testing.T) {
	qry := mocks.NewDataScheduleInterface(t)
	aui := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, aui)
	scheduleId := uint(1)
	result := &schedule.ScheduleDataEntity{
		ScheduleId:    scheduleId,
		CompanyID:     uint(1),
		Name:          "Ali",
		EffectiveDate: time.Now(),
		ScheduleIn:    "08:00",
		ScheduleOut:   "16:00",
		BreakStart:    "12:00",
		BreakEnd:      "13:00",
		Days:          29,
		Description:   "test description",
	}

	t.Run("Success Get Schedule By Id", func(t *testing.T) {
		qry.On("GetScheduleById", scheduleId).Return(result, nil).Once()
		data, err := srv.GetScheduleById(scheduleId)
		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, result, data)
	})
}

func TestUpdateSchedule(t *testing.T) {
	qry := mocks.NewDataScheduleInterface(t)
	aui := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, aui)
	scheduleId := uint(1)

	input := schedule.ScheduleDataEntity{
		CompanyID:     uint(1),
		Name:          "Ali",
		EffectiveDate: time.Now(),
		ScheduleIn:    "08:00",
		ScheduleOut:   "16:00",
		BreakStart:    "12:00",
		BreakEnd:      "13:00",
		Days:          29,
		Description:   "test description",
	}

	t.Run("Error From Validate", func(t *testing.T) {
		data := schedule.ScheduleDataEntity{
			CompanyID:     uint(2),
			Name:          "Ammar",
			EffectiveDate: time.Now(),
			ScheduleIn:    "08:00",
			ScheduleOut:   "16:00",
			BreakStart:    "12:00",
			BreakEnd:      "13:00",
			Days:          0,
			Description:   "test description",
		}
		err := srv.UpdateSchedule(scheduleId, data)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "days must be greater than 0")
	})

	t.Run("Success Update Schedule", func(t *testing.T) {
		qry.On("UpdateSchedule", scheduleId, input).Return(nil).Once()

		err := srv.UpdateSchedule(scheduleId, input)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
	})
}
