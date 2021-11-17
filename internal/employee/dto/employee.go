package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/employee/codaschema"
	"time"
)

type Employee struct {
	Location                string
	Name                    string
	StartDate               *time.Time
	ProbationEnd            *time.Time
	AnnualLeaveFrom         *time.Time
	SalaryBeforeIp          codatypes.MoneyRub
	SalaryAfterIp           codatypes.MoneyRub
	NetSalaryAfterProbation codatypes.MoneyRub
	EndDate                 *time.Time
	HourRate                codatypes.MoneyEur
	BankCurrencyControl     codatypes.MoneyRub
	BankService             codatypes.MoneyRub
	BankTotalFees           codatypes.MoneyRub
	Address                 string
	OpeningDateIp           *time.Time
	StartMonthName          string
	MattermostLogin         string
	RussianFullName         string
	Position                string
	INN                     string
	IsWorkingNow            bool
	EnglishFullName         string
	BankRequisites          string
	BillTo                  string
	GeneralSdLink           string
	PersonnelSdLink         string
	FinanceSdLink           string
	AclSdLink               string
	ContractDate            string
	ContractNumber          string
	LegalEntityName         string
	LegalEntity             *LegalEntity
}

func NewEmployeeFromRow(row *coda.Row) *Employee {
	empl := Employee{}
	errs := codatypes.NewErrorContainer()
	var err error

	if empl.Location, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.Location.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.Name, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.Name.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.StartDate, err = codatypes.ToDate(codaschema.ID.Table.AllEmployees.Cols.StartDate.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.ProbationEnd, err = codatypes.ToDate(codaschema.ID.Table.AllEmployees.Cols.ProbationEnd.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.AnnualLeaveFrom, err = codatypes.ToDate(codaschema.ID.Table.AllEmployees.Cols.AnnualLeaveFrom.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.SalaryBeforeIp, err = codatypes.ToRub(codaschema.ID.Table.AllEmployees.Cols.SalaryBeforeIP.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.SalaryAfterIp, err = codatypes.ToRub(codaschema.ID.Table.AllEmployees.Cols.SalaryAfterIP.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.NetSalaryAfterProbation, err = codatypes.ToRub(codaschema.ID.Table.AllEmployees.Cols.NetSalaryAfterProbation.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.EndDate, err = codatypes.ToDate(codaschema.ID.Table.AllEmployees.Cols.EndDate.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.HourRate, err = codatypes.ToEur(codaschema.ID.Table.AllEmployees.Cols.HourRate.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankCurrencyControl, err = codatypes.ToRub(codaschema.ID.Table.AllEmployees.Cols.BankCC.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankService, err = codatypes.ToRub(codaschema.ID.Table.AllEmployees.Cols.BankService.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankTotalFees, err = codatypes.ToRub(codaschema.ID.Table.AllEmployees.Cols.BankTotalFees.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.Address, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.Address.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.OpeningDateIp, err = codatypes.ToDate(codaschema.ID.Table.AllEmployees.Cols.OpeningDateIP.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.StartMonthName, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.StartMonth.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.MattermostLogin, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.MattermostLogin.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.RussianFullName, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.RussianFullName.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.Position, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.Position.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.INN, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.INN.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.IsWorkingNow, err = codatypes.ToBool(codaschema.ID.Table.AllEmployees.Cols.WorkingNow.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.EnglishFullName, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.EnglishFullName.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankRequisites, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.BankRequisites.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.BillTo, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.BillTo.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.GeneralSdLink, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.GeneralSD.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.PersonnelSdLink, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.PersonnelSD.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.FinanceSdLink, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.FinanceSD.ID, row); err != nil {
		errs.AddError(err)
	}
	if empl.AclSdLink, err = codatypes.ToString(codaschema.ID.Table.AllEmployees.Cols.ACLSD.ID, row); err != nil {
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

	errs.PanicIfError()

	return &empl
}
