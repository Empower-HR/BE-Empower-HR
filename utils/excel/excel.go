package excel

import (
	"bytes"

	"github.com/xuri/excelize/v2"
)


type ExcelUtilityInterface interface{
	DownloadExcel(data []struct{}) ([]byte, error)
}

type excelUntility struct {}

func NewExcelUtility() ExcelUtilityInterface{
	return &excelUntility{}
}

func (xu *excelUntility) DownloadExcel(data []struct{}) ([]byte, error) {
	f := excelize.NewFile();

	sheet := "Products";
	index, _ := f.NewSheet(sheet);

	// hapus sheet pertama yang di buat secara otomatis 
	f.DeleteSheet("Sheet1")

	// set sheet nya menjadi yang pertama
	f.SetActiveSheet(index)

	// header 
	headers := []string{}; // string untuk header data di excel

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1);
		f.SetCellValue(sheet, cell, header)
	};


	var buf bytes.Buffer;

    err := f.Write(&buf)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}