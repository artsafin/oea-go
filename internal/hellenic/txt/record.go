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
		lPadSp(strings.ReplaceAll(r.debitAccount, "-", ""), 16),
		lPadSp(strings.ReplaceAll(r.recipient.Account, "-", ""), 35),
		lPadSp(r.recipient.Amount.Humanize("#.###,##"), 18),
		r.recipient.Amount.CurrencyISO4217(),
		r.submissionDate,
		r.valueDate,
		rPadSp(limit(r.recipient.BeneficiaryName, 35), 35),
		rPadSp(limit(r.recipient.BeneficiaryAddress1, 35), 35),
		rPadSp(limit(r.recipient.BeneficiaryAddress2, 35), 35),
		rPadSp(limit(r.recipient.BeneficiaryBankName, 35), 35),
		rPadSp(limit(r.recipient.BeneficiaryBankAddress1, 35), 35),
		rPadSp(limit(r.recipient.BeneficiaryBankAddress2, 35), 35),
		rPadSp(limit(r.recipient.BeneficiaryBankAddress3, 35), 35),
		rPadSp(r.recipient.BeneficiarySWIFT, 11),
		r.transactionCosts,
		boolToYesNoFlag(r.receiveTransactionDetails),
		rPadSp(limit(r.note1, 35), 35),
		rPadSp(limit(r.note2, 35), 35),
		rPadSp("", 35), rPadSp("", 35), // Notes 3-4
		rPadSp("", 15),                                 // Sorting code
		rPadSp("", 11),                                 // Receiver’s Correspondent Swift Code
		rPadSp("", 35),                                 // Receiver’s Correspondent Bank Name
		rPadSp("", 35), rPadSp("", 35), rPadSp("", 35), // Receiver’s Correspondent Address 1-3
		rPadSp("", 11),                                 // Third Reimbursement Institution Swift Code
		rPadSp("", 35),                                 // Third Reimbursement Institution Bank Name
		rPadSp("", 35), rPadSp("", 35), rPadSp("", 35), // Third Reimbursement Institution Address 1-3
		rPadSp(r.recipient.IntermediarySWIFT, 11),
	)
}
