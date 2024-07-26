package handler

type UserRequest struct {
	PersonalData PersonalData  `json:"personal_data"`
	PayrollData  []PayrollData `json:"payroll_data"`
	CompanyData  CompanyData   `json:"company_data"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PersonalData struct {
	ProfilePicture string           `json:"profile_picture"`
	Name           string           `json:"name"`
	Email          string           `json:"email"`
	Password       string           `json:"password"`
	PhoneNumber    string           `json:"phone_number"`
	PlaceBirth     string           `json:"place_birth"`
	BirthDate      string           `json:"birth_date"`
	Gender         string           `json:"gender"`
	Religion       string           `json:"religion"`
	NIK            string           `json:"nik"`
	Address        string           `json:"address"`
	EmploymentData []EmploymentData `json:"employment_data"`
}

type EmploymentData struct {
	EmploymentStatus string        `json:"employment_status"`
	JoinDate         string        `json:"join_date"`
	Department       string        `json:"department,"`
	JobPosition      string        `json:"job_position"`
	JobLevel         string        `json:"job_level"`
	Schedule         string        `json:"schedule"`
	ApprovalLine     string        `json:"approval_line"`
	Manager          string        `json:"manager"`
	Payrolls         []PayrollData `json:"payrolls"`
}

type PayrollData struct {
	Salary        float64 `json:"salary"`
	BankName      string  `json:"bank_name"`
	AccountNumber int     `json:"account_number"`
}

type CompanyData struct {
	CompanyName string `json:"company_name"`
}
