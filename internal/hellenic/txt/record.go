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
		strings.ReplaceAll(r.debitAccount, "-", ""),
		strings.ReplaceAll(r.recipient.Account, "-", ""),
		r.recipient.Amount.Humanize("#.###,##"),
		r.recipient.Amount.CurrencyISO4217(),
		r.submissionDate,
		r.valueDate,
		limit(underscore(r.recipient.BeneficiaryName), 35),
		limit(underscore(r.recipient.BeneficiaryAddress1), 35),
		limit(underscore(r.recipient.BeneficiaryAddress2), 35),
		limit(underscore(r.recipient.BeneficiaryBankName), 35),
		limit(underscore(r.recipient.BeneficiaryBankAddress1), 35),
		limit(underscore(r.recipient.BeneficiaryBankAddress2), 35),
		limit(underscore(r.recipient.BeneficiaryBankAddress3), 35),
		underscore(r.recipient.BeneficiarySWIFT),
		r.transactionCosts,
		boolToYesNoFlag(r.receiveTransactionDetails),
		limit(underscore(r.note1), 35),
		limit(underscore(r.note2), 35),
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
