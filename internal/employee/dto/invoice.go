package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
	"oea-go/internal/employee/codaschema"
	"strings"
	"time"
)

const paymentDateDaysBeforeEndOfMonth = 2

type Invoice struct {
	ID        string
	InvoiceNo string

	MonthName           string
	Month               *Month
	Employee            *Employee
	EmployeeName        string
	EmployeeNameSlug    string
	PaymentChecksPassed bool
	Entries             Entries
	EUREntriesSum       codatypes.MoneyEur
	RUBEntriesSum       codatypes.MoneyRub
	EURRUBExpected      codatypes.MoneyRub
	EURRUBActual        codatypes.MoneyRub
	EURSubtotal         codatypes.MoneyEur
	EURRounding         codatypes.MoneyEur
	EURTotal            codatypes.MoneyEur
	hours               uint16
	hourRate            codatypes.MoneyEur
	SenderEntityName    string
	SenderDetails       *LegalEntity
	RecipientDetailsID  string
	RecipientDetails    *BankDetails
	PrevInvoiceID       string
	PrevInvoice         *Invoice
}

//region InvoiceDataProvider implementation

func (inv *Invoice) Filename() string {
	nameWithoutSpaces := strings.Replace(inv.EmployeeName, " ", "_", -1)
	nameLower := strings.ToLower(nameWithoutSpaces)

	return fmt.Sprintf("%s_%s.xlsx", inv.MonthName, nameLower)
}

func (inv *Invoice) BeneficiaryRequisites() string {
	if inv.RecipientDetails == nil {
		return "n/a"
	}
	return inv.RecipientDetails.BankRequisites
}

func (inv *Invoice) PayerRequisites() string {
	if inv.Employee == nil {
		return "n/a"
	}
	return strings.ReplaceAll(inv.Employee.LegalEntity.Requisites, "\n", "\r\n")
}

func (inv *Invoice) BeneficiaryName() string {
	if inv.Employee == nil {
		return "n/a"
	}
	return inv.Employee.EnglishFullName
}

func (inv *Invoice) PayerName() string {
	return ""
}

func (inv *Invoice) Number() string {
	return inv.InvoiceNo
}

func (inv *Invoice) DateYm() string {
	if inv.Month == nil {
		return "n/a"
	}
	return inv.Month.LastMonthDay.Format("January 2006")
}

func (inv *Invoice) HourRate() codatypes.MoneyEur {
	return inv.hourRate
}

func (inv *Invoice) Hours() uint16 {
	return inv.hours
}

func (inv *Invoice) TotalEur() codatypes.MoneyEur {
	return inv.EURTotal
}

func (inv *Invoice) FullMonthName() string {
	if inv.Month == nil {
		return "n/a"
	}
	return inv.Month.LastMonthDay.Format("January")
}

func (inv *Invoice) DateFull() string {
	if inv.Month == nil {
		return "n/a"
	}
	return fmt.Sprintf("%02d %s", time.Now().Day(), inv.Month.LastMonthDay.Format("Jan 2006"))
}

func (inv *Invoice) ContractNumber() string {
	return inv.Employee.ContractNumber
}

func (inv *Invoice) ContractDate() string {
	return inv.Employee.ContractDate
}

//endregion

func (inv *Invoice) DatePayment() string {
	if inv.Month == nil {
		return "n/a"
	}
	paymentDate := common.AddWorkingDate(*inv.Month.LastMonthDay, 0, 0, -paymentDateDaysBeforeEndOfMonth)
	now := time.Now()
	if paymentDate.Before(now) {
		paymentDate = now
	}
	return paymentDate.Format("2 Jan 2006")
}

func (inv *Invoice) RUBEntriesInEUR() codatypes.MoneyEur {
	return inv.EURSubtotal - inv.EUREntriesSum
}

func NewInvoiceFromRow(row *coda.Row) (*Invoice, error) {
	invoice := Invoice{}
	errs := codatypes.NewErrorContainer()
	var err error

	if invoice.ID, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.ID.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.InvoiceNo, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.InvoiceHash.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.MonthName, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.Month.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EmployeeName, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.Employee.ID, row); err != nil {
		errs.AddError(err)
	} else {
		invoice.EmployeeNameSlug = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(invoice.EmployeeName)), " ", "-")
	}
	if invoice.PaymentChecksPassed, err = codatypes.ToBool(codaschema.ID.Table.Invoice.Cols.PaymentChecksPassed.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EUREntriesSum, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.EUREntries.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.RUBEntriesSum, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.RUBEntries.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EURRUBExpected, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.EURRUBExpected.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EURRUBActual, err = codatypes.ToRub(codaschema.ID.Table.Invoice.Cols.EURRUBActual.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EURSubtotal, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.EURSubtotal.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EURRounding, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.EURRounding.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.EURTotal, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.EURTotal.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.hours, err = codatypes.ToUint16(codaschema.ID.Table.Invoice.Cols.Hours.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.hourRate, err = codatypes.ToEur(codaschema.ID.Table.Invoice.Cols.HourRate.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.SenderEntityName, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.SenderDetails.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.RecipientDetailsID, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.RecipientDetails.ID, row); err != nil {
		errs.AddError(err)
	}
	if invoice.PrevInvoiceID, err = codatypes.ToString(codaschema.ID.Table.Invoice.Cols.PreviousInvoice.ID, row); err != nil {
		errs.AddError(err)
	}

	return &invoice, common.JoinErrors(errs)
}
