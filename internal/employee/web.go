package employee

import (
	"archive/zip"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"oea-go/internal/common"
	"oea-go/internal/employee/dto"
	"oea-go/internal/excel"
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
	months, _ := h.client.GetMonths() // TODO pass error
	curMonthIndex := months.IndexOfTime(time.Now())

	from := curMonthIndex-1
	to := curMonthIndex+num-1

	if from < 0 {
		from = 0
	}
	if to > len(*months) {
		to = len(*months)
	}

	return (*months)[from:to]
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

	invoices := h.client.GetInvoices(month, With{Corrections: true, PrevInvoice: true, Employees: true})

	return Page{
		SelectedMonth: month,
		Months:        h.getLastMonths(limitForMonths),
		Invoices:      invoices,
	}
}

func (h Handler) DownloadInvoice(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	employee, containsEmployee := vars["employee"]
	if !containsEmployee {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	invoice := h.client.GetInvoiceForMonthAndEmployee(month, employee)

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", invoice.Filename()))
	err := excel.RenderInvoice(resp, common.MustAsset("resources/invoice_template_empl.xlsx"), invoice)
	if err != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, err)
	}
}

func (h Handler) DownloadPayrollReport(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	invoices := h.client.GetInvoices(month, With{Employees: true, PrevInvoice: true, Corrections: true})

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"payroll_report_%s.xlsx\"", month))

	renderErr := excel.RenderPayrollReport(resp, invoices)
	if renderErr != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, renderErr)
	}
}

func (h Handler) DownloadAllInvoices(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	common.WriteMemProfile("before_getinvoices")

	invoices := h.client.GetInvoices(month, With{Employees: true})

	common.WriteMemProfile("after_getinvoices")

	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"invoices_%s.zip\"", month))

	zipWriter := zip.NewWriter(resp)
	defer zipWriter.Close()

	invoiceTpl := common.MustAsset("resources/invoice_template_empl.xlsx")

	for _, invoice := range invoices {
		name := invoice.Filename()

		zipFileWriter, zipErr := zipWriter.Create(name)
		if zipErr != nil {
			log.Printf("skipping %s: %v", name, zipErr)
			return
		}

		renderErr := excel.RenderInvoice(zipFileWriter, invoiceTpl, invoice)

		if renderErr != nil {
			log.Printf("skipping %s: %v", name, renderErr)
			return
		}
	}

	common.WriteMemProfile("after_all_render")
}
