package dto

import (
	"github.com/artsafin/go-coda"
	"oea-go/common"
	"time"
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
	HourRate         common.MoneyEur
	EurFixedRate     common.MoneyRub
	EurRateWorst     common.MoneyRub
	ReturnOfRounding common.MoneyEur
	Subtotal         common.MoneyEur
	HourRateRounding common.MoneyEur
	TotalEur         common.MoneyEur
	Hours            uint16
	Filename         string
	ExpensesRub      common.MoneyRub
	ExpensesEur      common.MoneyEur
	ActuallySpent    common.MoneyRub
	PendingSpend     common.MoneyRub
	Balance          common.MoneyRub
	Expenses         []*Expense
	PrevInvoiceID    string
	ApprovalLink     string
}

func (i Invoice) InvoiceDateYm() string {
	return i.Date.Format("January 2006")
}

func (i Invoice) InvoiceDateFull() string {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	return i.Date.Format("02.01.2006")
}

func (i *Invoice) HasPendingSpends() bool {
	return i.PendingSpend != 0
}

func NewInvoiceFromRow(row *coda.Row) *Invoice {
	invoice := Invoice{}
	errs := make([]common.UnexpectedFieldTypeErr, 0)
	var err *common.UnexpectedFieldTypeErr

	if invoice.No, err = common.ToString(Ids.Invoices.Cols.No, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Status, err = common.ToString(Ids.Invoices.Cols.Status, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Number, err = common.ToUint16(Ids.Invoices.Cols.Number, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Date, err = common.ToDate(Ids.Invoices.Cols.Date, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.HourRate, err = common.ToEur(Ids.Invoices.Cols.HourRate, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.EurFixedRate, err = common.ToRub(Ids.Invoices.Cols.EurFixedRate, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.EurRateWorst, err = common.ToRub(Ids.Invoices.Cols.EurRateWorst, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.ReturnOfRounding, err = common.ToEur(Ids.Invoices.Cols.ReturnOfRounding, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Subtotal, err = common.ToEur(Ids.Invoices.Cols.Subtotal, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.HourRateRounding, err = common.ToEur(Ids.Invoices.Cols.HourRateRounding, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.TotalEur, err = common.ToEur(Ids.Invoices.Cols.TotalEur, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Hours, err = common.ToUint16(Ids.Invoices.Cols.Hours, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Filename, err = common.ToString(Ids.Invoices.Cols.Filename, row); err != nil {
		errs = append(errs, *err)
	}

	if invoice.ExpensesRub, err = common.ToRub(Ids.Invoices.Cols.ExpensesRub, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.ExpensesEur, err = common.ToEur(Ids.Invoices.Cols.ExpensesEur, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.ActuallySpent, err = common.ToRub(Ids.Invoices.Cols.ActuallySpent, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.PendingSpend, err = common.ToRub(Ids.Invoices.Cols.PendingSpend, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.Balance, err = common.ToRub(Ids.Invoices.Cols.Balance, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.PrevInvoiceID, err = common.ToString(Ids.Invoices.Cols.PrevInvoice, row); err != nil {
		errs = append(errs, *err)
	}
	if invoice.ApprovalLink, err = common.ToString(Ids.Invoices.Cols.ApprovalLink, row); err != nil {
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
