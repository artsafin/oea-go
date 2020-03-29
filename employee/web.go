package employee

import (
	"net/http"
	"oea-go/common"
	"oea-go/employee/dto"
	"time"
)

const limitForMonths = 10

type Page struct {
	SelectedMonth string
	Months        dto.Months
	Invoices      dto.Invoices
}

func (p Page) IsMonthSelected(mon dto.Month) bool {
	return p.SelectedMonth == mon.ID
}

type Handler struct {
	config *common.Config
	client *Requests
}

func NewHandler(cfg *common.Config) *Handler {
	return &Handler{
		config: cfg,
		client: NewRequests(cfg.BaseUri, cfg.ApiTokenEm, cfg.DocIdEm),
	}
}

func (h Handler) getLastMonths(num int) dto.Months {
	months := h.client.GetMonths()

	var curMonthIndex int
	var err error

	if curMonthIndex, err = months.IndexOfYearMonth(time.Now()); err != nil {
		curMonthIndex = 0
	}

	return (*months)[curMonthIndex-1 : curMonthIndex+num-1]
}

func (h Handler) Home(vars map[string]string, req *http.Request) interface{} {
	return Page{
		Months: h.getLastMonths(limitForMonths),
	}
}

func (h Handler) Month(vars map[string]string, req *http.Request) interface{} {
	month, containsMonth := vars["month"]

	if !containsMonth {
		return h.Home(vars, req)
	}

	invoices := h.client.GetInvoices(month, With{Corrections: true, PrevInvoice: true})

	return Page{
		SelectedMonth: month,
		Months:        h.getLastMonths(limitForMonths),
		Invoices:      *invoices,
	}
}
