package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/employee/codaschema"
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

	if d.Name, err = codatypes.ToString(codaschema.ID.Table.BeneficiaryBank.Cols.Name.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.Address1, err = codatypes.ToString(codaschema.ID.Table.BeneficiaryBank.Cols.Address1.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.Address2, err = codatypes.ToString(codaschema.ID.Table.BeneficiaryBank.Cols.Address2.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.Address3, err = codatypes.ToString(codaschema.ID.Table.BeneficiaryBank.Cols.Address3.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.BeneficiarySWIFT, err = codatypes.ToString(codaschema.ID.Table.BeneficiaryBank.Cols.BeneficiarySWIFT.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.IntermediarySWIFT, err = codatypes.ToString(codaschema.ID.Table.BeneficiaryBank.Cols.IntermediarySWIFT.ID, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return d
}
