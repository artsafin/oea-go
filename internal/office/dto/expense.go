package dto

import (
	"fmt"
	"oea-go/internal/common"
	"time"

	"github.com/artsafin/go-coda"
)

type Expense struct {
	ID              string
	Invoice         string
	Subject         string
	Category        string
	Comment         string
	AmountRub       common.MoneyRub
	AmountEur       common.MoneyEur
	Status          string
	ActuallySpent   common.MoneyRub
	RejectionReason string
	PendingSpend    common.MoneyRub
	Balance         common.MoneyRub
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
	TotalEur common.MoneyEur
	TotalRub common.MoneyRub
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
	errs := make([]common.UnexpectedFieldTypeErr, 0)
	var err *common.UnexpectedFieldTypeErr

	if expense.ID, err = common.ToString(Ids.Expenses.Cols.ID, row); err != nil {
		errs = append(errs, *err)
	}

	if expense.Invoice, err = common.ToString(Ids.Expenses.Cols.Invoice, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.Subject, err = common.ToString(Ids.Expenses.Cols.Subject, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.Category, err = common.ToString(Ids.Expenses.Cols.Category, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.Comment, err = common.ToString(Ids.Expenses.Cols.Comment, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.AmountRub, err = common.ToRub(Ids.Expenses.Cols.AmountRub, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.AmountEur, err = common.ToEur(Ids.Expenses.Cols.AmountEur, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.Status, err = common.ToString(Ids.Expenses.Cols.Status, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.ActuallySpent, err = common.ToRub(Ids.Expenses.Cols.ActuallySpent, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.RejectionReason, err = common.ToString(Ids.Expenses.Cols.RejectionReason, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.PendingSpend, err = common.ToRub(Ids.Expenses.Cols.PendingSpend, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.Balance, err = common.ToRub(Ids.Expenses.Cols.Balance, row); err != nil {
		errs = append(errs, *err)
	}
	if expense.LastCashOutDate, err = common.ToDate(Ids.Expenses.Cols.LastCashOutDate, row); err != nil {
		errs = append(errs, *err)
	}

	if len(errs) > 0 {
		stringErr := ""
		for _, err := range errs {
			stringErr += err.Error() + "; "
		}

		panic(stringErr)
	}

	return &expense
}
