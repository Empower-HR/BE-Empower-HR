package service

import (
	schedule "be-empower-hr/features/Schedule"
	"be-empower-hr/utils"
	"errors"
)

type scheduleService struct {
	scheduleData   schedule.DataScheduleInterface
	accountUtility utils.AccountUtilityInterface
}

func New(sc schedule.DataScheduleInterface, au utils.AccountUtilityInterface) schedule.ServiceScheduleInterface {
	return &scheduleService{
		scheduleData:   sc,
		accountUtility: au,
	}
}

// CreateSchedule implements schedule.ServiceScheduleInterface.
func (ss *scheduleService) CreateSchedule(schedule schedule.ScheduleDataEntity) (uint, error) {
	if schedule.Days <= 0 {
		return 0, errors.New("days must be greater than 0")
	}

	_, err := ss.accountUtility.NumberLoop(schedule.Days)
	if err != nil {
		return 0, err
	}

	if schedule.Name == "" || schedule.ScheduleIn == "" || schedule.ScheduleOut == "" {
		return 0, errors.New("name, effective_date, schedule_in, and schedule_out cannot be empty")
	}
	return ss.scheduleData.CreateSchedule(schedule)
}

// DeleteSchedule implements schedule.ServiceScheduleInterface.
func (ss *scheduleService) DeleteSchedule(scheduleid uint) error {
	return ss.scheduleData.DeleteSchedule(scheduleid)
}

// GetAllSchedule implements schedule.ServiceScheduleInterface.
func (ss *scheduleService) GetAllSchedule() ([]schedule.ScheduleDataEntity, error) {
	return ss.scheduleData.GetAllSchedule()
}

// GetScheduleById implements schedule.ServiceScheduleInterface.
func (ss *scheduleService) GetScheduleById(scheduleid uint) (*schedule.ScheduleDataEntity, error) {
	return ss.scheduleData.GetScheduleById(scheduleid)
}

// UpdateSchedule implements schedule.ServiceScheduleInterface.
func (ss *scheduleService) UpdateSchedule(scheduleid uint, account schedule.ScheduleDataEntity) error {
	// if account.Name == "" || account.EffectiveDate == "" || account.ScheduleIn == "" || account.ScheduleOut == "" {
	// 	return errors.New("nama/effective_date/schedule_in/schedule_out tidak boleh kosong")
	// }

	if account.Days <= 0 {
		return errors.New("days must be greater than 0")
	}

	_, err := ss.accountUtility.NumberLoop(account.Days)
	if err != nil {
		return err
	}

	return ss.scheduleData.UpdateSchedule(scheduleid, account)
}
