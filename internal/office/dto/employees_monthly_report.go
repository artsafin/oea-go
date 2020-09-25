package dto

import (
	"errors"
	"fmt"
	"html/template"
	"oea-go/internal/common"
	emplDto "oea-go/internal/employee/dto"
	"strings"
)

const (
	CatSalaries          = "Salaries"
	CatTaxes             = "Taxes"
	CatPatents           = "Patents"
	CatPaymentService    = "Payment services"
	CatDayoff            = "Day offs"
	CatGeneralCorrection = "Correction"
)

func NewEmployeesMonthlyReport(month string) *EmployeesMonthlyReport {
	return &EmployeesMonthlyReport{
		Month:      month,
		Categories: make(map[string]*EmployeesReportCategory),
	}
}

type EmployeesMonthlyReport struct {
	Month      string
	Categories map[string]*EmployeesReportCategory
	Total      common.MoneyRub
}

func (cats *EmployeesMonthlyReport) AddItemsFromInvoices(invoices emplDto.Invoices) {
	for _, inv := range invoices {
		cats.addItemsFromInvoice(inv)
	}
}

func (cats *EmployeesMonthlyReport) addItemsFromInvoice(invoice *emplDto.Invoice) {
	location := "n/a"
	if invoice.Employee != nil {
		location = invoice.Employee.Location
	}
	payment := EmployeeReportLine{
		Location: location,
		Name:     invoice.EmployeeName,
	}

	cats.addItem(CatSalaries, payment.WithAmount(invoice.BaseSalary))
	cats.addItem(CatPaymentService, payment.WithAmount(invoice.BankFees))

	for _, corr := range invoice.Corrections {
		comment := template.HTML(fmt.Sprintf("%s [%s]", strings.Replace(corr.Comment, "\n", "<br>", -1), corr.SubCategory()))
		cats.addItem(corr.Category, payment.WithAmountAndComment(corr.TotalCorrectionRub, comment))
	}

	if invoice.PatentRub > 0 {
		cats.addItem(CatPatents, payment.WithAmount(invoice.PatentRub))
	}

	if invoice.TaxesRub > 0 {
		cats.addItem(CatTaxes, payment.WithAmount(invoice.TaxesRub))
	}

	if invoice.UnpaidDay > 0 {
		cats.addItem(CatDayoff, payment.WithAmount(invoice.UnpaidDay))
	}
}

func (cats *EmployeesMonthlyReport) HasCategory(cat string) bool {
	_, found := cats.Categories[cat]

	return found
}

func (cats *EmployeesMonthlyReport) GetCategoryByName(cat string) (*EmployeesReportCategory, error) {
	if cat, found := cats.Categories[cat]; found {
		return cat, nil
	}

	return nil, errors.New("category not found: " + cat)
}

func (cats *EmployeesMonthlyReport) addItem(cat string, payment *EmployeeReportLine) {
	if _, hasCategory := cats.Categories[cat]; !hasCategory {
		cats.Categories[cat] = &EmployeesReportCategory{
			Name:     cat,
			Total:    common.MoneyRub(0),
			Payments: make([]*EmployeeReportLine, 0),
		}
	}

	cats.Categories[cat].AddPayment(payment)
	cats.Total += payment.Amount
}
