package schedule

import (
	"time"
)

type ScheduleDataEntity struct {
	ScheduleId    uint      `json:"id"`
	CompanyID     uint      `json:"company_id"`
	Name          string    `json:"name"`
	EffectiveDate time.Time `json:"effective_date"`
	ScheduleIn    string    `json:"schedule_in"`
	ScheduleOut   string    `json:"schedule_out"`
	BreakStart    string    `json:"break_start"`
	BreakEnd      string    `json:"break_end"`
	Days          string    `json:"repeat_until"`
	Description   string    `json:"description"`
}

type CompanyDataEntity struct {
	ID             uint
	CompanyPicture string
	CompanyName    string
	Email          string
	PhoneNumber    string
	Address        string
	Npwp           int
	CompanyAddress string
	Signature      string
}

type DataScheduleInterface interface {
	CreateSchedule(schedule ScheduleDataEntity) (uint, error)
	UpdateSchedule(scheduleid uint, account ScheduleDataEntity) error
	DeleteSchedule(scheduleid uint) error
	GetAllSchedule() ([]ScheduleDataEntity, error)
	GetScheduleById(scheduleid uint) (*ScheduleDataEntity, error)
}

type Query interface {
	GetCompany(ID uint) (CompanyDataEntity, error)
	UpdateCompany(ID uint, updateCompany CompanyDataEntity) error
}

type ServiceScheduleInterface interface {
	CreateSchedule(schedule ScheduleDataEntity) (uint, error)
	UpdateSchedule(scheduleid uint, account ScheduleDataEntity) error
	DeleteSchedule(scheduleid uint) error
	GetAllSchedule() ([]ScheduleDataEntity, error)
	GetScheduleById(scheduleid uint) (*ScheduleDataEntity, error)
}
