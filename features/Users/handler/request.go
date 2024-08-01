package handler

import (
	users "be-empower-hr/features/Users"
)

type UserRequest struct {
	Name        string `json:"name"`
	WorkEmail   string `json:"work_email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Department  string `json:"department"`
	JobPosition string `json:"job_position"`
	Company     string `json:"company_name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAdminRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
	PlaceBirth     string `json:"place_birth" form:"place_birth"`
	BirthDate      string `json:"birth_date" form:"birth_date"`
	Gender         string `json:"gender" form:"gender"`
	Religion       string `json:"religion" form:"religion"`
	NIK            string `json:"nik" form:"nik"`
	Address        string `json:"address" form:"address"`
}

type EmploymentData struct {
	EmploymentStatus string `json:"employment_status"`
	JoinDate         string `json:"join_date"`
	Department       string `json:"department"`
	JobPosition      string `json:"job_position"`
	JobLevel         string `json:"job_level"`
	Schedule         string `json:"schedule"`
	ApprovalLine     string `json:"approval_line"`
	Manager          string `json:"manager"`
}

type PersonalData struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	Password    string `json:"password"`
	PlaceBirth  string `json:"place_birth"`
	BirthDate   string `json:"birth_date"`
	Gender      string `json:"gender"`
	Status      string `json:"status"`
	Religion    string `json:"religion"`
	NIK         string `json:"nik"`
	Address     string `json:"address"`
}

type Payroll struct {
	Salary        float64 `json:"salary"`
	BankName      string  `json:"bank_name"`
	AccountNumber int     `json:"account_number"`
}

type Leaves struct {
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
	Status     string `json:"status"`
	TotalLeave int    `json:"total_leave"`
}

func ToModelEmploymentData(ed EmploymentData) users.EmploymentDataEntity {
	return users.EmploymentDataEntity{
		EmploymentStatus: ed.EmploymentStatus,
		JoinDate:         ed.JoinDate,
		Department:       ed.Department,
		JobPosition:      ed.JobPosition,
		JobLevel:         ed.JobLevel,
		Schedule:         ed.Schedule,
		ApprovalLine:     ed.ApprovalLine,
		Manager:          ed.Manager,
	}
}

func ToModelPersonalData(pd PersonalData) users.PersonalDataEntity {
	return users.PersonalDataEntity{
		ProfilePicture: "",
		Name:           pd.Name,
		Email:          pd.Email,
		PhoneNumber:    pd.PhoneNumber,
		Password:       pd.Password,
		PlaceBirth:     pd.PlaceBirth,
		BirthDate:      pd.BirthDate,
		Gender:         pd.Gender,
		Status:         pd.Status,
		Religion:       pd.Religion,
		NIK:            pd.NIK,
		Address:        pd.Address,
	}
}

func ToModelPayroll(py Payroll) users.PayrollDataEntity {
	return users.PayrollDataEntity{
		Salary:        py.Salary,
		BankName:      py.BankName,
		AccountNumber: py.AccountNumber,
	}
}

func ToModelLeaves(pd Leaves) users.LeavesDataEntity {
	return users.LeavesDataEntity{
		StartDate:  pd.StartDate,
		EndDate:    pd.EndDate,
		Reason:     pd.Reason,
		Status:     pd.Status,
		TotalLeave: pd.TotalLeave,
	}
}

type NewEmployeeRequest struct {
	PersonalData   PersonalData   `json:"personal"`
	EmploymentData EmploymentData `json:"employment"`
	Payroll        Payroll        `json:"payroll"`
	Leaves         Leaves         `json:"leaves"`
}
