package invoices

import (
	"context"
	"fmt"
	"oea-go/internal/employee/codaschema"
	"sort"
	"strings"
)

func sortInvoices(invs []codaschema.Invoice) {
	sort.Slice(invs, func(i, j int) bool {
		monthCmp := strings.Compare(invs[i].Month.String(), invs[j].Month.String())
		if monthCmp != 0 {
			return monthCmp < 0
		}

		return strings.Compare(invs[i].Employee.String(), invs[j].Employee.String()) < 0
	})
}

func FindByMonthID(doc *codaschema.CodaDocument, monthID string, with codaschema.Tables) (list []codaschema.Invoice, err error) {
	ctx := context.Background()

	invs, order, err := doc.MapOfInvoice(ctx)
	if err != nil {
		return nil, err
	}

	doc.AddInvoicesToCache(invs)

	with.Months = true
	err = doc.LoadRelationsInvoice(ctx, invs, with)
	if err != nil {
		return nil, err
	}

	for _, rowid := range order {
		if m := invs[rowid].Month.First(); m != nil && m.ID == monthID {
			list = append(list, *invs[rowid])
		}
	}

	sortInvoices(list)

	return
}

func FindByEmployeeAndMonthID(doc *codaschema.CodaDocument, employeeName, monthID string) (inv codaschema.Invoice, err error) {
	ctx := context.Background()

	invs, _, err := doc.MapOfInvoice(ctx)
	if err != nil {
		return codaschema.Invoice{}, err
	}

	err = doc.LoadRelationsInvoice(ctx, invs, codaschema.Tables{
		Months:       true,
		AllEmployees: true,
		BankDetails:  true,
		LegalEntity:  true,
	})
	if err != nil {
		return codaschema.Invoice{}, err
	}

	for _, invoice := range invs {
		m := invoice.Month.First()
		e := invoice.Employee.First()

		if m != nil && e != nil && m.ID == monthID && e.Name == employeeName {
			return *invoice, nil
		}
	}

	return codaschema.Invoice{}, fmt.Errorf("no invoice found for %v and %v", employeeName, monthID)
}

func FindPayableByMonthID(doc *codaschema.CodaDocument, monthID string, with codaschema.Tables) (list []codaschema.Invoice, err error) {
	invs, err := FindByMonthID(doc, monthID, with)
	if err != nil {
		return nil, err
	}

	for _, inv := range invs {
		if inv.PaymentChecksPassed {
			list = append(list, inv)
		}
	}

	sortInvoices(list)

	return
}
