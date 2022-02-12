package codaschema

import "oea-go/internal/codatypes"

func (inv Invoices) HasPendingSpends() bool {
	return inv.PendingSpendMoney() != 0
}

func (inv Invoices) HourRateMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.HourRate * 100)
}

func (inv Invoices) EURFixedRateMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.EURFixedRate * 100)
}

func (inv Invoices) EURRateWorstMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.EURRateWorst * 100)
}

func (inv Invoices) ExpensesRUBMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.ExpensesRUB * 100)
}

func (inv Invoices) ExpensesEURMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.ExpensesEUR * 100)
}

func (inv Invoices) ReturnOfRoundingMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.ReturnOfRounding * 100)
}

func (inv Invoices) SubtotalMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.Subtotal * 100)
}

func (inv Invoices) HourRateRoundingMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.HourRateRounding * 100)
}

func (inv Invoices) TotalEURMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.TotalEUR * 100)
}

func (inv Invoices) ActuallySpentMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.ActuallySpent * 100)
}

func (inv Invoices) PendingSpendMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.PendingSpend * 100)
}

func (inv Invoices) BalanceMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.Balance * 100)
}
