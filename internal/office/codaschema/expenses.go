package codaschema

import (
	"fmt"
	"html/template"
	"oea-go/internal/codatypes"
	"regexp"
)

// replaceAllStringSubmatchFunc taken from https://gist.github.com/elliotchance/d419395aa776d632d897
func replaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			if v[i] == -1 || v[i+1] == -1 {
				groups = append(groups, "")
			} else {
				groups = append(groups, str[v[i]:v[i+1]])
			}
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}

func (e Expenses) CommentHTML() template.HTML {
	linkRepl := regexp.MustCompile(`(?m)\[([^]]+)\]\(([^)]+)\)`)

	return template.HTML(replaceAllStringSubmatchFunc(linkRepl, e.Comment, func(g []string) string {
		return fmt.Sprintf(`<a href="%s">%s</a>`, g[1], g[2])
	}))
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
