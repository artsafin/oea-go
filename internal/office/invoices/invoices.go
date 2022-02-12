package invoices

import (
	"context"
	"fmt"
	"github.com/artsafin/coda-go-client/codaapi"
	"oea-go/internal/office/codaschema"
	"sort"
)

func sortInvoices(list []codaschema.Invoices) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Date.Unix() < list[j].Date.Unix()
	})
}
func sortInvoicesReverse(list []codaschema.Invoices) {
	sort.Slice(list, func(i, j int) bool {
		return list[j].Date.Unix() < list[i].Date.Unix()
	})
}

func GetRecent(doc *codaschema.CodaDocument) (list []codaschema.Invoices, err error) {
	list, err = doc.ListInvoices(
		context.Background(),
		codaapi.ListRows.Query(codaschema.ID.Table.Invoices.Cols.IsRecent.ID, "true"),
		codaapi.ListRows.SortBy(codaapi.RowsSortByNatural),
	)
	if err != nil {
		return
	}

	sortInvoicesReverse(list)

	return
}

func FindByName(doc *codaschema.CodaDocument, name string, with codaschema.Tables) (codaschema.Invoices, error) {
	invs, _, err := doc.MapOfInvoices(
		context.Background(),
		codaapi.ListRows.Query(codaschema.ID.Table.Invoices.Cols.No.ID, name),
	)

	if err != nil {
		return codaschema.Invoices{}, err
	}

	err = doc.LoadRelationsInvoices(context.Background(), invs, with)
	if err != nil {
		return codaschema.Invoices{}, err
	}

	for _, inv := range invs {
		return *inv, nil
	}

	return codaschema.Invoices{}, fmt.Errorf("could not find invoice %v", name)
}

func GetHistory(doc *codaschema.CodaDocument) (History, error) {
	sentInvoices := make([]codaschema.Invoices, 0)
	total := grandTotal{}

	invoices, err := doc.ListInvoices(context.Background())
	if err != nil {
		return History{}, err
	}

	for _, inv := range invoices {
		if inv.Status != "" {
			sentInvoices = append(sentInvoices, inv)
			total.AddInvoice(inv)
		}
	}

	sortInvoices(invoices)

	if len(sentInvoices) == 0 {
		return History{}, nil
	}

	lastInvoice := sentInvoices[len(sentInvoices)-1]

	expMap, _, err := doc.MapOfExpenses(context.Background())
	if err != nil {
		return History{}, err
	}
	lastInvoice.PlannedExpenses.Hydrate(expMap)

	return History{
		FirstInvoice: sentInvoices[0],
		PastInvoices: sentInvoices[:len(sentInvoices)-1],
		LastInvoice:  lastInvoice,
		GrandTotal:   total,
	}, nil
}
