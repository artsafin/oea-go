package invoices

import (
	"oea-go/internal/codatypes"
	"oea-go/internal/office/codaschema"
)

type expenseGroup struct {
	Name     string
	Items    []codaschema.Expenses
	TotalEur codatypes.MoneyEur
	TotalRub codatypes.MoneyRub
}

type ExpenseGroupMap map[string]*expenseGroup

func (e *expenseGroup) Add(expense codaschema.Expenses) {
	e.Items = append(e.Items, expense)
	e.TotalEur += expense.AmountEURMoney()
	e.TotalRub += expense.AmountRUBMoney()
}

func GroupExpensesByCategory(expenses []codaschema.Expenses) (m ExpenseGroupMap) {
	m = make(ExpenseGroupMap)

	for _, expense := range expenses {
		if _, catExists := m[expense.Category]; !catExists {
			m[expense.Category] = &expenseGroup{Name: expense.Category}
		}
		m[expense.Category].Add(expense)
	}

	return
}
