package handler

type UserResponse struct {
	PersonalData   PersonalData     `json:"personal_data"`
	EmploymentData []EmploymentData `json:"employment_data"`
	PayrollData    []PayrollData    `json:"payroll_data"`
	CompanyName    string           `json:"company_name"`
}
