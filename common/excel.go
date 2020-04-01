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

	interpolateCell(f, "A8", "%SSSS%", data.BeneficiaryRequisites())
	interpolateCell(f, "A21", "%SSSS%", data.PayerRequisites())
	interpolateCell(f, "AM5", "%D%", fmt.Sprint(data.Number()))
	interpolateCell(f, "AM6", "%MMMM DD YYYY%", data.DateFull())
	interpolateCell(f, "D28", "%MMMM YYYY%", data.DateYm())
	interpolateCell(f, "Z28", "%D%", data.HourRate().String())
	interpolateCell(f, "AE28", "%D%", fmt.Sprint(data.Hours()))
	interpolateCell(f, "AH28", "%D%", data.TotalEur().String())
	interpolateCell(f, "AH30", "%D%", data.TotalEur().String())
	interpolateCell(f, "AM32", "%D%", data.TotalEur().String())
	interpolateCell(f, "U38", "%SSSS%", data.BeneficiaryName())
	interpolateCell(f, "U44", "%SSSS%", data.PayerName())

	euros := int(data.TotalEur() / 100)
	cents := int(data.TotalEur()) - euros*100
	words := fmt.Sprintf("%s EURO and %02d cents", num2words.Convert(euros), cents)
	interpolateCell(f, "AM33", "%SSSS%", words)

	if err := f.Write(wr); err != nil {
		panic(err)
	}
}
