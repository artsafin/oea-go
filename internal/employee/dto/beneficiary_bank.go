package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
)

type BeneficiaryBank struct {
	RowID             string
	Name              string
	Address1          string
	Address2          string
	Address3          string
	BeneficiarySWIFT  string
	IntermediarySWIFT string
}

func NewBeneficiaryBankFromRow(row *coda.Row) BeneficiaryBank {
	d := BeneficiaryBank{}
	errs := codatypes.NewErrorContainer()
	var err error

	d.RowID = row.Id

	if d.Name, err = codatypes.ToString(Ids.BeneficiaryBank.Cols.Name, row); err != nil {
		errs.AddError(err)
	}
	if d.Address1, err = codatypes.ToString(Ids.BeneficiaryBank.Cols.Address1, row); err != nil {
		errs.AddError(err)
	}
	if d.Address2, err = codatypes.ToString(Ids.BeneficiaryBank.Cols.Address2, row); err != nil {
		errs.AddError(err)
	}
	if d.Address3, err = codatypes.ToString(Ids.BeneficiaryBank.Cols.Address3, row); err != nil {
		errs.AddError(err)
	}
	if d.BeneficiarySWIFT, err = codatypes.ToString(Ids.BeneficiaryBank.Cols.BeneficiarySWIFT, row); err != nil {
		errs.AddError(err)
	}
	if d.IntermediarySWIFT, err = codatypes.ToString(Ids.BeneficiaryBank.Cols.IntermediarySWIFT, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return d
}
