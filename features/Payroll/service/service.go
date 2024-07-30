package service

import (
	payroll "be-empower-hr/features/Payroll"
	"be-empower-hr/utils/pdf"
	"errors"
)

type payrollService struct {
	payrollData payroll.DataPayrollInterface
	pdfUtils pdf.PdfUtilityInterface
}

func New(pd payroll.DataPayrollInterface, p pdf.PdfUtilityInterface) payroll.ServicePayrollInterface {
	return &payrollService{
		payrollData: pd,
		pdfUtils: p,
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

func (ps *payrollService) GetPayrollDownload(ID uint) (error){

	result, err := ps.payrollData.GetPayrollDownload(ID);
	if err != nil {
		return errors.New("error retrieving payroll records")
	};

	err = ps.pdfUtils.DownloadPdfPayroll(result, "Payslip" + result.EmploymentName + ".pdf");
	if err != nil {
		return errors.New("error download pdf")
	}
	
	return nil;
}
