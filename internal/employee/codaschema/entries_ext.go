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

func (e Entries) EURTotal() codatypes.MoneyEur {
	return codatypes.MoneyEur(100 * (e.EURAmount + e.RUBAmountInEUR))
}
