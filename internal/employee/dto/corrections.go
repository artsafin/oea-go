package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"html/template"
	"oea-go/internal/codatypes"
	"oea-go/internal/employee/codaschema"
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

	if corr.PaymentInvoice, err = codatypes.ToString(codaschema.ID.Table.Corrections.Cols.PaymentInvoice.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.Comment, err = codatypes.ToString(codaschema.ID.Table.Corrections.Cols.Comment.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.TotalCorrectionRub, err = codatypes.ToRub(codaschema.ID.Table.Corrections.Cols.TotalCorrectionRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.Category, err = codatypes.ToString(codaschema.ID.Table.Corrections.Cols.Category.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.AbsoluteCorrectionRub, err = codatypes.ToRub(codaschema.ID.Table.Corrections.Cols.AbsoluteCorrectionRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.AbsoluteCorrectionEur, err = codatypes.ToEur(codaschema.ID.Table.Corrections.Cols.AbsoluteCorrectionEUR.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.AbsCorrectionEurInRub, err = codatypes.ToRub(codaschema.ID.Table.Corrections.Cols.AbsCorrectionEURInRUB.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDayType, err = codatypes.ToString(codaschema.ID.Table.Corrections.Cols.PerDayType.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.NumberOfDays, err = codatypes.ToFloat64(codaschema.ID.Table.Corrections.Cols.NumberOfDays.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.CostOfDay, err = codatypes.ToRub(codaschema.ID.Table.Corrections.Cols.CostOfDay.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDay, err = codatypes.ToRub(codaschema.ID.Table.Corrections.Cols.PerDay.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDayCoefficient, err = codatypes.ToFloat64(codaschema.ID.Table.Corrections.Cols.PerDayCoefficient.ID, row); err != nil {
		errs.AddError(err)
	}
	if corr.PerDayCalculationInvoice, err = codatypes.ToString(codaschema.ID.Table.Corrections.Cols.PerDayCalculationInvoice.ID, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &corr
}
