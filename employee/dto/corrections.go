package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/common"
	"strings"
)

type Correction struct {
	PaymentInvoice           string
	Comment                  string
	TotalCorrectionRub       common.MoneyRub
	Category                 string
	AbsoluteCorrectionRub    common.MoneyRub
	AbsoluteCorrectionEur    common.MoneyEur
	AbsCorrectionEurInRub    common.MoneyRub
	PerDayType               string
	NumberOfDays             float64
	CostOfDay                common.MoneyRub
	PerDay                   common.MoneyRub
	PerDayCoefficient        float64
	PerDayCalculationInvoice string
}

func (corr *Correction) LongComment() string {
	return fmt.Sprintf("[%s] %s: %s", corr.SubCategory(), corr.Category, corr.Comment)
}

func (corr *Correction) SubCategory() string {
	subcats := make([]string, 0)

	if corr.PerDayType != "" {
		subcats = append(subcats, fmt.Sprintf("%.2f days @%.2f", corr.NumberOfDays, corr.PerDayCoefficient))
	}

	if corr.AbsoluteCorrectionRub != 0 {
		subcats = append(subcats, fmt.Sprintf("Abs RUB"))
	}

	if corr.AbsoluteCorrectionEur != 0 {
		subcats = append(subcats, fmt.Sprintf("Abs EUR"))
	}

	return strings.Join(subcats, ", ")
}

func NewCorrectionFromRow(row *coda.Row) *Correction {
	corr := Correction{}
	errs := make([]common.UnexpectedFieldTypeErr, 0)
	var err *common.UnexpectedFieldTypeErr

	if corr.PaymentInvoice, err = common.ToString(Ids.Corrections.Cols.PaymentInvoice, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.Comment, err = common.ToString(Ids.Corrections.Cols.Comment, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.TotalCorrectionRub, err = common.ToRub(Ids.Corrections.Cols.TotalCorrectionRub, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.Category, err = common.ToString(Ids.Corrections.Cols.Category, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.AbsoluteCorrectionRub, err = common.ToRub(Ids.Corrections.Cols.AbsoluteCorrectionRub, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.AbsoluteCorrectionEur, err = common.ToEur(Ids.Corrections.Cols.AbsoluteCorrectionEur, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.AbsCorrectionEurInRub, err = common.ToRub(Ids.Corrections.Cols.AbsCorrectionEurInRub, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.PerDayType, err = common.ToString(Ids.Corrections.Cols.PerDayType, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.NumberOfDays, err = common.ToFloat64(Ids.Corrections.Cols.NumberOfDays, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.CostOfDay, err = common.ToRub(Ids.Corrections.Cols.CostOfDay, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.PerDay, err = common.ToRub(Ids.Corrections.Cols.PerDay, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.PerDayCoefficient, err = common.ToFloat64(Ids.Corrections.Cols.PerDayCoefficient, row); err != nil {
		errs = append(errs, *err)
	}
	if corr.PerDayCalculationInvoice, err = common.ToString(Ids.Corrections.Cols.PerDayCalculationInvoice, row); err != nil {
		errs = append(errs, *err)
	}

	if len(errs) > 0 {
		stringErr := ""
		for _, err := range errs {
			stringErr += err.Error() + "; "
		}

		panic(stringErr)
	}

	return &corr
}
