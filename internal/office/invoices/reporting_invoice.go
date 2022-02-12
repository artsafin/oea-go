package invoices

import (
	"fmt"
	"oea-go/internal/codatypes"
	"oea-go/internal/office/codaschema"
)

func NewReportingInvoice(inv *codaschema.Invoices) reportingInvoice {
	return reportingInvoice{*inv}
}

type reportingInvoice struct {
	inv codaschema.Invoices
}

func (r reportingInvoice) Filename() string {
	return r.inv.Filename
}

func (r reportingInvoice) BeneficiaryRequisites() string {
	return ""
}

func (r reportingInvoice) PayerRequisites() string {
	return ""
}

func (r reportingInvoice) BeneficiaryName() string {
	return ""
}

func (r reportingInvoice) PayerName() string {
	return ""
}

func (r reportingInvoice) Number() string {
	return fmt.Sprint(r.inv.Number)
}

func (r reportingInvoice) DateYm() string {
	return r.inv.Date.Format("January 2006")
}

func (r reportingInvoice) HourRate() codatypes.MoneyEur {
	return r.inv.HourRateMoney()
}

func (r reportingInvoice) Hours() uint16 {
	return uint16(r.inv.Hours)
}

func (r reportingInvoice) TotalEur() codatypes.MoneyEur {
	return r.inv.TotalEURMoney()
}

func (r reportingInvoice) DateFull() string {
	return r.inv.Date.Format("02.01.2006")
}

func (r reportingInvoice) ContractNumber() string {
	return "1"
}

func (r reportingInvoice) ContractDate() string {
	return "..."
}
