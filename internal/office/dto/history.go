package dto

import "oea-go/internal/common"

type GrandTotal struct {
	ExpensesRub   common.MoneyRub
	ActuallySpent common.MoneyRub
	PendingSpend  common.MoneyRub
	Balance       common.MoneyRub
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
