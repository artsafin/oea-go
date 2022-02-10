package codaschema

import (
	"fmt"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
	"strings"
	"time"
)

//region InvoiceDataProvider implementation

func (inv *Invoice) Filename() string {
	nameWithoutSpaces := strings.Replace(inv.Employee.FirstRefName(), " ", "_", -1)
	nameLower := strings.ToLower(nameWithoutSpaces)

	return fmt.Sprintf("%s_%s.xlsx", inv.Month.FirstRefName(), nameLower)
}

func (inv *Invoice) BeneficiaryRequisites() string {
	if inv.RecipientDetails == nil {
		return "n/a"
	}
	return inv.RecipientDetails.BankRequisites
}

func (inv *Invoice) PayerRequisites() string {
	if inv.SenderDetails == nil {
		return "n/a"
	}
	return strings.ReplaceAll(inv.SenderDetails.Requisites, "\n", "\r\n")
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
	return inv.contractNumber
}

func (inv *Invoice) ContractDate() string {
	return inv.contractDate
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
