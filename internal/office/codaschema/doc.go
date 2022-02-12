package codaschema

import (
	"context"
)

func (doc *CodaDocument) AddInvoicesToCache(m map[RowID]*Invoices) {
	doc.relationsCache.Store("Invoices", m)
}

func (doc *CodaDocument) MapOfInvoicesWithCache(ctx context.Context) (m map[RowID]*Invoices, err error) {
	var _invoicesInter interface{}
	var ok bool

	if _invoicesInter, ok = doc.relationsCache.Load("Invoices"); !ok {
		m, _, err = doc.MapOfInvoices(ctx)

		if err != nil {
			doc.AddInvoicesToCache(m)
		}

		return
	}

	return _invoicesInter.(map[RowID]*Invoices), nil
}

func (doc *CodaDocument) AddExpensesToCache(m map[RowID]*Expenses) {
	doc.relationsCache.Store("Expenses", m)
}
