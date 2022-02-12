package web

import (
	"oea-go/internal/office/codaschema"
	"oea-go/internal/office/invoices"
	"time"
)

type officeData struct {
	PrevInvoice   *codaschema.Invoices
	Invoice       codaschema.Invoices
	ExpenseGroups invoices.ExpenseGroupMap
	History       invoices.History
}

type page struct {
	Error           error
	Timestamp       time.Time
	SidebarInvoices []codaschema.Invoices
	SelectedInvoice string
	Office          officeData
}

func (p page) Now() page {
	p.Timestamp = time.Now()

	return p
}
