package txt

import (
	"oea-go/internal/codatypes"
	"testing"
)
import "github.com/stretchr/testify/assert"

func Test_worldwideBankRecord_String(t *testing.T) {
	type fields struct {
		debitAccount              string
		submissionDate            string
		valueDate                 string
		transactionCosts          transactionCosts
		receiveTransactionDetails bool
		recipient                 RecipientBankDetails
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple",
			fields: fields{
				debitAccount:              "DebitAccount",
				submissionDate:            "20060102",
				valueDate:                 "20060103",
				transactionCosts:          TransactionCostsOurs,
				receiveTransactionDetails: true,
				recipient: RecipientBankDetails{
					Account:                 "RecipientAccount",
					Amount:                  codatypes.MoneyEur(773133798),
					BeneficiaryName:         "RecipientBeneficiaryName",
					BeneficiaryAddress1:     "RecipientBeneficiaryAddress1",
					BeneficiaryAddress2:     "RecipientBeneficiaryAddress2",
					BeneficiaryBankName:     "RecipientBeneficiaryBankName",
					BeneficiaryBankAddress1: "RecipientBeneficiaryBankAddress1",
					BeneficiaryBankAddress2: "RecipientBeneficiaryBankAddress2",
					BeneficiaryBankAddress3: "RecipientBeneficiaryBankAddress3",
					BeneficiarySWIFT:        "RecipientBeneficiarySWIFT",
					IntermediarySWIFT:       "RecipientIntermediarySWIFT",
				},
			},
			want: "I|" +
				"DebitAccount|RecipientAccount|7.731.337,98|EUR|20060102|20060103|" +
				"RecipientBeneficiaryName|RecipientBeneficiaryAddress1|RecipientBeneficiaryAddress2|" +
				"RecipientBeneficiaryBankName|RecipientBeneficiaryBankAddress1|RecipientBeneficiaryBankAddress2|RecipientBeneficiaryBankAddress3|" +
				"RecipientBeneficiarySWIFT|" +
				"O|Y|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|" +
				"RecipientIntermediarySWIFT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &worldwideBankRecord{
				debitAccount:              tt.fields.debitAccount,
				submissionDate:            tt.fields.submissionDate,
				valueDate:                 tt.fields.valueDate,
				transactionCosts:          tt.fields.transactionCosts,
				receiveTransactionDetails: tt.fields.receiveTransactionDetails,
				recipient:                 tt.fields.recipient,
			}
			got := r.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
