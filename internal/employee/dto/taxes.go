package dto

import (
	"oea-go/internal/common"
	"time"
)

type TaxCalculation struct {
	Invoice            string
	OpeningDateIp      *time.Time
	PeriodStart        *time.Time
	PeriodEnd          *time.Time
	AmountIpDays       uint16
	MedicalFund        common.MoneyRub
	PensionFundFixed   common.MoneyRub
	PensionFundPercent common.MoneyRub
	Amount             common.MoneyRub
}
