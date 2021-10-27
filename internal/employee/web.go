package employee

import (
	"archive/zip"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"oea-go/internal/common"
	"oea-go/internal/common/config"
	"oea-go/internal/employee/dto"
	"oea-go/internal/excel"
	"oea-go/internal/hellenic"
	"oea-go/resources"
	"time"
)

const limitForMonths = 10

type Page struct {
	SelectedMonth string
	Months        dto.Months
	Invoices      dto.Invoices
	Error         error
}

func NewErrorPage(err error) Page {
	return Page{Error: err}
}

func (p Page) IsMonthSelected(mon dto.Month) bool {
	return p.SelectedMonth == mon.ID
}

type Handler struct {
	config *config.Config
	client *Requests
	logger *zap.SugaredLogger
}

func NewHandler(cfg *config.Config, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		config: cfg,
		client: NewRequests(cfg.BaseUri, cfg.ApiTokenEm, cfg.DocIdEm),
		logger: logger,
	}
}

func (h Handler) getLastMonths(num int) dto.Months {
	months, _ := h.client.GetMonths() // TODO pass error
	curMonthIndex := months.IndexOfTime(time.Now())

	from := curMonthIndex - 1
	to := curMonthIndex + num - 1

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

	invoices, err := h.client.GetInvoices(month, With{Corrections: true, PrevInvoice: true, Employees: true})
	if err != nil {
		return NewErrorPage(err)
	}

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
		fmt.Fprint(resp, "no month provided")
		return
	}
	employee, containsEmployee := vars["employee"]
	if !containsEmployee {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(resp, "no employee provided")
		return
	}

	invoice, err := h.client.GetInvoiceForMonthAndEmployee(month, employee)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(resp, err)
		return
	}

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", invoice.Filename()))
	err = excel.RenderInvoice(resp, resources.MustReadBytes("assets/invoice_template_empl.xlsx"), invoice)
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
		fmt.Fprint(resp, "no month provided")
		return
	}
	invoices, err := h.client.GetInvoices(month, With{Employees: true, PrevInvoice: true, Corrections: true})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, err)
		return
	}

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

func (h Handler) DownloadHellenicPayroll(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(resp, "no month provided")
		return
	}
	invoices, err := h.client.GetInvoices(month, With{Employees: true, BankDetails: true})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, err)
		return
	}

	resp.Header().Add("Content-Type", "text/plain")
	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"payroll_%s.txt\"", month))

	renderErr := hellenic.CreatePayrollFile(resp, invoices, h.config.PayrollDebitAccount, time.Now())
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
		fmt.Fprint(resp, "no month provided")
		return
	}

	common.WriteMemProfile("before_getinvoices")

	invoices, err := h.client.GetInvoices(month, With{Employees: true})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, err)
		return
	}

	common.WriteMemProfile("after_getinvoices")

	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"invoices_%s.zip\"", month))

	zipWriter := zip.NewWriter(resp)
	defer zipWriter.Close()

	invoiceTpl := resources.MustReadBytes("assets/invoice_template_empl.xlsx")

	for _, invoice := range invoices {
		name := invoice.Filename()

		zipFileWriter, zipErr := zipWriter.Create(name)
		if zipErr != nil {
			h.logger.Warnf("skipping %s: %v", name, zipErr)
			return
		}

		renderErr := excel.RenderInvoice(zipFileWriter, invoiceTpl, invoice)

		if renderErr != nil {
			h.logger.Warnf("skipping %s: %v", name, renderErr)
			return
		}
	}

	common.WriteMemProfile("after_all_render")
}
