package datapayroll

import "gorm.io/gorm"

type PayrollData struct {
	gorm.Model
	EmploymentDataID uint
	Salary           float64
	BankName         string
	AccountNumber    int
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
	Status         string
	Religion       string
	NIK            string
	Address        string
	Role           string
}

type EmploymentData struct {
	gorm.Model
	PersonalDataID   uint
	EmploymentStatus string
	JoinDate         string
	Department       string
	JobPosition      string
	JobLevel         string
	Schedule         string
	ApprovalLine     string
	Manager          string
	Payrolls         []PayrollData `gorm:"foreignKey:EmploymentDataID"`
}
