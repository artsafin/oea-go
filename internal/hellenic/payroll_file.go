package hellenic

import (
	"fmt"
	"io"
	"oea-go/internal/employee/codaschema"
	"oea-go/internal/hellenic/txt"
	"time"
)

func CreatePayrollFile(wr io.Writer, invoices []codaschema.Invoice, submissionDate time.Time, valueDate time.Time) error {
	f := txt.NewFile(submissionDate)

	for _, invoice := range invoices {

		if invoice.RecipientDetails.First() == nil ||
			invoice.RecipientDetails.First().BeneficiaryBank.First() == nil ||
			invoice.Employee.First() == nil ||
			invoice.Employee.First().LegalEntity.First() == nil {
			continue
		}

		legalEntity := invoice.Employee.First().LegalEntity.First()
		rcpt := invoice.RecipientDetails.First()
		rcptBank := rcpt.BeneficiaryBank.First()

		senderBankDetails := txt.SenderBankDetails{
			Account:   legalEntity.AccountNumber,
			ValueDate: valueDate,
			Notes: []string{
				fmt.Sprintf("Invoice Number:%s", invoice.InvoiceHash),
				fmt.Sprintf("Paid from %s", legalEntity.OfficialName),
			},
		}

		recipientBankDetails := txt.RecipientBankDetails{
			Account:                 rcpt.Account,
			Amount:                  invoice.EURTotalMoney(),
			BeneficiaryName:         invoice.Employee.String(),
			BeneficiaryAddress1:     rcpt.Address1,
			BeneficiaryAddress2:     rcpt.Address2,
			BeneficiaryBankName:     rcptBank.Name,
			BeneficiaryBankAddress1: rcptBank.Address1,
			BeneficiaryBankAddress2: rcptBank.Address2,
			BeneficiaryBankAddress3: rcptBank.Address3,
			BeneficiarySWIFT:        rcptBank.BeneficiarySWIFT,
			IntermediarySWIFT:       rcptBank.IntermediarySWIFT,
		}

		f.AddRecord(senderBankDetails, recipientBankDetails)
	}

	_, err := wr.Write([]byte(f.String()))

	return err
}
