package companies

import (
	dataschedule "be-empower-hr/features/Schedule/data_schedule"
	datausers "be-empower-hr/features/Users/data_users"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type CompanyDataEntity struct {
	ID             uint
	CompanyPicture string
	CompanyName    string
	Email          string
	PhoneNumber    string
	Npwp           int
	CompanyAddress string
	Signature      string
	Schedules      []dataschedule.ScheduleData
	PersonalData   []datausers.PersonalData
}

type Handler interface {
	GetCompany() echo.HandlerFunc
	UpdateCompany() echo.HandlerFunc
}

type Query interface {
	GetCompany() (CompanyDataEntity, error)
	UpdateCompany(ID uint, updateCompany CompanyDataEntity) error
	GetCompanyID(ID uint) (CompanyDataEntity, error)
}

type Service interface {
	GetCompany(ID uint) (CompanyDataEntity, error)
	UpdateCompany(ID uint,companyPicture, companySignature *multipart.FileHeader, updateCompany CompanyDataEntity) error
}
