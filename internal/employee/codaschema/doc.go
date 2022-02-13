package codaschema

func (doc *CodaDocument) AddInvoicesToCache(m map[RowID]*Invoice) {
	doc.relationsCache.Store("Invoice", m)
}
