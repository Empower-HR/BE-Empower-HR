package handler

import (
	"be-empower-hr/app/middlewares"
	payroll "be-empower-hr/features/Payroll"
	"be-empower-hr/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PayrollHandler struct {
	payrollService payroll.ServicePayrollInterface
}

func New(ps payroll.ServicePayrollInterface) *PayrollHandler {
	return &PayrollHandler{
		payrollService: ps,
	}
}

func (ph *PayrollHandler) CreatePayroll(c echo.Context) error {
	var req PayrollRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	payrollData := payroll.PayrollDataEntity{
		EmploymentDataID: req.EmploymentDataID,
		Salary:           req.Salary,
		BankName:         req.BankName,
		AccountNumber:    req.AccountNumber,
	}

	payroll, err := ph.payrollService.CreatePayroll(payrollData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Payroll entry created successfully", payroll))
}

func (ph *PayrollHandler) GetAllPayroll(c echo.Context) error {
	payrolls, err := ph.payrollService.GetAllPayroll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Team members retrieved successfully", payrolls))
}

func (ph *PayrollHandler ) DownloadPayrollPdf(c echo.Context) error {
	personalID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if personalID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	ID, err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Invalid request parameters", nil))
	}

	err = ph.payrollService.GetPayrollDownload(uint(ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Download failed", nil))
	}
	return c.File("./Payroll.pdf")
}
