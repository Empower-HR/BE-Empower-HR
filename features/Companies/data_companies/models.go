package datacompanies

import (
	dataschedule "be-empower-hr/features/Schedule/data_schedule"
	datausers "be-empower-hr/features/Users/data_users"

	"gorm.io/gorm"
)

type CompanyData struct {
	gorm.Model
	CompanyPicture string
	CompanyName    string
	Email          string
	PhoneNumber    string
	Address        string
	Npwp           int
	CompanyAddress string
	Signature      string
	Schedules      []dataschedule.ScheduleData `gorm:"foreignKey:CompanyID"`
	PersonalData   []datausers.PersonalData    `gorm:"foreignKey:CompanyID"`
}
