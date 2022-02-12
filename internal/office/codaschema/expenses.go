package codaschema

import (
	"fmt"
	"html/template"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
)

func (e Expenses) CommentHTML() template.HTML {
	return template.HTML(common.MarkdownToHTML(e.Comment))
}

func (e Expenses) HasPendingSpends() bool {
	return e.PendingSpendMoney() != 0
}

func (e Expenses) String() string {
	return fmt.Sprintf("%s", e.ID)
}

func (e Expenses) AmountRUBMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(e.AmountRUB * 100)
}

func (e Expenses) AmountEURMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(e.AmountEUR * 100)
}

func (e Expenses) ActuallySpentMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(e.ActuallySpent * 100)
}

func (e Expenses) PendingSpendMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(e.PendingSpend * 100)
}

func (e Expenses) BalanceMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(e.Balance * 100)
}
