package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
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
	PatentFee               codatypes.MoneyRub
	LastSalary              codatypes.MoneyRub
	EnglishFullName         string
	BankRequisites          string
	BillTo                  string
	GeneralSdLink           string
	PersonnelSdLink         string
	FinanceSdLink           string
	AclSdLink               string
	ContractDate            string
	ContractNumber          string
}

func NewEmployeeFromRow(row *coda.Row) *Employee {
	empl := Employee{}
	errs := codatypes.NewErrorContainer()
	var err error

	if empl.Location, err = codatypes.ToString(Ids.Employees.Cols.Location, row); err != nil {
		errs.AddError(err)
	}
	if empl.Name, err = codatypes.ToString(Ids.Employees.Cols.Name, row); err != nil {
		errs.AddError(err)
	}
	if empl.StartDate, err = codatypes.ToDate(Ids.Employees.Cols.StartDate, row); err != nil {
		errs.AddError(err)
	}
	if empl.ProbationEnd, err = codatypes.ToDate(Ids.Employees.Cols.ProbationEnd, row); err != nil {
		errs.AddError(err)
	}
	if empl.AnnualLeaveFrom, err = codatypes.ToDate(Ids.Employees.Cols.AnnualLeaveFrom, row); err != nil {
		errs.AddError(err)
	}
	if empl.SalaryBeforeIp, err = codatypes.ToRub(Ids.Employees.Cols.SalaryBeforeIp, row); err != nil {
		errs.AddError(err)
	}
	if empl.SalaryAfterIp, err = codatypes.ToRub(Ids.Employees.Cols.SalaryAfterIp, row); err != nil {
		errs.AddError(err)
	}
	if empl.NetSalaryAfterProbation, err = codatypes.ToRub(Ids.Employees.Cols.NetSalaryAfterProbation, row); err != nil {
		errs.AddError(err)
	}
	if empl.EndDate, err = codatypes.ToDate(Ids.Employees.Cols.EndDate, row); err != nil {
		errs.AddError(err)
	}
	if empl.HourRate, err = codatypes.ToEur(Ids.Employees.Cols.HourRate, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankCurrencyControl, err = codatypes.ToRub(Ids.Employees.Cols.BankCurrencyControl, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankService, err = codatypes.ToRub(Ids.Employees.Cols.BankService, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankTotalFees, err = codatypes.ToRub(Ids.Employees.Cols.BankTotalFees, row); err != nil {
		errs.AddError(err)
	}
	if empl.Address, err = codatypes.ToString(Ids.Employees.Cols.Address, row); err != nil {
		errs.AddError(err)
	}
	if empl.OpeningDateIp, err = codatypes.ToDate(Ids.Employees.Cols.OpeningDateIp, row); err != nil {
		errs.AddError(err)
	}
	if empl.StartMonthName, err = codatypes.ToString(Ids.Employees.Cols.StartMonth, row); err != nil {
		errs.AddError(err)
	}
	if empl.MattermostLogin, err = codatypes.ToString(Ids.Employees.Cols.MattermostLogin, row); err != nil {
		errs.AddError(err)
	}
	if empl.RussianFullName, err = codatypes.ToString(Ids.Employees.Cols.RussianFullName, row); err != nil {
		errs.AddError(err)
	}
	if empl.Position, err = codatypes.ToString(Ids.Employees.Cols.Position, row); err != nil {
		errs.AddError(err)
	}
	if empl.INN, err = codatypes.ToString(Ids.Employees.Cols.INN, row); err != nil {
		errs.AddError(err)
	}
	if empl.IsWorkingNow, err = codatypes.ToBool(Ids.Employees.Cols.IsWorkingNow, row); err != nil {
		errs.AddError(err)
	}
	if empl.PatentFee, err = codatypes.ToRub(Ids.Employees.Cols.PatentFee, row); err != nil {
		errs.AddError(err)
	}
	if empl.LastSalary, err = codatypes.ToRub(Ids.Employees.Cols.LastSalary, row); err != nil {
		errs.AddError(err)
	}
	if empl.EnglishFullName, err = codatypes.ToString(Ids.Employees.Cols.EnglishFullName, row); err != nil {
		errs.AddError(err)
	}
	if empl.BankRequisites, err = codatypes.ToString(Ids.Employees.Cols.BankRequisites, row); err != nil {
		errs.AddError(err)
	}
	if empl.BillTo, err = codatypes.ToString(Ids.Employees.Cols.BillTo, row); err != nil {
		errs.AddError(err)
	}
	if empl.GeneralSdLink, err = codatypes.ToString(Ids.Employees.Cols.GeneralSdLink, row); err != nil {
		errs.AddError(err)
	}
	if empl.PersonnelSdLink, err = codatypes.ToString(Ids.Employees.Cols.PersonnelSdLink, row); err != nil {
		errs.AddError(err)
	}
	if empl.FinanceSdLink, err = codatypes.ToString(Ids.Employees.Cols.FinanceSdLink, row); err != nil {
		errs.AddError(err)
	}
	if empl.AclSdLink, err = codatypes.ToString(Ids.Employees.Cols.AclSdLink, row); err != nil {
		errs.AddError(err)
	}
	if empl.ContractDate, err = codatypes.ToString(Ids.Employees.Cols.ContractDate, row); err != nil {
		errs.AddError(err)
	}
	if empl.ContractNumber, err = codatypes.ToString(Ids.Employees.Cols.ContractNumber, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &empl
}
