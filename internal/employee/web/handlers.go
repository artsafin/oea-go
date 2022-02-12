package web

import (
	"archive/zip"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"oea-go/internal/common"
	"oea-go/internal/employee/codaschema"
	"oea-go/internal/employee/invoices"
	"oea-go/internal/employee/months"
	"oea-go/internal/employee/payroll"
	"oea-go/internal/excel"
	"oea-go/internal/hellenic"
	"oea-go/resources"
	"sync"
	"time"
)

type handlers struct {
	doc    *codaschema.CodaDocument
	logger *zap.SugaredLogger
}

func NewHandlers(doc *codaschema.CodaDocument, logger *zap.SugaredLogger) *handlers {
	return &handlers{
		doc:    doc,
		logger: logger,
	}
}

func (h handlers) writeErr(resp http.ResponseWriter, err interface{}, status ...int) {
	common.WriteHTTPErrAndLog(h.logger, resp, err, status...)
}

func (h handlers) Home(vars map[string]string, req *http.Request) interface{} {
	ms, err := months.GetNearest(h.doc)
	return page{
		Months: ms,
		Error:  err,
	}
}

func (h handlers) Month(vars map[string]string, req *http.Request) interface{} {
	month, containsMonth := vars["month"]

	if !containsMonth {
		return h.Home(vars, req)
	}

	var err error
	pg := page{SelectedMonth: month}
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		pg.Invoices, err = invoices.FindByMonthID(h.doc, month, codaschema.Tables{
			Months:       true,
			Entries:      true,
			Invoice:      true,
			AllEmployees: true,
			BankDetails:  true,
			LegalEntity:  true,
		})
	}()
	go func() {
		defer wg.Done()
		pg.Months, err = months.GetNearest(h.doc)
	}()
	wg.Wait()

	if err != nil {
		return newErrorPage(err)
	}

	return pg
}

func (h handlers) DownloadInvoice(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		h.writeErr(resp, "no month provided")
		return
	}
	employee, containsEmployee := vars["employee"]
	if !containsEmployee {
		h.writeErr(resp, "no employee provided")
		return
	}

	invoice, err := invoices.FindByEmployeeAndMonthID(h.doc, month, employee)
	if err != nil {
		h.writeErr(resp, err)
		return
	}

	reportInv := invoices.NewReportingInvoice(&invoice)

	tpl, err := resources.Open("assets/invoice_template_empl.xlsx")
	if err != nil {
		h.writeErr(resp, err)
		return
	}
	defer tpl.Close()

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, reportInv.Filename()))
	err = excel.RenderInvoice(resp, tpl, reportInv)
	if err != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		h.writeErr(resp, err, http.StatusInternalServerError)
	}
}

func (h handlers) DownloadPayrollReport(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		h.writeErr(resp, "no month provided")
		return
	}
	invs, err := invoices.FindByMonthID(h.doc, month, codaschema.Tables{AllEmployees: true, Entries: true})
	if err != nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="payroll_report_%s.xlsx"`, month))

	err = excel.RenderPayrollReport(resp, invs)
	if err != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		h.writeErr(resp, err, http.StatusInternalServerError)
	}
}

func (h handlers) DownloadHellenicPayroll(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		h.writeErr(resp, "no month provided")
		return
	}

	payableInvoices, err := invoices.FindPayableByMonthID(h.doc, month, codaschema.Tables{AllEmployees: true, BankDetails: true, LegalEntity: true})

	if err != nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}

	executionDate, err := payroll.GetScheduleByMonth(h.doc, month)
	if err != nil || executionDate == nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}

	resp.Header().Add("Content-Type", "text/plain")
	resp.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="payroll_%s.txt"`, executionDate.Format("020106")))

	err = hellenic.CreatePayrollFile(resp, payableInvoices, time.Now(), *executionDate)
	if err != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		h.writeErr(resp, err, http.StatusInternalServerError)
	}
}

func (h handlers) DownloadAllInvoices(resp http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	month, containsMonth := vars["month"]
	if !containsMonth {
		h.writeErr(resp, "no month provided")
		return
	}

	common.WriteMemProfile("before_getinvoices")

	invs, err := invoices.FindByMonthID(h.doc, month, codaschema.Tables{AllEmployees: true, BankDetails: true, LegalEntity: true})
	if err != nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}

	common.WriteMemProfile("after_getinvoices")

	resp.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="invoices_%s.zip"`, month))

	zipWriter := zip.NewWriter(resp)
	defer zipWriter.Close()

	invoiceTpl, err := resources.Open("assets/invoice_template_empl.xlsx")
	if err != nil {
		h.writeErr(resp, err)
		return
	}
	defer invoiceTpl.Close()

	for _, invoice := range invs {
		repInv := invoices.NewReportingInvoice(&invoice)
		name := repInv.Filename()

		zipFileWriter, zipErr := zipWriter.Create(name)
		if zipErr != nil {
			h.logger.Warnf("skipping %s: %v", name, zipErr)
			return
		}

		renderErr := excel.RenderInvoice(zipFileWriter, invoiceTpl, repInv)

		if renderErr != nil {
			h.logger.Warnf("skipping %s: %v", name, renderErr)
			return
		}
	}

	common.WriteMemProfile("after_all_render")
}