package invoices

import (
	"fmt"
	"oea-go/internal/codatypes"
	"oea-go/internal/employee/codaschema"
	"strings"
	"time"
)

func NewReportingInvoice(inv *codaschema.Invoice) reportingInvoice {
	return reportingInvoice{*inv}
}

type reportingInvoice struct {
	inv codaschema.Invoice
}

func (r reportingInvoice) Filename() string {
	nameWithoutSpaces := strings.Replace(r.inv.Employee.FirstRefName(), " ", "_", -1)
	nameLower := strings.ToLower(nameWithoutSpaces)

	return fmt.Sprintf("%s_%s.xlsx", r.inv.Month.FirstRefName(), nameLower)
}

func (r reportingInvoice) BeneficiaryRequisites() string {
	if res := r.inv.RecipientDetails.First(); res != nil {
		return res.BankRequisites
	}
	return "n/a"
}

func (r reportingInvoice) PayerRequisites() string {
	if res := r.inv.SenderDetails.First(); res != nil {
		return strings.ReplaceAll(res.Requisites, "\n", "\r\n")
	}
	return "n/a"
}

func (r reportingInvoice) BeneficiaryName() string {
	if res := r.inv.Employee.First(); res != nil {
		return res.EnglishFullName
	}
	return "n/a"
}

func (r reportingInvoice) PayerName() string {
	return ""
}

func (r reportingInvoice) Number() string {
	return r.inv.InvoiceHash
}

func (r reportingInvoice) DateYm() string {
	if res := r.inv.Month.First(); res != nil {
		return res.Month.Format("January 2006")
	}
	return "n/a"
}

func (r reportingInvoice) HourRate() codatypes.MoneyEur {
	return r.inv.HourRateMoney()
}

func (r reportingInvoice) Hours() uint16 {
	return uint16(r.inv.Hours)
}

func (r reportingInvoice) TotalEur() codatypes.MoneyEur {
	return r.inv.EURTotalMoney()
}

func (r reportingInvoice) FullMonthName() string {
	if res := r.inv.Month.First(); res != nil {
		return res.Month.Format("January")
	}
	return "n/a"
}

func (r reportingInvoice) DateFull() string {
	if res := r.inv.Month.First(); res != nil {
		return fmt.Sprintf("%02d %s", time.Now().Day(), res.Month.Format("Jan 2006"))
	}
	return "n/a"
}

func (r reportingInvoice) ContractNumber() string {
	return r.inv.ContractNumber
}

func (r reportingInvoice) ContractDate() string {
	return r.inv.ContractDate
}
