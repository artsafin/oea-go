package dto

import (
	"html/template"
	"oea-go/internal/codatypes"
)

type EmployeeReportLine struct {
	Location string
	Name     string
	Comment  template.HTML
	Amount   codatypes.MoneyRub
}

func (p EmployeeReportLine) clone() EmployeeReportLine {
	return EmployeeReportLine{
		Location: p.Location,
		Name:     p.Name,
		Comment:  p.Comment,
		Amount:   p.Amount,
	}
}

func (p EmployeeReportLine) WithAmount(amount codatypes.MoneyRub) *EmployeeReportLine {
	newp := p.clone()
	newp.Amount = amount
	return &newp
}

func (p EmployeeReportLine) WithAmountAndComment(amount codatypes.MoneyRub, comment template.HTML) *EmployeeReportLine {
	newp := p.clone()
	newp.Amount = amount
	newp.Comment = comment
	return &newp
}
