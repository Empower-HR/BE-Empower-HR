package service

import (
	payroll "be-empower-hr/features/Payroll"
	"errors"
)

type payrollService struct {
	payrollData payroll.DataPayrollInterface
}

func New(pd payroll.DataPayrollInterface) payroll.ServicePayrollInterface {
	return &payrollService{
		payrollData: pd,
	}
}

func (ps *payrollService) CreatePayroll(p payroll.PayrollDataEntity) (payroll.PayrollDataEntity, error) {
	if p.BankName == "" {
		return payroll.PayrollDataEntity{}, errors.New("bank name cannot be empty")
	}
	return ps.payrollData.CreatePayroll(p)
}

func (ps *payrollService) GetAllPayroll() ([]payroll.PayrollResponse, error) {
	return ps.payrollData.GetAllPayroll()
}
