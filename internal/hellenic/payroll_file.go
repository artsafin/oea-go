package hellenic

import (
	"io"
	"oea-go/internal/employee/dto"
	"oea-go/internal/hellenic/txt"
	"time"
)

func CreatePayrollFile(wr io.Writer, invoices dto.Invoices, account string, date time.Time) error {
	f := txt.NewFile(date, account)

	for _, invoice := range invoices {

		if invoice.BankDetails == nil {
			continue
		}

		f.AddRecord(txt.RecipientBankDetails{
			Account:                 invoice.BankDetails.Account,
			Amount:                  invoice.AmountRequestedEur,
			BeneficiaryName:         invoice.EmployeeName,
			BeneficiaryAddress1:     invoice.BankDetails.Address1,
			BeneficiaryAddress2:     invoice.BankDetails.Address2,
			BeneficiaryBankName:     invoice.BankDetails.Bank.Name,
			BeneficiaryBankAddress1: invoice.BankDetails.Bank.Address1,
			BeneficiaryBankAddress2: invoice.BankDetails.Bank.Address2,
			BeneficiaryBankAddress3: invoice.BankDetails.Bank.Address3,
			BeneficiarySWIFT:        invoice.BankDetails.Bank.BeneficiarySWIFT,
			IntermediarySWIFT:       invoice.BankDetails.Bank.IntermediarySWIFT,
		})
	}

	_, err := wr.Write([]byte(f.String()))

	return err
}
