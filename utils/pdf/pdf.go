package pdf

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type PdfUtilityInterface interface {
	DownloadPdf(url string) (string, error)
	UploadPdf(url, filePath string) error
}

type pdfUtility struct{}

func NewPdfUtility() PdfUtilityInterface {
	return &pdfUtility{}
}

func (pu *pdfUtility) DownloadPdf(url string) (string, error) {
	// url := "https://example.com/path/to/your.pdf"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return "", err
	}
	defer response.Body.Close()

	file, err := os.Create("your.pdf")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return "", err
	}

	return "PDF downloaded successfully", nil
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
