package dto

import "oea-go/internal/common"

type EmployeesReportCategory struct {
	Name     string
	Total    common.MoneyRub
	Payments []*EmployeeReportLine
}

func (cat *EmployeesReportCategory) AddPayment(payment *EmployeeReportLine) {
	cat.Payments = append(cat.Payments, payment)
	cat.Total += payment.Amount
}
