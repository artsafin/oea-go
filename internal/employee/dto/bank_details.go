package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
)

type BankDetails struct {
	ID       string
	Account  string
	Address1 string
	Address2 string
	Bank     BeneficiaryBank
}

func NewBankDetailsFromRow(row *coda.Row) BankDetails {
	d := BankDetails{
		Bank: BeneficiaryBank{},
	}
	errs := codatypes.NewErrorContainer()
	var err error

	if d.ID, err = codatypes.ToString(Ids.BankDetails.Cols.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.Account, err = codatypes.ToString(Ids.BankDetails.Cols.Account, row); err != nil {
		errs.AddError(err)
	}
	if d.Address1, err = codatypes.ToString(Ids.BankDetails.Cols.Address1, row); err != nil {
		errs.AddError(err)
	}
	if d.Address2, err = codatypes.ToString(Ids.BankDetails.Cols.Address2, row); err != nil {
		errs.AddError(err)
	}

	var benBankValue codatypes.StructuredValue
	if benBankValue, err = codatypes.ToStructuredValue(Ids.BankDetails.Cols.BeneficiaryBank, row); err != nil {
		errs.AddError(err)
	}
	d.Bank.RowID = benBankValue.RowId

	errs.PanicIfError()

	return d
}
