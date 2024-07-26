package companies

import (
	dataschedule "be-empower-hr/features/Schedule/data_schedule"
	datausers "be-empower-hr/features/Users/data_users"
)

type CompanyDataEntity struct {
	CompanyPicture string
	CompanyName    string
	Email          string
	PhoneNumber    string
	Address        string
	Npwp           int
	CompanyAddress string
	Signature      string
	Schedules      []dataschedule.ScheduleData
	PersonalData   []datausers.PersonalData
}
