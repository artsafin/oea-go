package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/common"
	"strings"
	"time"
)

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

	return strings.Compare(invs[i].Employee, invs[j].Employee) < 0
}

// Swap swaps the elements with indexes i and j.
func (invs Invoices) Swap(i, j int) {
	invs[i], invs[j] = invs[j], invs[i]
}

type Invoice struct {
	Id                   string
	InvoiceNo            string
	Month                string
	Employee             string
	PreviousInvoiceId    string
	AmountRub            common.MoneyRub
	EurRubExpected       common.MoneyRub
	RequestedSubtotalEur common.MoneyEur
	RoundingPrevMonEur   common.MoneyEur
	Rounding             common.MoneyEur
	AmountRequestedEur   common.MoneyEur
	Hours                uint16
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
	Corrections          []*Correction
	MonthData            *Month
	PrevInvoice          *Invoice
}

func (inv *Invoice) FullMonthName() string {
	return inv.MonthData.LastMonthDay.Format("January")
}

func (inv *Invoice) InvoiceDate() string {
	return fmt.Sprintf("%02d %s", time.Now().Day(), inv.MonthData.LastMonthDay.Format("Jan 2006"))
}

func (inv *Invoice) PaymentDate() string {
	return fmt.Sprintf("%02d %s", time.Now().Day(), inv.MonthData.LastMonthDay.Format("Jan 2006"))
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
	if invoice.Employee, err = common.ToString(Ids.Invoices.Cols.Employee, row); err != nil {
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
	if invoice.Hours, err = common.ToUint16(Ids.Invoices.Cols.Hours, row); err != nil {
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

	if len(errs) > 0 {
		stringErr := ""
		for _, err := range errs {
			stringErr += err.Error() + "; "
		}

		panic(stringErr)
	}

	return &invoice
}
