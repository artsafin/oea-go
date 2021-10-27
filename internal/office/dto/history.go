package dto

import (
	"oea-go/internal/codatypes"
)

type GrandTotal struct {
	ExpensesRub   codatypes.MoneyRub
	ActuallySpent codatypes.MoneyRub
	PendingSpend  codatypes.MoneyRub
	Balance       codatypes.MoneyRub
}

func (t *GrandTotal) AddInvoice(inv *Invoice) {
	t.ExpensesRub += inv.ExpensesRub
	t.ActuallySpent += inv.ActuallySpent
	t.PendingSpend += inv.PendingSpend
	t.Balance += inv.Balance
}

type History struct {
	FirstInvoice *Invoice
	LastInvoice  *Invoice
	Invoices     []*Invoice
	GrandTotal   GrandTotal
}
