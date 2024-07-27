package handler

import users "be-empower-hr/features/Users"

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

func ToModelEmploymentData(ed EmploymentData) users.EmploymentDataEntity {
	return users.EmploymentDataEntity{
	EmploymentStatus : ed.EmploymentStatus,
	JoinDate         : ed.JoinDate,
	Department       : ed.Department,
	JobPosition      : ed.JobPosition,
	JobLevel         : ed.JobLevel,
	Schedule         : ed.Schedule,
	ApprovalLine     : ed.ApprovalLine,
	Manager          : ed.Manager,
	}
}
