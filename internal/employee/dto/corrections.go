package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"html/template"
	"oea-go/internal/codatypes"
	"strings"
)

type Correction struct {
	PaymentInvoice           string
	Comment                  string
	TotalCorrectionRub       codatypes.MoneyRub
	Category                 string
	AbsoluteCorrectionRub    codatypes.MoneyRub
	AbsoluteCorrectionEur    codatypes.MoneyEur
	AbsCorrectionEurInRub    codatypes.MoneyRub
	PerDayType               string
	NumberOfDays             float64
	CostOfDay                codatypes.MoneyRub
	PerDay                   codatypes.MoneyRub
	PerDayCoefficient        float64
	PerDayCalculationInvoice string
}

func (corr *Correction) LongComment() template.HTML {
	htmlComment := strings.Replace(corr.Comment, "\n", "<br>", -1)
	return template.HTML(fmt.Sprintf("<code>%s - %s</code><br>%s", corr.Category, corr.SubCategory(), htmlComment))
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
	errs := codatypes.NewErrorContainer()
	var err error

	if corr.PaymentInvoice, err = codatypes.ToString(Ids.Corrections.Cols.PaymentInvoice, row); err != nil {
		errs.AddError(err)
	}
	if corr.Comment, err = codatypes.ToString(Ids.Corrections.Cols.Comment, row); err != nil {
		errs.AddError(err)
	}
	if corr.TotalCorrectionRub, err = codatypes.ToRub(Ids.Corrections.Cols.TotalCorrectionRub, row); err != nil {
		errs.AddError(err)
	}
	if corr.Category, err = codatypes.ToString(Ids.Corrections.Cols.Category, row); err != nil {
		errs.AddError(err)
	}
	if corr.AbsoluteCorrectionRub, err = codatypes.ToRub(Ids.Corrections.Cols.AbsoluteCorrectionRub, row); err != nil {
		errs.AddError(err)
	}
	if corr.AbsoluteCorrectionEur, err = codatypes.ToEur(Ids.Corrections.Cols.AbsoluteCorrectionEur, row); err != nil {
		errs.AddError(err)
	}
	if corr.AbsCorrectionEurInRub, err = codatypes.ToRub(Ids.Corrections.Cols.AbsCorrectionEurInRub, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDayType, err = codatypes.ToString(Ids.Corrections.Cols.PerDayType, row); err != nil {
		errs.AddError(err)
	}
	if corr.NumberOfDays, err = codatypes.ToFloat64(Ids.Corrections.Cols.NumberOfDays, row); err != nil {
		errs.AddError(err)
	}
	if corr.CostOfDay, err = codatypes.ToRub(Ids.Corrections.Cols.CostOfDay, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDay, err = codatypes.ToRub(Ids.Corrections.Cols.PerDay, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDayCoefficient, err = codatypes.ToFloat64(Ids.Corrections.Cols.PerDayCoefficient, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDayCalculationInvoice, err = codatypes.ToString(Ids.Corrections.Cols.PerDayCalculationInvoice, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &corr
}
