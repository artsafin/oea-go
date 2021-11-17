package txt

import (
	"oea-go/internal/codatypes"
	"strings"
	"time"
)

type SenderBankDetails struct {
	Account   string
	ValueDate time.Time
	Notes     []string
}

func (d *SenderBankDetails) Note1() string {
	if len(d.Notes) >= 1 {
		return d.Notes[0]
	}
	return ""
}

func (d *SenderBankDetails) Note2() string {
	if len(d.Notes) >= 2 {
		return d.Notes[1]
	}
	return ""
}

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
	submissionDate time.Time
	header         *header
	records        []worldwideBankRecord
}

func NewFile(submissionDate time.Time) file {
	return file{
		submissionDate: submissionDate,
		header:         &header{uploadType: "M", submissionDate: "_", debitAccount: "_"},
		records:        make([]worldwideBankRecord, 0),
	}
}

func (f *file) AddRecord(sendDetails SenderBankDetails, rcpDetails RecipientBankDetails) {
	f.header.increment(rcpDetails.Amount)

	rec := worldwideBankRecord{
		debitAccount:              sendDetails.Account,
		submissionDate:            f.submissionDate.Format("20060102"),
		valueDate:                 sendDetails.ValueDate.Format("20060102"),
		transactionCosts:          TransactionCostsOurs,
		receiveTransactionDetails: true,
		recipient:                 rcpDetails,
		note1:                     sendDetails.Note1(),
		note2:                     sendDetails.Note2(),
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
