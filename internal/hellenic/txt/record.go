package txt

import "fmt"

type worldwideBankRecord struct {
	debitAccount              string
	submissionDate            string
	valueDate                 string
	transactionCosts          transactionCosts
	receiveTransactionDetails bool
	recipient                 RecipientBankDetails
}

func (r *worldwideBankRecord) String() string {
	return fmt.Sprintf("I|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		underscore(r.debitAccount),
		underscore(r.recipient.Account),
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
		"_", "_", "_", "_", // Notes 1-4
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
