package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/divan/num2words"
	"oea-go/office/dto"
)

const sheetName = "Sheet1"

func RenderExcelTemplate(outdir string, templateSource []byte, invoice *dto.Invoice) {
	f, err := excelize.OpenReader(bytes.NewReader(templateSource))
	if err != nil {
		panic(err)
	}

	f.SetCellValue(sheetName, "AM5", strings.Replace(f.GetCellValue(sheetName, "AM5"), "%D%", fmt.Sprint(invoice.Number), 1))
	f.SetCellValue(sheetName, "AM6", strings.Replace(f.GetCellValue(sheetName, "AM6"), "%MMMM DD YYYY%", invoice.InvoiceDateFull(), 1))
	f.SetCellValue(sheetName, "D28", strings.Replace(f.GetCellValue(sheetName, "D28"), "%MMMM YYYY%", invoice.InvoiceDateYm(), 1))
	f.SetCellValue(sheetName, "Z28", strings.Replace(f.GetCellValue(sheetName, "Z28"), "%D%", invoice.HourRate.String(), 1))
	f.SetCellValue(sheetName, "AE28", strings.Replace(f.GetCellValue(sheetName, "AE28"), "%D%", fmt.Sprint(invoice.Hours), 1))
	f.SetCellValue(sheetName, "AH28", strings.Replace(f.GetCellValue(sheetName, "AH28"), "%D%", invoice.TotalEur.String(), 1))
	f.SetCellValue(sheetName, "AH30", strings.Replace(f.GetCellValue(sheetName, "AH30"), "%D%", invoice.TotalEur.String(), 1))
	f.SetCellValue(sheetName, "AM32", strings.Replace(f.GetCellValue(sheetName, "AM32"), "%D%", invoice.TotalEur.String(), 1))

	euros := int(invoice.TotalEur / 100)
	cents := int(invoice.TotalEur) - euros*100
	words := fmt.Sprintf("%s EURO and %02d cents", num2words.Convert(euros), cents)
	f.SetCellValue(sheetName, "AM33", strings.Replace(f.GetCellValue(sheetName, "AM33"), "%SSSS%", words, 1))

	f.SaveAs(fmt.Sprintf("%s/%s.xlsx", outdir, invoice.Filename))
}
