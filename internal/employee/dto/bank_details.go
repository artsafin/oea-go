package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
	"oea-go/internal/employee/codaschema"
)

type BankDetails struct {
	ID             string
	BankRequisites string
	Account        string
	Address1       string
	Address2       string
	Bank           BeneficiaryBank
}

func NewBankDetailsFromRow(row *coda.Row) (BankDetails, error) {
	d := BankDetails{
		Bank: BeneficiaryBank{},
	}
	errs := codatypes.NewErrorContainer()
	var err error

	if d.ID, err = codatypes.ToString(codaschema.ID.Table.BankDetails.Cols.ID.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.BankRequisites, err = codatypes.ToString(codaschema.ID.Table.BankDetails.Cols.BankRequisites.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.Account, err = codatypes.ToString(codaschema.ID.Table.BankDetails.Cols.Account.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.Address1, err = codatypes.ToString(codaschema.ID.Table.BankDetails.Cols.Address1.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.Address2, err = codatypes.ToString(codaschema.ID.Table.BankDetails.Cols.Address2.ID, row); err != nil {
		errs.AddError(err)
	}

	var benBankValue codatypes.StructuredValue
	if benBankValue, err = codatypes.ToStructuredValue(codaschema.ID.Table.BankDetails.Cols.BeneficiaryBank.ID, row); err != nil {
		errs.AddError(err)
	}
	d.Bank.RowID = benBankValue.RowId

	return d, common.JoinErrors(errs)
}
