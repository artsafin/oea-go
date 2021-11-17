package txt

import (
	"fmt"
	"strings"
)

type worldwideBankRecord struct {
	debitAccount              string
	submissionDate            string
	valueDate                 string
	transactionCosts          transactionCosts
	receiveTransactionDetails bool
	recipient                 RecipientBankDetails
	note1                     string
	note2                     string
}

func (r *worldwideBankRecord) String() string {
	return fmt.Sprintf("I|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		underscore(strings.ReplaceAll(r.debitAccount, "-", "")),
		underscore(strings.ReplaceAll(r.recipient.Account, "-", "")),
		underscore(r.recipient.Amount.Humanize("#.###,##")),
		underscore(r.recipient.Amount.CurrencyISO4217()),
		underscore(r.submissionDate),
		underscore(r.valueDate),
		underscore(r.recipient.BeneficiaryName),
		underscore(r.recipient.BeneficiaryAddress1),
		underscore(r.recipient.BeneficiaryAddress2),
		underscore(r.recipient.BeneficiaryBankName),
		underscore(r.recipient.BeneficiaryBankAddress1),
		underscore(r.recipient.BeneficiaryBankAddress2),
		underscore(r.recipient.BeneficiaryBankAddress3),
		underscore(r.recipient.BeneficiarySWIFT),
		r.transactionCosts,
		boolToYesNoFlag(r.receiveTransactionDetails),
		underscore(r.note1),
		underscore(r.note2),
		"_", "_", // Notes 3-4
		"_",           // Sorting code
		"_",           // Receiver’s Correspondent Swift Code
		"_",           // Receiver’s Correspondent Bank Name
		"_", "_", "_", // Receiver’s Correspondent Address 1-3
		"_",           // Third Reimbursement Institution Swift Code
		"_",           // Third Reimbursement Institution Bank Name
		"_", "_", "_", // Third Reimbursement Institution Address 1-3
		underscore(r.recipient.IntermediarySWIFT),
	)
}
