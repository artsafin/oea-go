package invoices

import (
	"oea-go/internal/codatypes"
	"oea-go/internal/office/codaschema"
)

type grandTotal struct {
	ExpensesRub   codatypes.MoneyRub
	ActuallySpent codatypes.MoneyRub
	PendingSpend  codatypes.MoneyRub
	Balance       codatypes.MoneyRub
}

func (t *grandTotal) AddInvoice(inv codaschema.Invoices) {
	t.ExpensesRub += inv.ExpensesRUBMoney()
	t.ActuallySpent += inv.ActuallySpentMoney()
	t.PendingSpend += inv.PendingSpendMoney()
	t.Balance += inv.BalanceMoney()
}

type History struct {
	FirstInvoice codaschema.Invoices
	LastInvoice  codaschema.Invoices
	PastInvoices []codaschema.Invoices
	GrandTotal   grandTotal
}
