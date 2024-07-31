package datapayroll

import (
	payroll "be-empower-hr/features/Payroll"
	"be-empower-hr/utils"

	"gorm.io/gorm"
)

type payrollQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) payroll.DataPayrollInterface {
	payrollQuery := &payrollQuery{
		db: db,
	}
	return payrollQuery
}

func (pq *payrollQuery) CreatePayroll(p payroll.PayrollDataEntity) (payroll.PayrollDataEntity, error) {
	newPayroll := PayrollData{
		EmploymentDataID: p.EmploymentDataID,
		Salary:           p.Salary,
		BankName:         p.BankName,
		AccountNumber:    p.AccountNumber,
	}

	result := pq.db.Create(&newPayroll)
	if result.Error != nil {
		return payroll.PayrollDataEntity{}, result.Error
	}

	createdPayroll := payroll.PayrollDataEntity{
		PayrollID:        newPayroll.ID,
		EmploymentDataID: newPayroll.EmploymentDataID,
		Salary:           newPayroll.Salary,
		BankName:         newPayroll.BankName,
		AccountNumber:    newPayroll.AccountNumber,
	}
	return createdPayroll, nil
}

func (pq *payrollQuery) GetAllPayroll() ([]payroll.PayrollResponse, error) {
	var payrolls []PayrollData
	if err := pq.db.Find(&payrolls).Error; err != nil {
		return nil, err
	}     
	var result []payroll.PayrollResponse
	for _, p := range payrolls {
		emp, err := pq.GetEmpById(p.EmploymentDataID)
		if err != nil {
			return nil, err
		}
		personal, err := pq.GetUserById(emp.PersonalDataID)
		if err != nil {
			return nil, err
		}
		date, err := utils.DateToString(p.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, payroll.PayrollResponse{
			ID:             p.ID,
			EmploymentName: personal.Name,
			Date:           date,
			Position:       emp.JobPosition,
		})
	}
	return result, nil
}


func (pq *payrollQuery) GetPayrollDownload(ID uint) (payroll.PayrollResponsePDF, error){
	var payrolls PayrollData

	 err := pq.db.Where("id = ?", ID).First(&payrolls).Error

	 if err != nil {
		return payroll.PayrollResponsePDF{}, err
	}
	var result payroll.PayrollResponsePDF

	emp, err := pq.GetEmpById(payrolls.EmploymentDataID)
	if err != nil {
		return payroll.PayrollResponsePDF{}, err
	}
	personal, err := pq.GetUserById(emp.PersonalDataID)
	if err != nil {
		return payroll.PayrollResponsePDF{}, err
	}
	date, err := utils.DateToString(payrolls.CreatedAt)
	if err != nil {
		return payroll.PayrollResponsePDF{}, err
	}

	result = payroll.PayrollResponsePDF{
		ID: payrolls.ID ,
		EmploymentName: personal.Name,
		Date: date,
		Position: emp.JobPosition,
		Salary: payrolls.Salary,
		BankName: payrolls.BankName,
		AccountNumber: payrolls.AccountNumber,
	};

	return result, nil;
}

func (pq *payrollQuery) GetEmpById(id uint) (payroll.EmploymentDataEntity, error) {
	var emp EmploymentData
	err := pq.db.Where("id = ?", id).First(&emp).Error
	if err != nil {
		return payroll.EmploymentDataEntity{}, err
	}
	empEntity := payroll.EmploymentDataEntity{
		EmploymentDataID: emp.ID,
		PersonalDataID:   emp.PersonalDataID,
		JobPosition:      emp.JobPosition,
	}
	return empEntity, nil
}

func (pq *payrollQuery) GetUserById(id uint) (payroll.PersonalDataEntity, error) {
	var personal PersonalData
	err := pq.db.Where("id = ?", id).First(&personal).Error
	if err != nil {
		return payroll.PersonalDataEntity{}, err
	}
	personalData := payroll.PersonalDataEntity{
		PersonalDataID: personal.ID,
		Name:           personal.Name,
	}
	return personalData, nil
}



