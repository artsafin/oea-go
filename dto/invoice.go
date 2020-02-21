package dto

import (
	"time"

	"github.com/phouse512/go-coda"
)

type Invoices []*Invoice

func (invs Invoices) Len() int {
	return len(invs)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (invs Invoices) Less(i, j int) bool {
	return invs[i].Date.Unix() < invs[j].Date.Unix()
}

// Swap swaps the elements with indexes i and j.
func (invs Invoices) Swap(i, j int) {
	invs[i], invs[j] = invs[j], invs[i]
}

type Invoice struct {
	No               string
	Status           string
	Number           uint16
	Date             *time.Time
	HourRate         MoneyEur
	EurFixedRate     MoneyRub
	EurRateWorst     MoneyRub
	ReturnOfRounding MoneyEur
	Subtotal         MoneyEur
	HourRateRounding MoneyEur
	TotalEur         MoneyEur
	Hours            uint16
	Filename         string
	ExpensesRub      MoneyRub
	ExpensesEur      MoneyEur
	ActuallySpent    MoneyRub
	PendingSpend     MoneyRub
	Balance          MoneyRub
	Expenses         []*Expense
}

func (i *Invoice) InvoiceDateYm() string {
	return i.Date.Format("January 2006")
}

func (i *Invoice) InvoiceDateFull() string {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	return i.Date.Format("02.01.2006")
}

func (i *Invoice) HasPendingSpends() bool {
	return i.PendingSpend != 0
}

func NewInvoiceFromRow(row *coda.Row) *Invoice {
	invoice := Invoice{}
	errs := make([]UnexpectedFieldTypeErr, 0)
	var err *UnexpectedFieldTypeErr

	if invoice.No, err = toString(Ids.Invoices.Cols.No, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Status, err = toString(Ids.Invoices.Cols.Status, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Number, err = toUint16(Ids.Invoices.Cols.Number, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Date, err = toDate(Ids.Invoices.Cols.Date, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.HourRate, err = toEur(Ids.Invoices.Cols.HourRate, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.EurFixedRate, err = toRub(Ids.Invoices.Cols.EurFixedRate, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.EurRateWorst, err = toRub(Ids.Invoices.Cols.EurRateWorst, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.ReturnOfRounding, err = toEur(Ids.Invoices.Cols.ReturnOfRounding, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Subtotal, err = toEur(Ids.Invoices.Cols.Subtotal, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.HourRateRounding, err = toEur(Ids.Invoices.Cols.HourRateRounding, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.TotalEur, err = toEur(Ids.Invoices.Cols.TotalEur, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Hours, err = toUint16(Ids.Invoices.Cols.Hours, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Filename, err = toString(Ids.Invoices.Cols.Filename, row); err != nil {
		errs = append(errs, *err)
	}

	if invoice.ExpensesRub, err = toRub(Ids.Invoices.Cols.ExpensesRub, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.ExpensesEur, err = toEur(Ids.Invoices.Cols.ExpensesEur, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.ActuallySpent, err = toRub(Ids.Invoices.Cols.ActuallySpent, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.PendingSpend, err = toRub(Ids.Invoices.Cols.PendingSpend, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Balance, err = toRub(Ids.Invoices.Cols.Balance, row); err != nil {
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
