package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
	"oea-go/internal/employee/codaschema"
)

type Employee struct {
	Name            string
	ContractDate    string
	ContractNumber  string
	Location        string
	LegalEntityName string
	LegalEntity     *LegalEntity
	EnglishFullName string
	GeneralSdLink   string
	FinanceSdLink   string
}

func NewEmployeeFromRow(row *coda.Row) (*Employee, error) {
	empl := Employee{}
	errs := codatypes.NewErrorContainer()
	var err error

	if empl.Location, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.Location.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.Name, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.Name.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.EnglishFullName, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.EnglishFullName.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.ContractDate, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.ContractDate.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.ContractNumber, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.ContractNumber.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.LegalEntityName, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.LegalEntity.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.GeneralSdLink, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.GeneralSD.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.FinanceSdLink, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.FinanceSD.ID, row); err != nil {
		errs.AddError(err)
	}

	return &empl, common.JoinErrors(errs)
}
