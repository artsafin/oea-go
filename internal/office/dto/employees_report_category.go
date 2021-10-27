package dto

import (
	"oea-go/internal/codatypes"
)

type EmployeesReportCategory struct {
	Name     string
	Total    codatypes.MoneyRub
	Payments []*EmployeeReportLine
}

func (cat *EmployeesReportCategory) AddPayment(payment *EmployeeReportLine) {
	cat.Payments = append(cat.Payments, payment)
	cat.Total += payment.Amount
}
