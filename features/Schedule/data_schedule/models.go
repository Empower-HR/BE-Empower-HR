package dataschedule

import (
	"gorm.io/gorm"
)

type ScheduleData struct {
	gorm.Model
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

type PersonalData struct {
	gorm.Model
	CompanyID      uint
	ProfilePicture string
	Name           string
	Email          string
	Password       string
	PhoneNumber    string
	PlaceBirth     string
	BirthDate      string
	Gender         string
	Religion       string
	NIK            string
	Address        string
	Role           string
}
