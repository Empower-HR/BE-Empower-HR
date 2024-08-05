package service_test

import (
	payroll "be-empower-hr/features/Payroll"
	service "be-empower-hr/features/Payroll/service"
	"be-empower-hr/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePayroll(t *testing.T) {
	qry := mocks.NewDataPayrollInterface(t)
	pui := mocks.NewPdfUtilityInterface(t)
	srv := service.New(qry, pui)
	input := payroll.PayrollDataEntity{
		// PayrollID:        uint(1),
		EmploymentDataID: uint(1),
		Salary:           5000000,
		BankName:         "BCA",
		AccountNumber:    63762999,
	}

	t.Run("Error Bank Name", func(t *testing.T) {
		data := payroll.PayrollDataEntity{
			PayrollID:        uint(1),
			EmploymentDataID: uint(1),
			Salary:           5000000,
			BankName:         "",
			AccountNumber:    63762999,
		}
		_, err := srv.CreatePayroll(data)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "bank name cannot be empty")
	})

	t.Run("Success Create Payroll", func(t *testing.T) {
		qry.On("CreatePayroll", input).Return(input, nil).Once()

		data, err := srv.CreatePayroll(input)

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, input, data)
	})
}

func TestGetAllPayroll(t *testing.T) {
	qry := mocks.NewDataPayrollInterface(t)
	pui := mocks.NewPdfUtilityInterface(t)
	srv := service.New(qry, pui)
	result := []payroll.PayrollResponse{
		{
			ID:             uint(1),
			EmploymentName: "Joko",
			Date:           "02-01-2024",
			Position:       "Software Engineer",
		},
	}

	t.Run("Success Get All Payroll", func(t *testing.T) {
		qry.On("GetAllPayroll").Return(result, nil).Once()
		data, err := srv.GetAllPayroll()

		qry.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, result, data)
	})
}

func TestGetPayrollDownload(t *testing.T) {
	qry := new(mocks.DataPayrollInterface)
	pui := new(mocks.PdfUtilityInterface)
	srv := service.New(qry, pui)

	t.Run("Success Get Payroll Download", func(t *testing.T) {
		id := uint(1)
		payrollResult := payroll.PayrollResponsePDF{
			ID:             uint(1),
			EmploymentName: "JohnDoe",
			Date:           "02-01-2024",
			Position:       "Software Engineer",
			Salary:         6000000,
			BankName:       "BCA",
			AccountNumber:  63762999,
		}

		qry.On("GetPayrollDownload", id).Return(payrollResult, nil).Once()
		pui.On("DownloadPdfPayroll", payrollResult, "PayslipJohnDoe.pdf").Return(nil).Once()

		err := srv.GetPayrollDownload(id)
		assert.NoError(t, err)

		qry.AssertExpectations(t)
		pui.AssertExpectations(t)
	})

	t.Run("Error retrieving payroll records", func(t *testing.T) {
		id := uint(1)
		qry.On("GetPayrollDownload", id).Return(payroll.PayrollResponsePDF{}, errors.New("error retrieving payroll records")).Once()

		err := srv.GetPayrollDownload(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "error retrieving payroll records")

		qry.AssertExpectations(t)
	})

	t.Run("Error downloading PDF", func(t *testing.T) {
		id := uint(1)
		payrollResult := payroll.PayrollResponsePDF{
			ID:             uint(1),
			EmploymentName: "JohnDoe",
			Date:           "02-01-2024",
			Position:       "Software Engineer",
			Salary:         6000000,
			BankName:       "BCA",
			AccountNumber:  63762999,
		}

		qry.On("GetPayrollDownload", id).Return(payrollResult, nil).Once()
		pui.On("DownloadPdfPayroll", payrollResult, "PayslipJohnDoe.pdf").Return(errors.New("error downloading pdf")).Once()

		err := srv.GetPayrollDownload(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "error download pdf")

		qry.AssertExpectations(t)
		pui.AssertExpectations(t)
	})
}
