package office

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"oea-go/common"
	empl "oea-go/employee"
	"oea-go/office/dto"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	CatSalaries          = "Salaries"
	CatTaxes             = "Taxes"
	CatPatents           = "Patents"
	CatPaymentService    = "Payment services"
	CatDayoff            = "Day offs"
	CatGeneralCorrection = "Correction"
)

func loadOfficeData(invoiceID, baseUri, apiToken, docId string) OfficeTemplateData {
	req := NewRequests(baseUri, apiToken, docId)

	var invoice *dto.Invoice
	var prevInvoice *dto.Invoice
	var expensesByCategory dto.ExpenseGroupMap
	var history *dto.History

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println("Loading invoice", invoiceID, "...")
		invoice = req.GetInvoice(invoiceID)
		if invoice.PrevInvoiceID != "" {
			fmt.Println("Loading prev invoice", invoice.PrevInvoiceID, "...")
			prevInvoice = req.GetInvoice(invoice.PrevInvoiceID)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Loading expenses...")
		expenses := req.GetExpenses(invoiceID)
		expensesByCategory = dto.GroupExpensesByCategory(expenses)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Loading history...")
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

func loadEmployeesData(baseUri, apiToken, docId string) *dto.EmployeesPaymentCategories {
	req := empl.NewRequests(baseUri, apiToken, docId)

	month, err := req.GetCurrentMonth()
	if err != nil {
		panic(err)
	}

	log.Println("Fetching month", month)

	invoices := req.GetInvoices(month, empl.With{Corrections: true})

	sort.Sort(invoices)

	emplCats := dto.NewEmployeesPaymentCategories(month)

	for _, invoice := range invoices {

		emplCats.AddItem(CatSalaries, &dto.EmployeePayment{
			Name:   invoice.EmployeeName,
			Amount: invoice.BaseSalary,
		})
		emplCats.AddItem(CatPaymentService, &dto.EmployeePayment{
			Name:   invoice.EmployeeName,
			Amount: invoice.BankFees,
		})

		for _, corr := range invoice.Corrections {
			emplCats.AddItem(corr.Category, &dto.EmployeePayment{
				Name:    invoice.EmployeeName,
				Comment: template.HTML(fmt.Sprintf("%s [%s]", strings.Replace(corr.Comment, "\n", "<br>", -1), corr.SubCategory())),
				Amount:  corr.TotalCorrectionRub,
			})
		}

		if invoice.PatentRub > 0 {
			emplCats.AddItem(CatPatents, &dto.EmployeePayment{
				Name:   invoice.EmployeeName,
				Amount: invoice.PatentRub,
			})
		}

		if invoice.TaxesRub > 0 {
			emplCats.AddItem(CatTaxes, &dto.EmployeePayment{
				Name:   invoice.EmployeeName,
				Amount: invoice.TaxesRub,
			})
		}

		if invoice.UnpaidDay > 0 {
			emplCats.AddItem(CatDayoff, &dto.EmployeePayment{
				Name:   invoice.EmployeeName,
				Amount: invoice.UnpaidDay,
			})
		}
	}

	return emplCats
}

func buildApprovalRequestHtml(officeData OfficeTemplateData, employeesData *dto.EmployeesPaymentCategories) string {
	tpl := template.Must(template.New("post").Parse(string(common.MustAsset("resources/post.go.html"))))

	var sb strings.Builder
	err := tpl.Execute(&sb, TemplateData{
		Timestamp: time.Now(),
		Office:    officeData,
		Employees: employeesData,
	})

	if err != nil {
		panic(err)
	}

	return sb.String()
}

type Page struct {
	InvoiceID string
	Body      template.HTML
}

type handler struct {
	config *common.Config
}

func NewHandler(cfg *common.Config) *handler {
	return &handler{config: cfg}
}

func (h handler) Home(vars map[string]string, req *http.Request) interface{} {
	client := NewRequests(h.config.BaseUri, h.config.ApiTokenOf, h.config.DocIdOf)
	invoiceId, _ := client.GetLastInvoice()

	return Page{InvoiceID: invoiceId}
}

func (h handler) ShowInvoiceData(vars map[string]string, req *http.Request) interface{} {
	officeData := loadOfficeData(vars["invoice"], h.config.BaseUri, h.config.ApiTokenOf, h.config.DocIdOf)
	employeesData := loadEmployeesData(h.config.BaseUri, h.config.ApiTokenEm, h.config.DocIdEm)

	html := buildApprovalRequestHtml(officeData, employeesData)
	return Page{
		InvoiceID: vars["invoice"],
		Body:      template.HTML(html),
	}
}

func (h handler) DownloadInvoice(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	data := loadOfficeData(vars["invoice"], h.config.BaseUri, h.config.ApiTokenOf, h.config.DocIdOf)

	resp.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.xlsx\"", data.Invoice.Filename()))

	common.RenderExcelTemplate(resp, common.MustAsset("resources/invoice_template.xlsx"), &data.Invoice)
}
