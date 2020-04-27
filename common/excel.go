package common

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/divan/num2words"
)

const sheetName = "Sheet1"

type InvoiceDataProvider interface {
	BeneficiaryRequisites() string
	PayerRequisites() string
	Number() string
	DateFull() string
	DateYm() string
	HourRate() MoneyEur
	Hours() uint16
	TotalEur() MoneyEur
	BeneficiaryName() string
	PayerName() string
	Filename() string
	ContractNumber() string
	ContractDate() string
}

func interpolateCell(file *excelize.File, cell, placeholder, value string) {
	rawValue := file.GetCellValue(sheetName, cell)
	newValue := strings.Replace(rawValue, placeholder, value, 1)
	file.SetCellValue(sheetName, cell, newValue)
}

func RenderExcelTemplate(wr io.Writer, templateSource []byte, data InvoiceDataProvider) {
	f, err := excelize.OpenReader(bytes.NewReader(templateSource))
	if err != nil {
		panic(err)
	}

	// Header
	interpolateCell(f, "AM4", "%D%", fmt.Sprint(data.Number())) // Invoice Number
	interpolateCell(f, "AM5", "%MMMM DD YYYY%", data.DateFull()) // Invoice Date
	interpolateCell(f, "AM6", "%S%", data.ContractNumber()) // Contract Number
	interpolateCell(f, "AM7", "%S%", data.ContractDate()) // Contract Date

	// Requisites
	interpolateCell(f, "A9", "%SSSS%", data.BeneficiaryRequisites())
	interpolateCell(f, "A22", "%SSSS%", data.PayerRequisites())

	// Invoice line
	interpolateCell(f, "D29", "%MMMM YYYY%", data.DateYm()) // Service type
	interpolateCell(f, "Z29", "%D%", data.HourRate().String()) // Price, EUR
	interpolateCell(f, "AE29", "%D%", fmt.Sprint(data.Hours())) // Qty
	interpolateCell(f, "AH29", "%D%", data.TotalEur().String()) // Cost

	// Subtotal
	interpolateCell(f, "AH31", "%D%", data.TotalEur().String())

	// Total
	interpolateCell(f, "AM33", "%D%", data.TotalEur().String())

	euros := int(data.TotalEur() / 100)
	cents := int(data.TotalEur()) - euros*100
	words := fmt.Sprintf("%s EURO and %02d cents", num2words.Convert(euros), cents)
	interpolateCell(f, "AM34", "%SSSS%", words)

	// Signs
	interpolateCell(f, "U39", "%SSSS%", data.BeneficiaryName())
	interpolateCell(f, "U45", "%SSSS%", data.PayerName())

	if err := f.Write(wr); err != nil {
		panic(err)
	}
}
