package datausers

import (
	dataleaves "be-empower-hr/features/Leaves/data_leaves"
	datapayroll "be-empower-hr/features/Payroll/data_payroll"
	users "be-empower-hr/features/Users"
	"time"

	"gorm.io/gorm"
)

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
	EmploymentData []EmploymentData        `gorm:"foreignKey:PersonalDataID"`
	Leaves         []dataleaves.LeavesData `gorm:"foreignKey:PersonalDataID"`
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
	Payrolls         []datapayroll.PayrollData `gorm:"foreignKey:EmploymentDataID"`
}

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
}

type PayrollData struct {
	gorm.Model
	EmploymentDataID uint
	Salary           float64
	BankName         string
	AccountNumber    int
}

type LeavesData struct {
	gorm.Model
	StartDate      time.Time
	EndDate        time.Time
	Reason         string
	Status         string
	TotalLeave     int
	PersonalDataID uint
}


func ToQueryEmploymentEmployee(input users.EmploymentDataEntity) EmploymentData {
	return EmploymentData{
		PersonalDataID 	 : input.PersonalDataID,
		EmploymentStatus : input.EmploymentStatus,
		JoinDate         : input.JoinDate,
		Department       : input.Department,
		JobPosition      : input.JobPosition,
		JobLevel         : input.JobLevel,
		Schedule         : input.Schedule,
		ApprovalLine     : input.ApprovalLine,
		Manager          : input.Manager,
	}
}