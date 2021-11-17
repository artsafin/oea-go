package hellenic

import (
	"fmt"
	"io"
	"oea-go/internal/employee/dto"
	"oea-go/internal/hellenic/txt"
	"time"
)

func CreatePayrollFile(wr io.Writer, invoices dto.Invoices, submissionDate time.Time, valueDate time.Time) error {
	f := txt.NewFile(submissionDate)

	for _, invoice := range invoices {

		if invoice.BankDetails == nil {
			continue
		}

		senderBankDetails := txt.SenderBankDetails{
			Account:   invoice.Employee.LegalEntity.AccountNumber,
			ValueDate: valueDate,
			Notes: []string{
				fmt.Sprintf("Invoice Number:%s", invoice.InvoiceNo),
				fmt.Sprintf("Paid from %s", invoice.Employee.LegalEntity.OfficialName),
			},
		}

		recipientBankDetails := txt.RecipientBankDetails{
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
		}

		f.AddRecord(senderBankDetails, recipientBankDetails)
	}

	_, err := wr.Write([]byte(f.String()))

	return err
}
