package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"oea-go/internal/common"
	"oea-go/internal/excel"
	"oea-go/internal/office/codaschema"
	"oea-go/internal/office/invoices"
	"os"
	"sync"
	"time"
)

type handlers struct {
	filesDir string
	doc      *codaschema.CodaDocument
	logger   *zap.SugaredLogger
}

func NewHandlers(filesDir string, doc *codaschema.CodaDocument, logger *zap.SugaredLogger) *handlers {
	return &handlers{
		filesDir: filesDir,
		doc:      doc,
		logger:   logger,
	}
}

func (h handlers) writeErr(resp http.ResponseWriter, err interface{}, status ...int) {
	common.WriteHTTPErrAndLog(h.logger, resp, err, status...)
}

func (h handlers) Home(vars map[string]string, req *http.Request) interface{} {
	invs, err := invoices.GetRecent(h.doc)

	return page{SidebarInvoices: invs, Error: err}.Now()
}

func (h handlers) ShowInvoiceData(vars map[string]string, req *http.Request) interface{} {
	if _, ok := vars["invoice"]; !ok {
		return page{Error: fmt.Errorf("invoice not passed")}.Now()
	}

	invoiceID := vars["invoice"]

	var wg sync.WaitGroup
	var err error
	var invoice codaschema.Invoices
	var hist invoices.History

	wg.Add(2)

	go func(){
		defer wg.Done()

		invoice, err = invoices.FindByName(h.doc, invoiceID, codaschema.Tables{Invoices: true, Expenses: true})
	}()
	go func() {
		defer wg.Done()

		hist, err = invoices.GetHistory(h.doc)
	}()

	wg.Wait()

	if err != nil {
		return page{Error: err}.Now()
	}

	return page{
		Timestamp: time.Now(),
		Office: officeData{
			PrevInvoice:   invoice.PrevInvoice.First(),
			Invoice:       invoice,
			ExpenseGroups: invoices.GroupExpensesByCategory(invoice.PlannedExpenses.All()),
			History:       hist,
		},
		Error:           err,
		SelectedInvoice: invoiceID,
	}
}

func (h handlers) DownloadInvoice(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if _, ok := vars["invoice"]; !ok {
		h.writeErr(resp, "no invoice passed")
		return
	}

	invoiceID := vars["invoice"]

	invoice, err := invoices.FindByName(h.doc, invoiceID, codaschema.Tables{Invoices: true, Expenses: true})

	if err != nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}

	tplFile, err := os.Open(h.filesDir + "/invoice_template.xlsx")
	if err != nil {
		h.writeErr(resp, err, http.StatusInternalServerError)
		return
	}
	defer tplFile.Close()

	repInv := invoices.NewReportingInvoice(&invoice)

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.xlsx"`, repInv.Filename()))
	err = excel.RenderInvoice(resp, tplFile, repInv)

	if err != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		h.writeErr(resp, err, http.StatusInternalServerError)
	}
}
