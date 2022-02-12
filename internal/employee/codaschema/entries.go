package codaschema

import (
	"fmt"
	"html/template"
	"oea-go/internal/codatypes"
	"strings"
)

func (e Entries) LongComment() template.HTML {
	htmlComment := strings.Replace(e.Comment, "\n", "<br>", -1)
	return template.HTML(fmt.Sprintf("<code>%s</code><br>%s", e.Type.FirstRefName(), htmlComment))
}

func (e Entries) EURTotal() float64 {
	return e.EURAmount + e.RUBAmountInEUR
}

func (e Entries) EURAmountMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(e.EURAmount * 100)
}
func (e Entries) RUBAmountMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(e.RUBAmount * 100)
}
func (e Entries) RUBAmountInEURMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(e.RUBAmountInEUR * 100)
}