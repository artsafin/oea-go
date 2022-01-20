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

		if invoice.RecipientDetails == nil {
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
			Account:                 invoice.RecipientDetails.Account,
			Amount:                  invoice.EURTotal,
			BeneficiaryName:         invoice.EmployeeName,
			BeneficiaryAddress1:     invoice.RecipientDetails.Address1,
			BeneficiaryAddress2:     invoice.RecipientDetails.Address2,
			BeneficiaryBankName:     invoice.RecipientDetails.Bank.Name,
			BeneficiaryBankAddress1: invoice.RecipientDetails.Bank.Address1,
			BeneficiaryBankAddress2: invoice.RecipientDetails.Bank.Address2,
			BeneficiaryBankAddress3: invoice.RecipientDetails.Bank.Address3,
			BeneficiarySWIFT:        invoice.RecipientDetails.Bank.BeneficiarySWIFT,
			IntermediarySWIFT:       invoice.RecipientDetails.Bank.IntermediarySWIFT,
		}

		f.AddRecord(senderBankDetails, recipientBankDetails)
	}

	_, err := wr.Write([]byte(f.String()))

	return err
}
