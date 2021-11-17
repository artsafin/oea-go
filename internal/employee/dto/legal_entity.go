package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/employee/codaschema"
)

type LegalEntity struct {
	RowID         string
	EntityName    string
	OfficialName  string
	AccountNumber string
}

func NewLegalEntityFromRow(row *coda.Row) LegalEntity {
	d := LegalEntity{}
	errs := codatypes.NewErrorContainer()
	var err error

	d.RowID = row.Id

	if d.EntityName, err = codatypes.ToString(codaschema.ID.Table.LegalEntity.Cols.EntityName.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.OfficialName, err = codatypes.ToString(codaschema.ID.Table.LegalEntity.Cols.OfficialName.ID, row); err != nil {
		errs.AddError(err)
	}
	if d.AccountNumber, err = codatypes.ToString(codaschema.ID.Table.LegalEntity.Cols.AccountNumber.ID, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return d
}
