package codaschema

import (
	"html/template"
	"oea-go/internal/common"
)

func (e AllEmployees) GeneralSDURL() template.URL {
	return template.URL(common.MarkdownLinkToURL(e.GeneralSD))
}

func (e AllEmployees) FinanceSDURL() template.URL {
	return template.URL(common.MarkdownLinkToURL(e.FinanceSD))
}
