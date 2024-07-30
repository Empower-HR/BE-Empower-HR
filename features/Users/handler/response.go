package handler

type UserResponse struct {
	PersonalDataID uint   `json:"personal_data_id"`
	CompanyID      uint   `json:"company_id"`
	Name           string `json:"name"`
	WorkEmail      string `json:"work_email"`
	PhoneNumber    string `json:"phone_number"`
	Department     string `json:"department"`
	JobPosition    string `json:"job_position"`
	CompanyName    string `json:"company_name"`
}

type ProfileResponse struct {
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	PlaceBirthDate string `json:"place_birth"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Religion       string `json:"religion"`
	NIK            string `json:"nik"`
	Address        string `json:"address"`
	EmploymentData []EmploymentDataResponse
}

type EmploymentDataResponse struct {
	EmploymentStatus string `json:"employment_status"`
	JoinDate         string `json:"join_date"`
	Department       string `json:"department"`
	JobPosition      string `json:"job_position"`
	JobLevel         string `json:"job_level"`
	Schedule         string `json:"schedule"`
	ApprovalLine     string `json:"approval_line"`
	Manager          string `json:"manager"`
}

type AllUsersResponse struct {
	PersonalDataID   uint   `json:"id"`
	Name             string `json:"name"`
	JobPosition      string `json:"job_position"`
	JobLevel         string `json:"job_level"`
	EmploymentStatus string `json:"employment_status"`
	JoinDate         string `json:"join_date"`
}

type PersonalDataResponse struct {
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone"`
	PlaceBirth     string `json:"place_birth"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Status         string `json:"status"`
	Religion       string `json:"religion"`
	NIK            string `json:"nik"`
	Address        string `json:"address"`
}

type PayrollResponse struct {
	Salary        float64 `json:"salary"`
	BankName      string  `json:"bank_name"`
	AccountNumber int     `json:"account_number"`
}

type DashboardStatsResponses struct {
	CurrentDate              string  `json:"current_date"`
	PersonalDataName         string  `json:"name"`
	TotalUsers               int64   `json:"total_users,omitempty"`
	MalePercentage           float64 `json:"male"`
	FemalePercentage         float64 `json:"female"`
	ContractUsersPercentage  float64 `json:"contract"`
	PermanentUsersPercentage float64 `json:"permanent"`
	PayrollRecords           int64   `json:"payroll,omitempty"`
	LeavesRecords            int64   `json:"leaves,omitempty"`
	AttendanceRecords        int64   `json:"attendance,omitempty"`
}
