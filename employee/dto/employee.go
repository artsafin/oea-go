package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/common"
	"time"
)

type Employee struct {
	Location                string
	Name                    string
	StartDate               *time.Time
	ProbationEnd            *time.Time
	AnnualLeaveFrom         *time.Time
	SalaryBeforeIp          common.MoneyRub
	SalaryAfterIp           common.MoneyRub
	NetSalaryAfterProbation common.MoneyRub
	EndDate                 *time.Time
	HourRate                common.MoneyEur
	BankCurrencyControl     common.MoneyRub
	BankService             common.MoneyRub
	BankTotalFees           common.MoneyRub
	Address                 string
	OpeningDateIp           *time.Time
	StartMonthName          string
	MattermostLogin         string
	RussianFullName         string
	Position                string
	INN                     string
	IsWorkingNow            bool
	PatentFee               common.MoneyRub
	LastSalary              common.MoneyRub
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
	errs := make([]common.UnexpectedFieldTypeErr, 0)
	var err *common.UnexpectedFieldTypeErr

	if empl.Location, err = common.ToString(Ids.Employees.Cols.Location, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.Name, err = common.ToString(Ids.Employees.Cols.Name, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.StartDate, err = common.ToDate(Ids.Employees.Cols.StartDate, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.ProbationEnd, err = common.ToDate(Ids.Employees.Cols.ProbationEnd, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.AnnualLeaveFrom, err = common.ToDate(Ids.Employees.Cols.AnnualLeaveFrom, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.SalaryBeforeIp, err = common.ToRub(Ids.Employees.Cols.SalaryBeforeIp, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.SalaryAfterIp, err = common.ToRub(Ids.Employees.Cols.SalaryAfterIp, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.NetSalaryAfterProbation, err = common.ToRub(Ids.Employees.Cols.NetSalaryAfterProbation, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.EndDate, err = common.ToDate(Ids.Employees.Cols.EndDate, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.HourRate, err = common.ToEur(Ids.Employees.Cols.HourRate, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.BankCurrencyControl, err = common.ToRub(Ids.Employees.Cols.BankCurrencyControl, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.BankService, err = common.ToRub(Ids.Employees.Cols.BankService, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.BankTotalFees, err = common.ToRub(Ids.Employees.Cols.BankTotalFees, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.Address, err = common.ToString(Ids.Employees.Cols.Address, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.OpeningDateIp, err = common.ToDate(Ids.Employees.Cols.OpeningDateIp, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.StartMonthName, err = common.ToString(Ids.Employees.Cols.StartMonth, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.MattermostLogin, err = common.ToString(Ids.Employees.Cols.MattermostLogin, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.RussianFullName, err = common.ToString(Ids.Employees.Cols.RussianFullName, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.Position, err = common.ToString(Ids.Employees.Cols.Position, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.INN, err = common.ToString(Ids.Employees.Cols.INN, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.IsWorkingNow, err = common.ToBool(Ids.Employees.Cols.IsWorkingNow, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.PatentFee, err = common.ToRub(Ids.Employees.Cols.PatentFee, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.LastSalary, err = common.ToRub(Ids.Employees.Cols.LastSalary, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.EnglishFullName, err = common.ToString(Ids.Employees.Cols.EnglishFullName, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.BankRequisites, err = common.ToString(Ids.Employees.Cols.BankRequisites, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.BillTo, err = common.ToString(Ids.Employees.Cols.BillTo, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.GeneralSdLink, err = common.ToString(Ids.Employees.Cols.GeneralSdLink, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.PersonnelSdLink, err = common.ToString(Ids.Employees.Cols.PersonnelSdLink, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.FinanceSdLink, err = common.ToString(Ids.Employees.Cols.FinanceSdLink, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.AclSdLink, err = common.ToString(Ids.Employees.Cols.AclSdLink, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.ContractDate, err = common.ToString(Ids.Employees.Cols.ContractDate, row); err != nil {
		errs = append(errs, *err)
	}
	if empl.ContractNumber, err = common.ToString(Ids.Employees.Cols.ContractNumber, row); err != nil {
		errs = append(errs, *err)
	}

	if len(errs) > 0 {
		stringErr := ""
		for _, err := range errs {
			stringErr += err.Error() + "; "
		}

		panic(stringErr)
	}

	return &empl
}
