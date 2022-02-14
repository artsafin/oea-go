package web

import (
	"archive/zip"
	"bytes"
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

type CodaDocFactory interface {
	New() *codaschema.CodaDocument
}

type handlers struct {
	doc    CodaDocFactory
	logger *zap.SugaredLogger
}

func NewHandlers(doc CodaDocFactory, logger *zap.SugaredLogger) *handlers {
	return &handlers{
		doc:    doc,
		logger: logger,
	}
}

func (h handlers) writeErr(resp http.ResponseWriter, err interface{}, status ...int) {
	common.WriteHTTPErrAndLog(h.logger, resp, err, status...)
}

func (h handlers) Home(vars map[string]string, req *http.Request) interface{} {
	doc := h.doc.New()
	ms, err := months.GetNearest(doc)
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

	doc := h.doc.New()

	var errMonths, errInvoices error
	pg := page{SelectedMonth: month}
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		pg.Invoices, errInvoices = invoices.FindByMonthID(doc, month, codaschema.Tables{
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
		pg.Months, errMonths = months.GetNearest(doc)
	}()
	wg.Wait()

	if errInvoices != nil {
		return newErrorPage(errInvoices)
	}
	if errMonths != nil {
		return newErrorPage(errMonths)
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

	doc := h.doc.New()
	invoice, err := invoices.FindByEmployeeAndMonthID(doc, employee, month)
	if err != nil {
		h.writeErr(resp, err)
		return
	}

	reportInv := invoices.NewReportingInvoice(&invoice)

	tpl := resources.MustReadBytes("assets/invoice_template_empl.xlsx")

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, reportInv.Filename()))
	err = excel.RenderInvoice(resp, bytes.NewReader(tpl), reportInv)
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
	doc := h.doc.New()
	invs, err := invoices.FindByMonthID(doc, month, codaschema.Tables{AllEmployees: true, Entries: true})
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

	doc := h.doc.New()
	payableInvoices, err := invoices.FindPayableByMonthID(
		doc,
		month,
		codaschema.Tables{
			LoadRelationsRecursive: true,
			AllEmployees: true,
			BankDetails: true,
			LegalEntity: true,
			BeneficiaryBank: true,
		},
	)

	if err != nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}

	executionDate, err := payroll.GetScheduleByMonth(doc, month)
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

	doc := h.doc.New()
	invs, err := invoices.FindByMonthID(doc, month, codaschema.Tables{AllEmployees: true, BankDetails: true, LegalEntity: true})
	if err != nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}

	common.WriteMemProfile("after_getinvoices")

	resp.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="invoices_%s.zip"`, month))

	zipWriter := zip.NewWriter(resp)
	defer zipWriter.Close()

	invoiceTplBs := resources.MustReadBytes("assets/invoice_template_empl.xlsx")

	for _, invoice := range invs {
		repInv := invoices.NewReportingInvoice(&invoice)
		name := repInv.Filename()

		zipFileWriter, zipErr := zipWriter.Create(name)
		if zipErr != nil {
			h.logger.Warnf("zipWriter.Create: skipping %s: %v", name, zipErr)
			continue
		}

		renderErr := excel.RenderInvoice(zipFileWriter, bytes.NewReader(invoiceTplBs), repInv)

		if renderErr != nil {
			h.logger.Warnf("RenderInvoice: skipping %s: %v", name, renderErr)
			continue
		}
	}

	common.WriteMemProfile("after_all_render")
}
