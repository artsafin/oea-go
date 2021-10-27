package dto

import (
	"oea-go/internal/codatypes"
	"time"
)

type PatentCalculation struct {
	Invoice       string
	OpeningPatent *time.Time
	PeriodEnd     *time.Time
	FullMonths    uint16
	AnnualCost    codatypes.MoneyRub
	PeriodCost    codatypes.MoneyRub
	Period        string
}
