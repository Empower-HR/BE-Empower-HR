package handler

type PayrollRequest struct {
	EmploymentDataID uint    `json:"employee_id"`
	Salary           float64 `json:"salary"`
	BankName         string  `json:"bank_name"`
	AccountNumber    int     `json:"account_num"`
}
