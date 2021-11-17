package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
	"oea-go/internal/employee/codaschema"
	"strings"
	"time"
)

const paymentDateDaysBeforeEndOfMonth = 2

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
	EmployeeNameSlug     string
	Employee             *Employee
	BankDetailsID        string
	BankDetails          *BankDetails
	PreviousInvoiceId    string
	RequestedSubtotalRub codatypes.MoneyRub
	EurRubExpected       codatypes.MoneyRub
	RequestedSubtotalEur codatypes.MoneyEur
	RoundingPrevMonEur   codatypes.MoneyEur
	Rounding             codatypes.MoneyEur
	AmountRequestedEur   codatypes.MoneyEur
	hours                uint16
	EurRubActual         codatypes.MoneyRub
	AmountRubActual      codatypes.MoneyRub
	RateErrorRub         codatypes.MoneyRub
	CostOfDay            codatypes.MoneyRub
	OpeningDateIp        *time.Time
	CorrectionRub        codatypes.MoneyRub
	PatentRub            codatypes.MoneyRub
	TaxesRub             codatypes.MoneyRub
	BaseSalaryRub        codatypes.MoneyRub
	BaseSalaryEur        codatypes.MoneyEur
	BankFees             codatypes.MoneyRub
	RateErrorPrevMon     codatypes.MoneyRub
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

func (inv *Invoice) HourRate() codatypes.MoneyEur {
	if inv.Employee == nil {
		return codatypes.MoneyEur(0)
	}
	return inv.Employee.HourRate
}

func (inv *Invoice) Hours() uint16 {
	return inv.hours
}

func (inv *Invoice) TotalEur() codatypes.MoneyEur {
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
	now := time.Now()
	if paymentDate.Before(now) {
		paymentDate = now
	}
	return paymentDate.Format("2 Jan 2006")
}

func NewInvoiceFromRow(row *coda.Row) *Invoice {
	invoice := Invoice{}
	errs := codatypes.NewErrorContainer()
	var err error

	if invoice.Id, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.ID.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.InvoiceNo, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.InvoiceHash.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.Month, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.Month.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EmployeeName, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.Employee.ID, row); err != nil {
		errs.AddError(err)
	} else {
		invoice.EmployeeNameSlug = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(invoice.EmployeeName)), " ", "-")
	}
	if invoice.PreviousInvoiceId, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.PreviousInvoice.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.RequestedSubtotalRub, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.RUBSubtotal.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EurRubExpected, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.EURRUBExpected.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.RequestedSubtotalEur, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.RequestedSubtotalEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.RoundingPrevMonEur, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.RoundingPrevMonEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.Rounding, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.Rounding.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.AmountRequestedEur, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.RequestedTotalEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.hours, err = codatypes.ToUint16(codaschema.ID.Table.Invoice.Cols.Hours.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EurRubActual, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.EURRUBActual.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.AmountRubActual, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.AmountRUBActual.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.RateErrorRub, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.RateErrorRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.CostOfDay, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.CostOfDay.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.OpeningDateIp, err = codatypes.ToDate(codaschema.ID.Table.Invoice.Cols.OpeningDateIP.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.CorrectionRub, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.CorrectionRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.PatentRub, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.PatentRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.TaxesRub, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.TaxesRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.BaseSalaryRub, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.BaseSalaryRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.BaseSalaryEur, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.BaseSalaryEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.BankFees, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.BankFees.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.RateErrorPrevMon, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.RateErrorPM.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.PaymentChecksPassed, err = codatypes.ToBool(codaschema.ID.Table.Invoice.Cols.ToPay.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.BankDetailsID, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.BankDetails.ID, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &invoice
}
