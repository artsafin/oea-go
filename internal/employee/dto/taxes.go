package dto

import (
	"oea-go/internal/codatypes"
	"time"
)

type TaxCalculation struct {
	Invoice            string
	OpeningDateIp      *time.Time
	PeriodStart        *time.Time
	PeriodEnd          *time.Time
	AmountIpDays       uint16
	MedicalFund        codatypes.MoneyRub
	PensionFundFixed   codatypes.MoneyRub
	PensionFundPercent codatypes.MoneyRub
	Amount             codatypes.MoneyRub
}
