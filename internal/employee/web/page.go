package web

import "oea-go/internal/employee/codaschema"

func newErrorPage(err error) page {
	return page{Error: err}
}

type page struct {
	SelectedMonth string
	Months        []codaschema.Months
	Invoices      []codaschema.Invoice
	Error         error
}

func (p page) IsMonthSelected(mon codaschema.Months) bool {
	return p.SelectedMonth == mon.ID
}
