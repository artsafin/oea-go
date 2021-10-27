package txt

import (
	"fmt"
	"oea-go/internal/codatypes"
)

type header struct {
	uploadType                string
	totalNumberOfTransactions int32
	amount                    codatypes.MoneyEur
	submissionDate            string
	debitAccount              string
}

func (h *header) increment(amount codatypes.MoneyEur) {
	h.totalNumberOfTransactions++
	h.amount += amount
}

func (h *header) String() string {
	return fmt.Sprintf("%s|%d|%s|%s|%s",
		h.uploadType,
		h.totalNumberOfTransactions,
		h.amount.Humanize("#.###,##"),
		underscore(h.submissionDate),
		underscore(h.debitAccount),
	)
}
