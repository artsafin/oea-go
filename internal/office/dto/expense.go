package dto

import (
	"fmt"
	"oea-go/internal/codatypes"
	"oea-go/internal/office/codaschema"
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

	if expense.ID, err = codatypes.ToString(codaschema.ID.Table.Expenses.Cols.ID.ID, row); err != nil {
		errs.AddError(err)
	}

	if expense.Invoice, err = codatypes.ToString(codaschema.ID.Table.Expenses.Cols.Invoice.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.Subject, err = codatypes.ToString(codaschema.ID.Table.Expenses.Cols.Subject.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.Category, err = codatypes.ToString(codaschema.ID.Table.Expenses.Cols.Category.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.Comment, err = codatypes.ToString(codaschema.ID.Table.Expenses.Cols.Comment.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.AmountRub, err = codatypes.ToRub(codaschema.ID.Table.Expenses.Cols.AmountRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.AmountEur, err = codatypes.ToEur(codaschema.ID.Table.Expenses.Cols.AmountEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.Status, err = codatypes.ToString(codaschema.ID.Table.Expenses.Cols.Status.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.ActuallySpent, err = codatypes.ToRub(codaschema.ID.Table.Expenses.Cols.ActuallySpent.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.RejectionReason, err = codatypes.ToString(codaschema.ID.Table.Expenses.Cols.RejectionReason.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.PendingSpend, err = codatypes.ToRub(codaschema.ID.Table.Expenses.Cols.PendingSpend.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.Balance, err = codatypes.ToRub(codaschema.ID.Table.Expenses.Cols.Balance.ID, row); err != nil {
		errs.AddError(err)
	}
	if expense.LastCashOutDate, err = codatypes.ToDate(codaschema.ID.Table.Expenses.Cols.LastCashOutDate.ID, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &expense
}
