package dto

import "oea-go/common"

type Corrections struct {
	PaymentInvoice     string
	Comment            string
	TotalCorrectionRub common.MoneyRub
	Category           string

	AbsoluteCorrectionRub    common.MoneyRub
	AbsoluteCorrectionEur    common.MoneyEur
	AbsCorrectionEurInRub    common.MoneyRub
	PerDayType               string
	NumberOfDays             uint16
	CostOfDay                common.MoneyRub
	PerDay                   common.MoneyRub
	PerDayCoefficient        float64
	PerDayCalculationInvoice string
}
