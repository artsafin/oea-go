package dto

import (
	"fmt"
	"oea-go/internal/codatypes"
	"time"

	"github.com/artsafin/go-coda"
)

type Expense struct {
	ID              string
	Invoice         string
	Subject         string
	Category        string
	Comment         string
	AmountRub       codatypes.MoneyRub
	AmountEur       codatypes.MoneyEur
	Status          string
	ActuallySpent   codatypes.MoneyRub
	RejectionReason string
	PendingSpend    codatypes.MoneyRub
	Balance         codatypes.MoneyRub
	LastCashOutDate *time.Time
}

func (i *Expense) HasPendingSpends() bool {
	return i.PendingSpend != 0
}

func (e Expense) String() string {
	return fmt.Sprintf("%s", e.ID)
}

type expenseGroup struct {
	Name     string
	Items    []*Expense
	TotalEur codatypes.MoneyEur
	TotalRub codatypes.MoneyRub
}

type ExpenseGroupMap map[string]*expenseGroup

func newExpenseGroup(name string) *expenseGroup {
	return &expenseGroup{Name: name, Items: make([]*Expense, 0)}
}

func (e *expenseGroup) Add(expense *Expense) {
	e.Items = append(e.Items, expense)
	e.TotalEur += expense.AmountEur
	e.TotalRub += expense.AmountRub
}

func GroupExpensesByCategory(expenses []*Expense) ExpenseGroupMap {
	expensesByCategory := make(ExpenseGroupMap)
	for _, expense := range expenses {
		if _, catExists := expensesByCategory[expense.Category]; !catExists {
			expensesByCategory[expense.Category] = newExpenseGroup(expense.Category)
		}
		expensesByCategory[expense.Category].Add(expense)
	}

	return expensesByCategory
}

func NewExpenseFromRow(row *coda.Row) *Expense {
	expense := Expense{}
	errs := codatypes.NewErrorContainer()
	var err error

	if expense.ID, err = codatypes.ToString(Ids.Expenses.Cols.ID, row); err != nil {
		errs.AddError(err)
	}

	if expense.Invoice, err = codatypes.ToString(Ids.Expenses.Cols.Invoice, row); err != nil {
		errs.AddError(err)
	}
	if expense.Subject, err = codatypes.ToString(Ids.Expenses.Cols.Subject, row); err != nil {
		errs.AddError(err)
	}
	if expense.Category, err = codatypes.ToString(Ids.Expenses.Cols.Category, row); err != nil {
		errs.AddError(err)
	}
	if expense.Comment, err = codatypes.ToString(Ids.Expenses.Cols.Comment, row); err != nil {
		errs.AddError(err)
	}
	if expense.AmountRub, err = codatypes.ToRub(Ids.Expenses.Cols.AmountRub, row); err != nil {
		errs.AddError(err)
	}
	if expense.AmountEur, err = codatypes.ToEur(Ids.Expenses.Cols.AmountEur, row); err != nil {
		errs.AddError(err)
	}
	if expense.Status, err = codatypes.ToString(Ids.Expenses.Cols.Status, row); err != nil {
		errs.AddError(err)
	}
	if expense.ActuallySpent, err = codatypes.ToRub(Ids.Expenses.Cols.ActuallySpent, row); err != nil {
		errs.AddError(err)
	}
	if expense.RejectionReason, err = codatypes.ToString(Ids.Expenses.Cols.RejectionReason, row); err != nil {
		errs.AddError(err)
	}
	if expense.PendingSpend, err = codatypes.ToRub(Ids.Expenses.Cols.PendingSpend, row); err != nil {
		errs.AddError(err)
	}
	if expense.Balance, err = codatypes.ToRub(Ids.Expenses.Cols.Balance, row); err != nil {
		errs.AddError(err)
	}
	if expense.LastCashOutDate, err = codatypes.ToDate(Ids.Expenses.Cols.LastCashOutDate, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &expense
}
