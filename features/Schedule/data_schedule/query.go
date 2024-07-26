package dataschedule

import (
	schedule "be-empower-hr/features/Schedule"

	"gorm.io/gorm"
)

type scheduleQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) schedule.DataScheduleInterface {
	return &scheduleQuery{
		db: db,
	}
}

// CreateSchedule implements schedule.DataScheduleInterface.
func (sc *scheduleQuery) CreateSchedule(schedule schedule.ScheduleDataEntity) (uint, error) {
	newSchedule := ScheduleData{
		CompanyID:     schedule.CompanyID,
		Name:          schedule.Name,
		EffectiveDate: schedule.EffectiveDate,
		ScheduleIn:    schedule.ScheduleIn,
		ScheduleOut:   schedule.ScheduleOut,
		BreakStart:    schedule.BreakStart,
		BreakEnd:      schedule.BreakEnd,
		Days:          schedule.Days,
		Description:   schedule.Description,
	}
	if err := sc.db.Create(&newSchedule).Error; err != nil {
		return 0, err
	}
	return newSchedule.ID, nil
}

// DeleteSchedule implements schedule.DataScheduleInterface.
func (sc *scheduleQuery) DeleteSchedule(scheduleid uint) error {
	if err := sc.db.Delete(&ScheduleData{}, scheduleid).Error; err != nil {
		return err
	}
	return nil
}

// GetAllSchedule implements schedule.DataScheduleInterface.
func (sc *scheduleQuery) GetAllSchedule() ([]schedule.ScheduleDataEntity, error) {
	var schedules []ScheduleData
	if err := sc.db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	var result []schedule.ScheduleDataEntity
	for _, s := range schedules {
		result = append(result, schedule.ScheduleDataEntity{
			ScheduleId:    s.ID,
			CompanyID:     s.CompanyID,
			Name:          s.Name,
			EffectiveDate: s.EffectiveDate,
			ScheduleIn:    s.ScheduleIn,
			ScheduleOut:   s.ScheduleOut,
			BreakStart:    s.BreakStart,
			BreakEnd:      s.BreakEnd,
			Days:          s.Days,
			Description:   s.Description,
		})
	}
	return result, nil
}

// GetScheduleById implements schedule.DataScheduleInterface.
func (sc *scheduleQuery) GetScheduleById(scheduleid uint) (*schedule.ScheduleDataEntity, error) {
	var scheduleData ScheduleData
	if err := sc.db.First(&scheduleData, scheduleid).Error; err != nil {
		return nil, err
	}
	return &schedule.ScheduleDataEntity{
		ScheduleId:    scheduleData.ID,
		CompanyID:     scheduleData.CompanyID,
		Name:          scheduleData.Name,
		EffectiveDate: scheduleData.EffectiveDate,
		ScheduleIn:    scheduleData.ScheduleIn,
		ScheduleOut:   scheduleData.ScheduleOut,
		BreakStart:    scheduleData.BreakStart,
		BreakEnd:      scheduleData.BreakEnd,
		Days:          scheduleData.Days,
		Description:   scheduleData.Description,
	}, nil
}

// UpdateSchedule implements schedule.DataScheduleInterface.
func (sc *scheduleQuery) UpdateSchedule(scheduleid uint, account schedule.ScheduleDataEntity) error {
	var schedule ScheduleData
	if err := sc.db.First(&schedule, scheduleid).Error; err != nil {
		return err
	}
	schedule.CompanyID = account.CompanyID
	schedule.Name = account.Name
	schedule.EffectiveDate = account.EffectiveDate
	schedule.ScheduleIn = account.ScheduleIn
	schedule.ScheduleOut = account.ScheduleOut
	schedule.BreakStart = account.BreakStart
	schedule.BreakEnd = account.BreakEnd
	schedule.Days = account.Days
	schedule.Description = account.Description
	if err := sc.db.Save(&schedule).Error; err != nil {
		return err
	}
	return nil
}
