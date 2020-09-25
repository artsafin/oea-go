package dto

import (
	"html/template"
	"oea-go/internal/common"
)

type EmployeeReportLine struct {
	Location string
	Name     string
	Comment  template.HTML
	Amount   common.MoneyRub
}

func (p EmployeeReportLine) clone() EmployeeReportLine {
	return EmployeeReportLine{
		Location: p.Location,
		Name:     p.Name,
		Comment:  p.Comment,
		Amount:   p.Amount,
	}
}

func (p EmployeeReportLine) WithAmount(amount common.MoneyRub) *EmployeeReportLine {
	newp := p.clone()
	newp.Amount = amount
	return &newp
}

func (p EmployeeReportLine) WithAmountAndComment(amount common.MoneyRub, comment template.HTML) *EmployeeReportLine {
	newp := p.clone()
	newp.Amount = amount
	newp.Comment = comment
	return &newp
}
