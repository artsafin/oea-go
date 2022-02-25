package codaschema

import (
	"html/template"
)

func (e AllEmployees) GeneralSDURL() template.URL {
	return template.URL(e.GeneralSD)
}

func (e AllEmployees) FinanceSDURL() template.URL {
	return template.URL(e.FinanceSD)
}
