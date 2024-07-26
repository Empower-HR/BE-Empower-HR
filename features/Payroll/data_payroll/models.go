package datapayroll

import "gorm.io/gorm"

type PayrollData struct {
	gorm.Model
	EmploymentDataID uint
	Salary           float64
	BankName         string
	AccountNumber    int
}
