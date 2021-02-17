package office

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"oea-go/internal/common"
	"oea-go/internal/common/config"
	empl "oea-go/internal/employee"
	emplDto "oea-go/internal/employee/dto"
	"oea-go/internal/excel"
	"oea-go/internal/office/dto"
	"sort"
	"sync"
	"time"
)

func loadOfficeData(req *Requests, invoiceID string) OfficeTemplateData {
	var invoice *dto.Invoice
	var prevInvoice *dto.Invoice
	var expensesByCategory dto.ExpenseGroupMap
	var history *dto.History

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		log.Println("Loading invoice", invoiceID, "...")
		invoice = req.GetInvoice(invoiceID)
		if invoice.PrevInvoiceID != "" {
			log.Println("Loading prev invoice", invoice.PrevInvoiceID, "...")
			prevInvoice = req.GetInvoice(invoice.PrevInvoiceID)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Println("Loading expenses...")
		expenses := req.GetExpenses(invoiceID)
		expensesByCategory = dto.GroupExpensesByCategory(expenses)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Println("Loading history...")
		history = req.GetHistory()
		wg.Done()
	}()

	wg.Wait()

	return OfficeTemplateData{
		PrevInvoice:   *prevInvoice,
		Invoice:       *invoice,
		ExpenseGroups: expensesByCategory,
		History:       *history,
	}
}

func getMonthsN(req *empl.Requests, num int, now time.Time) emplDto.Months {
	months, _ := req.GetMonths() // TODO pass error

	curMonthIndex := months.IndexOfTime(now)

	if curMonthIndex == len(*months)-1 {
		num = 0
	}

	from := curMonthIndex
	to := curMonthIndex + num - 1

	if from < 0 {
		from = 0
	}
	if to > len(*months) {
		to = len(*months)
	}

	return (*months)[from:to]
}

func (h handler) loadTodayAndPastInvoices(req *empl.Requests, numPastInvoices int, today time.Time) dto.EmployeesHistoricReport {
	months := getMonthsN(req, numPastInvoices, today)
	numMonths := len(months)

	monthlyInvoices := make(chan emplDto.InvoicesPerMonth, numMonths)

	wg := sync.WaitGroup{}
	wg.Add(numMonths)

	for _, month := range months {
		go func(month *emplDto.Month) {
			defer wg.Done()
			h.logger.Infof("started loading invoices for month %v", month.ID)
			invoices, err := req.GetInvoices(month.ID, empl.With{Corrections: true, Employees: true})
			if err != nil {
				h.logger.Errorf("error during mass loading of invoices for month %v: %v", month.ID, err)
				return
			}
			sort.Sort(invoices)
			monthlyInvoices <- emplDto.InvoicesPerMonth{
				Invoices: invoices,
				Month:    month,
			}
			h.logger.Infof("finished loading invoices for month %v", month.ID)
		}(month)
	}

	wg.Wait()
	close(monthlyInvoices)

	hist := dto.EmployeesHistoricReport{}

	for invoices := range monthlyInvoices {
		rep := dto.NewEmployeesMonthlyReport(invoices.Month.ID)
		rep.AddItemsFromInvoices(invoices.Invoices)

		if invoices.Month.IsCurrent {
			hist.CurrentMonth = rep
		} else {
			hist.AppendHistoricReport(rep)
		}
	}

	hist.RunPostCalculations()

	return hist
}

func buildApprovalRequestHtml(officeData OfficeTemplateData, employeesData dto.EmployeesHistoricReport) string {
	tpl := template.Must(template.New("post").Parse(string(common.MustAsset("resources/post.go.html"))))

	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, TemplateData{
		Timestamp: time.Now(),
		Office:    officeData,
		Employees: employeesData,
	})

	if err != nil {
		panic(err)
	}

	return buf.String()
}

type Page struct {
	InvoiceID string
	Body      template.HTML
}

type handler struct {
	config *config.Config
	logger *zap.SugaredLogger
}

func NewHandler(cfg *config.Config, logger *zap.SugaredLogger) *handler {
	return &handler{config: cfg, logger: logger}
}

func (h handler) Home(vars map[string]string, req *http.Request) interface{} {
	client := NewRequests(h.config.BaseUri, h.config.ApiTokenOf, h.config.DocIdOf)
	invoiceId, _ := client.GetLastInvoice()

	return Page{InvoiceID: invoiceId}
}

func (h handler) ShowInvoiceData(vars map[string]string, req *http.Request) interface{} {
	officeClient := NewRequests(h.config.BaseUri, h.config.ApiTokenOf, h.config.DocIdOf)
	officeData := loadOfficeData(officeClient, vars["invoice"])

	emplClient := empl.NewRequests(h.config.BaseUri, h.config.ApiTokenEm, h.config.DocIdEm)
	employeesData := h.loadTodayAndPastInvoices(emplClient, 5, time.Now())

	html := buildApprovalRequestHtml(officeData, employeesData)

	return Page{
		InvoiceID: vars["invoice"],
		Body:      template.HTML(html),
	}
}

func (h handler) DownloadInvoice(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	officeClient := NewRequests(h.config.BaseUri, h.config.ApiTokenOf, h.config.DocIdOf)
	data := loadOfficeData(officeClient, vars["invoice"])

	resp.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.xlsx\"", data.Invoice.Filename()))

	templateSource, readErr := ioutil.ReadFile(h.config.FilesDir + "/invoice_template.xlsx")
	if readErr != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, readErr)
		return
	}

	excelErr := excel.RenderInvoice(resp, templateSource, &data.Invoice)

	if excelErr != nil {
		resp.Header().Del("Content-Type")
		resp.Header().Del("Content-Disposition")
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, excelErr)
	}
}
