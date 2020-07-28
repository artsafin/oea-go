package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/common"
	"strings"
	"time"
)

const paymentDateDaysBeforeEndOfMonth = 3

type Invoices []*Invoice

func (invs Invoices) Len() int {
	return len(invs)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (invs Invoices) Less(i, j int) bool {
	monthCmp := strings.Compare(invs[i].Month, invs[j].Month)
	if monthCmp != 0 {
		return monthCmp < 0
	}

	return strings.Compare(invs[i].EmployeeName, invs[j].EmployeeName) < 0
}

// Swap swaps the elements with indexes i and j.
func (invs Invoices) Swap(i, j int) {
	invs[i], invs[j] = invs[j], invs[i]
}

type Invoice struct {
	Id                   string
	InvoiceNo            string
	Month                string
	EmployeeName         string
	Employee             *Employee
	PreviousInvoiceId    string
	AmountRub            common.MoneyRub
	EurRubExpected       common.MoneyRub
	RequestedSubtotalEur common.MoneyEur
	RoundingPrevMonEur   common.MoneyEur
	Rounding             common.MoneyEur
	AmountRequestedEur   common.MoneyEur
	hours                uint16
	EurRubActual         common.MoneyRub
	AmountRubActual      common.MoneyRub
	RateErrorRub         common.MoneyRub
	CostOfDay            common.MoneyRub
	BankTotalFees        common.MoneyRub
	OpeningDateIp        *time.Time
	UnpaidDay            common.MoneyRub
	CorrectionRub        common.MoneyRub
	PatentRub            common.MoneyRub
	TaxesRub             common.MoneyRub
	BaseSalary           common.MoneyRub
	BankFees             common.MoneyRub
	RateErrorPrevMon     common.MoneyRub
	SelfEmployedTax      common.MoneyRub
	Corrections          []*Correction
	MonthData            *Month
	PrevInvoice          *Invoice
	PaymentChecksPassed  bool
}

func (inv *Invoice) Filename() string {
	nameWithoutSpaces := strings.Replace(inv.EmployeeName, " ", "_", -1)
	nameLower := strings.ToLower(nameWithoutSpaces)

	return fmt.Sprintf("%s_%s.xlsx", inv.Month, nameLower)
}

func (inv *Invoice) BeneficiaryRequisites() string {
	if inv.Employee == nil {
		return "n/a"
	}
	return inv.Employee.BankRequisites
}

func (inv *Invoice) PayerRequisites() string {
	if inv.Employee == nil {
		return "n/a"
	}
	return strings.ReplaceAll(inv.Employee.BillTo, "\n", "\r\n")
}

func (inv *Invoice) BeneficiaryName() string {
	if inv.Employee == nil {
		return "n/a"
	}
	return inv.Employee.EnglishFullName
}

func (inv *Invoice) PayerName() string {
	return ""
}

func (inv *Invoice) Number() string {
	return inv.InvoiceNo
}

func (inv *Invoice) DateYm() string {
	if inv.MonthData == nil {
		return "n/a"
	}
	return inv.MonthData.LastMonthDay.Format("January 2006")
}

func (inv *Invoice) HourRate() common.MoneyEur {
	if inv.Employee == nil {
		return common.MoneyEur(0)
	}
	return inv.Employee.HourRate
}

func (inv *Invoice) Hours() uint16 {
	return inv.hours
}

func (inv *Invoice) TotalEur() common.MoneyEur {
	return inv.AmountRequestedEur
}

func (inv *Invoice) FullMonthName() string {
	if inv.MonthData == nil {
		return "n/a"
	}
	return inv.MonthData.LastMonthDay.Format("January")
}

func (inv *Invoice) DateFull() string {
	if inv.MonthData == nil {
		return "n/a"
	}
	return fmt.Sprintf("%02d %s", time.Now().Day(), inv.MonthData.LastMonthDay.Format("Jan 2006"))
}

func (inv *Invoice) ContractNumber() string {
	return inv.Employee.ContractNumber
}

func (inv *Invoice) ContractDate() string {
	return inv.Employee.ContractDate
}

func (inv *Invoice) DatePayment() string {
	if inv.MonthData == nil {
		return "n/a"
	}
	paymentDate := common.AddWorkingDate(*inv.MonthData.LastMonthDay, 0, 0, -paymentDateDaysBeforeEndOfMonth)
	return paymentDate.Format("2 Jan 2006")
}

func NewInvoiceFromRow(row *coda.Row) *Invoice {
	invoice := Invoice{}
	errs := make([]common.UnexpectedFieldTypeErr, 0)
	var err *common.UnexpectedFieldTypeErr

	if invoice.Id, err = common.ToString(Ids.Invoices.Cols.Id, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.InvoiceNo, err = common.ToString(Ids.Invoices.Cols.InvoiceNo, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Month, err = common.ToString(Ids.Invoices.Cols.Month, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.EmployeeName, err = common.ToString(Ids.Invoices.Cols.Employee, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.PreviousInvoiceId, err = common.ToString(Ids.Invoices.Cols.PreviousInvoice, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.AmountRub, err = common.ToRub(Ids.Invoices.Cols.AmountRub, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.EurRubExpected, err = common.ToRub(Ids.Invoices.Cols.EurRubExpected, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.RequestedSubtotalEur, err = common.ToEur(Ids.Invoices.Cols.RequestedSubtotalEur, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.RoundingPrevMonEur, err = common.ToEur(Ids.Invoices.Cols.RoundingPrevMonEur, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Rounding, err = common.ToEur(Ids.Invoices.Cols.Rounding, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.AmountRequestedEur, err = common.ToEur(Ids.Invoices.Cols.AmountRequestedEur, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.hours, err = common.ToUint16(Ids.Invoices.Cols.Hours, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.EurRubActual, err = common.ToRub(Ids.Invoices.Cols.EurRubActual, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.AmountRubActual, err = common.ToRub(Ids.Invoices.Cols.AmountRubActual, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.RateErrorRub, err = common.ToRub(Ids.Invoices.Cols.RateErrorRub, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.CostOfDay, err = common.ToRub(Ids.Invoices.Cols.CostOfDay, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.BankTotalFees, err = common.ToRub(Ids.Invoices.Cols.BankTotalFees, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.OpeningDateIp, err = common.ToDate(Ids.Invoices.Cols.OpeningDateIp, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.UnpaidDay, err = common.ToRub(Ids.Invoices.Cols.UnpaidDay, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.CorrectionRub, err = common.ToRub(Ids.Invoices.Cols.CorrectionRub, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.PatentRub, err = common.ToRub(Ids.Invoices.Cols.PatentRub, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.TaxesRub, err = common.ToRub(Ids.Invoices.Cols.TaxesRub, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.BaseSalary, err = common.ToRub(Ids.Invoices.Cols.BaseSalary, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.BankFees, err = common.ToRub(Ids.Invoices.Cols.BankFees, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.RateErrorPrevMon, err = common.ToRub(Ids.Invoices.Cols.RateErrorPrevMon, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.PaymentChecksPassed, err = common.ToBool(Ids.Invoices.Cols.PaymentChecksPassed, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.SelfEmployedTax, err = common.ToRub(Ids.Invoices.Cols.SelfEmployedTax, row); err != nil {
		errs = append(errs, *err)
	}

	if len(errs) > 0 {
		stringErr := ""
		for _, err := range errs {
			stringErr += err.Error() + "; "
		}

		panic(stringErr)
	}

	return &invoice
}
