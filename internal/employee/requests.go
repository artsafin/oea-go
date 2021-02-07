package employee

import (
	"errors"
	"fmt"
	"github.com/artsafin/go-coda"
	"log"
	"oea-go/internal/common"
	"oea-go/internal/employee/dto"
	"sort"
	"strings"
)

type With struct {
	Corrections bool
	PrevInvoice bool
	Employees   bool
}

type Requests struct {
	Client *common.CodaClient
	DocId  string
}

func NewRequests(baseUri, apiTokenOf, docId string) *Requests {
	return &Requests{
		Client: common.NewCodaClient(baseUri, apiTokenOf),
		DocId:  docId,
	}
}

func (requests *Requests) GetMonths() (*dto.Months, error) {
	params := coda.ListRowsParameters{}
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Months.Id, params)
	if err != nil {
		log.Printf("ListTableRows error: %v\n", err)
		return nil, err
	}

	months := make(dto.Months, len(resp.Rows))

	for k, v := range resp.Rows {
		months[k] = dto.NewMonthFromRow(&v)
	}

	sort.Sort(sort.Reverse(months))

	return &months, nil
}

func (requests *Requests) GetCurrentMonth() (string, error) {
	currMonth, err := requests.Client.GetFormula(requests.DocId, dto.Ids.CodaFormulas.CurrentMonth)

	if err != nil {
		return "", err
	}
	if currMonth.Formula.Value == "" {
		return "", errors.New("current month not found")
	}

	str, ok := currMonth.Formula.Value.(string)

	if !ok {
		return "", errors.New("cannot convert current month to string")
	}

	return str, nil
}

func uniqueInvoiceById(invoices dto.Invoices) map[string]*dto.Invoice {
	result := make(map[string]*dto.Invoice)

	for _, v := range invoices {
		result[v.Id] = v
	}

	return result
}

func (requests *Requests) GetInvoiceForMonthAndEmployee(month, employee string) *dto.Invoice {
	invoices := requests.GetInvoices(month, With{Employees: true})

	for _, invoice := range invoices {
		if invoice.EmployeeName == employee {
			return invoice
		}
	}

	return nil
}

func (requests *Requests) GetInvoices(month string, with With) dto.Invoices {
	params := coda.ListRowsParameters{
		Query: fmt.Sprintf("%s:\"%s\"", dto.Ids.Invoices.Cols.Month, month),
	}
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Invoices.Id, params)

	if err != nil {
		panic(err)
	}

	invoices := make(dto.Invoices, len(resp.Rows))

	var corrs map[string][]*dto.Correction
	var employees map[string]*dto.Employee

	if len(invoices) > 0 && with.Corrections {
		corrs = requests.getCorrectionsIndexedByInvoice(month)
	}

	if len(invoices) > 0 && with.Employees {
		employees = requests.GetAllEmployees()
	}

	thisMonth, prevMonth, monthErr := requests.getMonthsData(month)
	if monthErr != nil {
		log.Println(err)
	}

	var prevInvoices map[string]*dto.Invoice

	if len(invoices) > 0 && with.PrevInvoice {
		if monthErr == nil {
			prevInvoices = uniqueInvoiceById(requests.GetInvoices(prevMonth.ID, With{}))
		}
	}

	for i, row := range resp.Rows {
		invoices[i] = dto.NewInvoiceFromRow(&row)
		if with.Corrections {
			invoices[i].Corrections = corrs[invoices[i].Id]
		}
		if with.PrevInvoice {
			invoices[i].PrevInvoice = prevInvoices[invoices[i].PreviousInvoiceId]
		}
		if with.Employees {
			invoices[i].Employee = employees[invoices[i].EmployeeName]
		}

		invoices[i].MonthData = thisMonth
	}

	sort.Sort(invoices)

	return invoices
}

func (requests *Requests) getMonthsData(month string) (*dto.Month, *dto.Month, error) {
	var err error
	months, err := requests.GetMonths()
	if err != nil {
		return nil, nil, err
	}
	var thisMonth *dto.Month
	if thisMonth, err = months.FindByName(month); err != nil {
		return nil, nil, err
	}

	prevMonthName := thisMonth.PreviousMonthLink
	var prevMonth *dto.Month

	if prevMonth, err = months.FindByName(prevMonthName); err != nil {
		return nil, nil, err
	}

	return thisMonth, prevMonth, nil
}

func (requests *Requests) getCorrectionsIndexedByInvoice(month string) map[string][]*dto.Correction {
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Corrections.Id, coda.ListRowsParameters{})

	if err != nil {
		panic(err)
	}

	result := make(map[string][]*dto.Correction)

	for _, row := range resp.Rows {
		corr := dto.NewCorrectionFromRow(&row)
		if strings.Contains(corr.PaymentInvoice, month) {
			result[corr.PaymentInvoice] = append(result[corr.PaymentInvoice], corr)
		}
	}

	return result
}

func (requests *Requests) GetAllEmployees() map[string]*dto.Employee {
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Employees.Id, coda.ListRowsParameters{})

	if err != nil {
		panic(err)
	}

	result := make(map[string]*dto.Employee)

	for _, row := range resp.Rows {
		empl := dto.NewEmployeeFromRow(&row)
		result[empl.Name] = empl
	}

	return result
}
