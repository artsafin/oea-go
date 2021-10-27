package txt

import (
	"oea-go/internal/codatypes"
	"strings"
	"time"
)

type RecipientBankDetails struct {
	Account                 string
	Amount                  codatypes.MoneyEur
	BeneficiaryName         string
	BeneficiaryAddress1     string
	BeneficiaryAddress2     string
	BeneficiaryBankName     string
	BeneficiaryBankAddress1 string
	BeneficiaryBankAddress2 string
	BeneficiaryBankAddress3 string
	BeneficiarySWIFT        string
	IntermediarySWIFT       string
}

type file struct {
	header  *header
	records []worldwideBankRecord
}

func NewFile(submissionDate time.Time, debitAccount string) file {
	return file{
		header:  &header{uploadType: "S", submissionDate: submissionDate.Format("20060102"), debitAccount: debitAccount},
		records: make([]worldwideBankRecord, 0),
	}
}

func (f *file) AddRecord(details RecipientBankDetails) {
	f.header.increment(details.Amount)

	rec := worldwideBankRecord{
		debitAccount:              f.header.debitAccount,
		submissionDate:            f.header.submissionDate,
		valueDate:                 f.header.submissionDate,
		transactionCosts:          TransactionCostsOurs,
		receiveTransactionDetails: true,
		recipient:                 details,
	}
	f.records = append(f.records, rec)
}

func (f *file) String() string {
	sb := strings.Builder{}
	sb.WriteString(f.header.String())
	sb.WriteString("\n")

	for _, r := range f.records {
		sb.WriteString(r.String())
		sb.WriteString("\n")
	}

	return sb.String()
}
