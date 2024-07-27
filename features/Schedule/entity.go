package schedule

import (
	"time"
)

type ScheduleDataEntity struct {
	ScheduleId    uint
	CompanyID     uint
	Name          string
	EffectiveDate time.Time
	ScheduleIn    string
	ScheduleOut   string
	BreakStart    string
	BreakEnd      string
	Days          int
	Description   string
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
