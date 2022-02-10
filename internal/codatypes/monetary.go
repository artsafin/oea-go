package codatypes

import (
	"github.com/dustin/go-humanize"
)

type MoneyRub int64
type MoneyEur int64

func (v MoneyRub) String() string {
	return EnPrinter.Sprintf("%.2f ₽", float64(v)/100)
}

func (v MoneyRub) Number() interface{} {
	return float64(v) / 100
}

func (v MoneyRub) Neg() MoneyRub {
	return -v
}

func (v MoneyEur) String() string {
	return EnPrinter.Sprintf("€%.2f", float64(v)/100)
}

func (v MoneyEur) Humanize(format string) string {
	return humanize.FormatFloat(format, float64(v)/100)
}

func (v MoneyEur) Number() float64 {
	return float64(v) / 100
}

func (v MoneyEur) CurrencyISO4217() string {
	return "EUR"
}

func (v MoneyEur) Neg() MoneyEur {
	return -v
}

func (v MoneyEur) ToRub(rate MoneyRub) MoneyRub {
	return MoneyRub(float64(v)/100) * rate
}
