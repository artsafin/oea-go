package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/office/codaschema"
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
	number           uint16
	Date             *time.Time
	hourRate         codatypes.MoneyEur
	EurFixedRate     codatypes.MoneyRub
	EurRateWorst     codatypes.MoneyRub
	ReturnOfRounding codatypes.MoneyEur
	Subtotal         codatypes.MoneyEur
	HourRateRounding codatypes.MoneyEur
	totalEur         codatypes.MoneyEur
	hours            uint16
	filename         string
	ExpensesRub      codatypes.MoneyRub
	ExpensesEur      codatypes.MoneyEur
	ActuallySpent    codatypes.MoneyRub
	PendingSpend     codatypes.MoneyRub
	Balance          codatypes.MoneyRub
	Expenses         []*Expense
	PrevInvoiceID    string
	ApprovalLink     string
}

func (i Invoice) Filename() string {
	return i.filename
}

func (i Invoice) BeneficiaryRequisites() string {
	return ""
}

func (i Invoice) PayerRequisites() string {
	return ""
}

func (i Invoice) BeneficiaryName() string {
	return ""
}

func (i Invoice) PayerName() string {
	return ""
}

func (i Invoice) Number() string {
	return fmt.Sprint(i.number)
}
func (i Invoice) HourRate() codatypes.MoneyEur {
	return i.hourRate
}
func (i Invoice) Hours() uint16 {
	return i.hours
}
func (i Invoice) TotalEur() codatypes.MoneyEur {
	return i.totalEur
}

func (i Invoice) DateYm() string {
	return i.Date.Format("January 2006")
}

func (i Invoice) DateFull() string {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	return i.Date.Format("02.01.2006")
}

func (i *Invoice) HasPendingSpends() bool {
	return i.PendingSpend != 0
}

func (i *Invoice) ContractNumber() string {
	return "1"
}

func (i *Invoice) ContractDate() string {
	return "..."
}

func NewInvoiceFromRow(row *coda.Row) *Invoice {
	invoice := Invoice{}
	errs := codatypes.NewErrorContainer()
	var err error

	if invoice.No, err = codatypes.ToString(codaschema.ID.Table.Invoices.Cols.No.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.Status, err = codatypes.ToString(codaschema.ID.Table.Invoices.Cols.Status.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.number, err = codatypes.ToUint16(codaschema.ID.Table.Invoices.Cols.Number.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.Date, err = codatypes.ToDate(codaschema.ID.Table.Invoices.Cols.Date.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.hourRate, err = codatypes.ToEur(codaschema.ID.Table.Invoices.Cols.HourRate.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EurFixedRate, err = codatypes.ToRub(codaschema.ID.Table.Invoices.Cols.EURFixedRate.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EurRateWorst, err = codatypes.ToRub(codaschema.ID.Table.Invoices.Cols.EURRateWorst.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.ReturnOfRounding, err = codatypes.ToEur(codaschema.ID.Table.Invoices.Cols.ReturnOfRounding.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.Subtotal, err = codatypes.ToEur(codaschema.ID.Table.Invoices.Cols.Subtotal.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.HourRateRounding, err = codatypes.ToEur(codaschema.ID.Table.Invoices.Cols.HourRateRounding.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.totalEur, err = codatypes.ToEur(codaschema.ID.Table.Invoices.Cols.TotalEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.hours, err = codatypes.ToUint16(codaschema.ID.Table.Invoices.Cols.Hours.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.filename, err = codatypes.ToString(codaschema.ID.Table.Invoices.Cols.Filename.ID, row); err != nil {
		errs.AddError(err)
	}

	if invoice.ExpensesRub, err = codatypes.ToRub(codaschema.ID.Table.Invoices.Cols.ExpensesRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.ExpensesEur, err = codatypes.ToEur(codaschema.ID.Table.Invoices.Cols.ExpensesEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.ActuallySpent, err = codatypes.ToRub(codaschema.ID.Table.Invoices.Cols.ActuallySpent.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.PendingSpend, err = codatypes.ToRub(codaschema.ID.Table.Invoices.Cols.PendingSpend.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.Balance, err = codatypes.ToRub(codaschema.ID.Table.Invoices.Cols.Balance.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.PrevInvoiceID, err = codatypes.ToString(codaschema.ID.Table.Invoices.Cols.PrevInvoice.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.ApprovalLink, err = codatypes.ToString(codaschema.ID.Table.Invoices.Cols.ApprovalLink.ID, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &invoice
}
