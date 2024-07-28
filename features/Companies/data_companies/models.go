package datacompanies

import (
	companies "be-empower-hr/features/Companies"
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
	Npwp           int
	CompanyAddress string
	Signature      string
	Schedules      []dataschedule.ScheduleData `gorm:"foreignKey:CompanyID"`
	PersonalData   []datausers.PersonalData    `gorm:"foreignKey:CompanyID"`
};


func (cd *CompanyData) ToCompanyEntity() companies.CompanyDataEntity{
	return companies.CompanyDataEntity{
		ID			   : cd.ID,
		CompanyPicture : cd.CompanyPicture,
		CompanyName    : cd.CompanyName,
		Email          : cd.Email,
		PhoneNumber    : cd.PhoneNumber,
		Npwp           : cd.Npwp,
		CompanyAddress : cd.CompanyAddress,
		Signature      : cd.Signature,
	}
};

func ToCompanyQuery(cmp companies.CompanyDataEntity) CompanyData {
	return CompanyData{
		CompanyPicture : cmp.CompanyPicture,
		CompanyName    : cmp.CompanyName,
		Email          : cmp.Email,
		PhoneNumber    : cmp.PhoneNumber,
		Npwp           : cmp.Npwp,
		CompanyAddress : cmp.CompanyAddress,
		Signature      : cmp.Signature,
	}
}
