package schedule

type ScheduleDataEntity struct {
	ScheduleId    uint
	CompanyID     uint
	Name          string
	EffectiveDate string
	ScheduleIn    string
	ScheduleOut   string
	BreakStart    string
	BreakEnd      string
	Days          int
	Description   string
}

type DataScheduleInterface interface {
	CreateSchedule(schedule ScheduleDataEntity) (uint, error)
	UpdateSchedule(scheduleid uint, account ScheduleDataEntity) error
	DeleteSchedule(scheduleid uint) error
	GetAllSchedule() ([]ScheduleDataEntity, error)
	GetScheduleById(scheduleid uint) (*ScheduleDataEntity, error)
}

type ServiceScheduleInterface interface {
	CreateSchedule(schedule ScheduleDataEntity) (uint, error)
	UpdateSchedule(scheduleid uint, account ScheduleDataEntity) error
	DeleteSchedule(scheduleid uint) error
	GetAllSchedule() ([]ScheduleDataEntity, error)
	GetScheduleById(scheduleid uint) (*ScheduleDataEntity, error)
}
