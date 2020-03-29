package dto

import (
	"oea-go/common"
	"time"
)

type PatentCalculation struct {
	Invoice       string
	OpeningPatent *time.Time
	PeriodEnd     *time.Time
	FullMonths    uint16
	AnnualCost    common.MoneyRub
	PeriodCost    common.MoneyRub
	Period        string
}
