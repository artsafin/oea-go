package excel

import (
	"bytes"
	"fmt"
	"io"
	"oea-go/internal/codatypes"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/divan/num2words"
)

const invoiceSheetName = "Sheet1"

type InvoiceDataProvider interface {
	BeneficiaryRequisites() string
	PayerRequisites() string
	Number() string
	DateFull() string
	DateYm() string
	HourRate() codatypes.MoneyEur
	Hours() uint16
	TotalEur() codatypes.MoneyEur
	BeneficiaryName() string
	PayerName() string
	Filename() string
	ContractNumber() string
	ContractDate() string
}

func invoiceCell(file *excelize.File, cell, placeholder, value string) {
	rawValue, _ := file.GetCellValue(invoiceSheetName, cell)
	newValue := strings.Replace(rawValue, placeholder, value, 1)
	file.SetCellValue(invoiceSheetName, cell, newValue)
}

func RenderInvoice(wr io.Writer, templateSource []byte, data InvoiceDataProvider) error {
	f, err := excelize.OpenReader(bytes.NewReader(templateSource))
	if err != nil {
		return err
	}

	// Header
	invoiceCell(f, "AM4", "%D%", fmt.Sprint(data.Number()))  // Invoice Number
	invoiceCell(f, "AM5", "%MMMM DD YYYY%", data.DateFull()) // Invoice Date
	invoiceCell(f, "AM6", "%S%", data.ContractNumber())      // Contract Number
	invoiceCell(f, "AM7", "%S%", data.ContractDate())        // Contract Date

	// Requisites
	invoiceCell(f, "A9", "%SSSS%", data.BeneficiaryRequisites())
	invoiceCell(f, "A22", "%SSSS%", data.PayerRequisites())

	// Invoice line
	invoiceCell(f, "D29", "%MMMM YYYY%", data.DateYm())     // Service type
	invoiceCell(f, "Z29", "%D%", data.HourRate().String())  // Price, EUR
	invoiceCell(f, "AE29", "%D%", fmt.Sprint(data.Hours())) // Qty
	invoiceCell(f, "AH29", "%D%", data.TotalEur().String()) // Cost

	// Subtotal
	invoiceCell(f, "AH31", "%D%", data.TotalEur().String())

	// Total
	invoiceCell(f, "AM33", "%D%", data.TotalEur().String())

	euros := int(data.TotalEur() / 100)
	cents := int(data.TotalEur()) - euros*100
	words := fmt.Sprintf("%s EURO and %02d cents", num2words.Convert(euros), cents)
	invoiceCell(f, "AM34", "%SSSS%", words)

	// Signs
	invoiceCell(f, "U39", "%SSSS%", data.BeneficiaryName())

	if err := f.Write(wr); err != nil {
		return err
	}

	return nil
}
