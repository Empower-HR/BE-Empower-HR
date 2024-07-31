package pdf

import (
	attendance "be-empower-hr/features/Attendance"
	payroll "be-empower-hr/features/Payroll"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type PdfUtilityInterface interface {
	DownloadPdf(items []attendance.Attandance, filename string) error
	DownloadPdfPayroll(item payroll.PayrollResponsePDF, filename string) error 
	UploadPdf(url, filePath string) error
}

type pdfUtility struct{}

func NewPdfUtility() PdfUtilityInterface {
	return &pdfUtility{}
}

func (pu *pdfUtility) DownloadPdf(items []attendance.Attandance, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Items List")

	// Set header font
	pdf.SetFont("Arial", "", 12)

	// Add table headers
	pdf.Ln(10)
	pdf.Cell(60, 10, "Date")
	pdf.Cell(60, 10, "Clock In")
	pdf.Cell(60, 10, "Clock Out")
	pdf.Ln(10)

	// Add table rows
	for _, item := range items {
		pdf.Cell(60, 10, item.Date)
		pdf.Cell(60, 10, item.Clock_in)
		pdf.Cell(60, 10, item.Clock_out)
		pdf.Ln(10)
	}

	// Save the file
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return err
	}
	return nil
}

func (pu *pdfUtility) DownloadPdfPayroll(item payroll.PayrollResponsePDF, filename string) error {
	pdf := gofpdf.New("L", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Payslip :" + item.EmploymentName)

	// Set header font
	pdf.SetFont("Arial", "", 12)

	// Add table headers with adjusted spacing
	pdf.Ln(10)
	pdf.Cell(25, 10, "Payroll_ID")
	pdf.Cell(40, 10, "Employment Name")
	pdf.Cell(25, 10, "Date")
	pdf.Cell(40, 10, "Position")
	pdf.Cell(25, 10, "Salary")
	pdf.Cell(40, 10, "Bank Name")
	pdf.Cell(40, 10, "Account Number")
	pdf.Ln(10)

	IdStr := strconv.Itoa(int(item.ID))
	salaryStr := strconv.FormatFloat(item.Salary, 'f', 2, 64)
	AccStr := strconv.Itoa(int(item.AccountNumber))

	// Add table rows with adjusted spacing
	pdf.Cell(25, 10, IdStr)
	pdf.Cell(40, 10, item.EmploymentName)
	pdf.Cell(25, 10, item.Date)
	pdf.Cell(40, 10, item.Position)
	pdf.Cell(25, 10, salaryStr)
	pdf.Cell(40, 10, item.BankName)
	pdf.Cell(40, 10, AccStr)
	pdf.Ln(10)

	// Save the file
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return err
	}
	return nil
}


func (pu *pdfUtility) UploadPdf(url, filePath string) error {
	// 	url := "http://localhost:8080/upload"
	// filePath := "path/to/your.pdf"

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return fmt.Errorf("could not create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("could not copy file data: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("could not close writer: %w", err)
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not do request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}

	fmt.Println(string(respBody))
	return nil
}
