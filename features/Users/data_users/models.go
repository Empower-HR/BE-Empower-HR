package datausers

import (
	dataatt "be-empower-hr/features/Attendance/data_attendance"
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
	Status         string
	Religion       string
	NIK            string
	Address        string
	Role           string
	EmploymentData []EmploymentData        `gorm:"foreignKey:PersonalDataID"`
	Leaves         []dataleaves.LeavesData `gorm:"foreignKey:PersonalDataID"`
	Attandance     []dataatt.Attandance    `gorm:"foreignKey:PersonalDataID"`
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

type DashboardStats struct {
	TotalUsers     int64
	MaleUsers      int64
	FemaleUsers    int64
	ContractUsers  int64
	PermanentUsers int64
	PayrollRecords int64
}

func ToQueryEmploymentEmployee(input users.EmploymentDataEntity) EmploymentData {
	return EmploymentData{
		PersonalDataID:   input.PersonalDataID,
		EmploymentStatus: input.EmploymentStatus,
		JoinDate:         input.JoinDate,
		Department:       input.Department,
		JobPosition:      input.JobPosition,
		JobLevel:         input.JobLevel,
		Schedule:         input.Schedule,
		ApprovalLine:     input.ApprovalLine,
		Manager:          input.Manager,
	}
}

func ToPersonalDataQuery(input users.PersonalDataEntity) PersonalData {
	return PersonalData{
		CompanyID:      input.CompanyID,
		ProfilePicture: input.ProfilePicture,
		Name:           input.Name,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		PlaceBirth:     input.PlaceBirth,
		BirthDate:      input.BirthDate,
		Gender:         input.Gender,
		Status:         input.Status,
		Religion:       input.Religion,
		NIK:            input.NIK,
		Address:        input.Address,
		Role:           input.Role,
	}
}

func ToEmploymentQuery(input users.EmploymentDataEntity) EmploymentData {
	return EmploymentData{
		PersonalDataID:   input.PersonalDataID,
		EmploymentStatus: input.EmploymentStatus,
		JoinDate:         input.JoinDate,
		Department:       input.Department,
		JobPosition:      input.JobPosition,
		JobLevel:         input.JobLevel,
		Schedule:         input.Schedule,
		ApprovalLine:     input.ApprovalLine,
		Manager:          input.Manager,
	}
}

func ToPayrollQuery(input users.PayrollDataEntity) PayrollData {
	return PayrollData{
		EmploymentDataID: input.EmploymentDataID,
		Salary:           input.Salary,
		BankName:         input.BankName,
		AccountNumber:    input.AccountNumber,
	}
}

func ToLeavesQuery(input users.LeavesDataEntity) LeavesData {
	return LeavesData{
		StartDate:      input.StartDate,
		EndDate:        input.EndDate,
		Reason:         input.Reason,
		Status:         input.Status,
		TotalLeave:     input.TotalLeave,
		PersonalDataID: input.PersonalDataID,
	}
}
