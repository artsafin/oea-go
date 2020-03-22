package employee

import (
	"html/template"
	"net/http"
	"oea-go/common"
	"oea-go/employee/dto"
	"time"
)

const limitForMonths = 10

type Page struct {
	Months dto.Months
	Body   template.HTML
}

type Handler struct {
	Config *common.Config
}

func (h Handler) Home(vars map[string]string, req *http.Request) interface{} {
	client := NewRequests(h.Config.BaseUri, h.Config.ApiTokenEm, h.Config.DocIdEm)
	months := client.GetLastMonths()

	var curMonthIndex int
	var err error

	if curMonthIndex, err = months.IndexOf(time.Now()); err != nil {
		curMonthIndex = 0
	}

	lastNMonths := (*months)[curMonthIndex-1 : curMonthIndex+limitForMonths-1]

	return Page{
		Months: lastNMonths,
	}
}
