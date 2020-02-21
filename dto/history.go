package dto

type GrandTotal struct {
	ExpensesRub   MoneyRub
	ActuallySpent MoneyRub
	PendingSpend  MoneyRub
	Balance       MoneyRub
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
