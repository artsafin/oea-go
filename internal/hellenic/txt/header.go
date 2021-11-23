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
	return fmt.Sprintf("%s|%5d|%s|%s|%s",
		h.uploadType,
		h.totalNumberOfTransactions,
		lPadSp(h.amount.Humanize("#.###,##"), 18),
		h.submissionDate,
		lPadSp(h.debitAccount, 16),
	)
}
