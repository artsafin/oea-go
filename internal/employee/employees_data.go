package employee

import (
	"github.com/artsafin/goda"
	"github.com/pkg/errors"
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

type Data interface {
	GetMonths() (*dto.Months, error)
	GetCurrentMonth() (string, error)
	GetInvoiceForMonthAndEmployee(month, employee string) (*dto.Invoice, error)
	GetInvoices(month string, with With) (dto.Invoices, error)
	GetAllEmployees() (map[string]*dto.Employee, error)
}

type codaData struct {
	doc *common.CodaDocument
}

func NewCodaData(baseUri, apiTokenOf, docId string) (Data, error) {
	cocl, err := common.NewCodaDocument(baseUri, apiTokenOf, docId)
	if err != nil {
		return nil, err
	}
	return &codaData{cocl}, nil
}

func (data *codaData) GetMonths() (*dto.Months, error) {
	resp, err := data.doc.ListRows(dto.Ids.Months.Id, &goda.ListRowsParams{})
	if err != nil {
		return nil, err
	}

	months := make(dto.Months, len(resp.Items))

	for k, v := range resp.Items {
		months[k] = dto.NewMonthFromRow(&v)
	}

	sort.Sort(sort.Reverse(months))

	return &months, nil
}

func (data *codaData) GetCurrentMonth() (string, error) {
	currMonth, err := data.doc.GetFormula(dto.Ids.CodaFormulas.CurrentMonth)

	if err != nil {
		return "", err
	}
	str, ok := currMonth.Value.(string)

	if !ok || str == "" {
		return "", errors.Errorf("current month not found")
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

func (data *codaData) GetInvoiceForMonthAndEmployee(month, employee string) (*dto.Invoice, error) {
	invoices, err := data.GetInvoices(month, With{Employees: true})
	if err != nil {
		return nil, err
	}

	for _, invoice := range invoices {
		if invoice.EmployeeName == employee {
			return invoice, nil
		}
	}

	return nil, errors.Errorf("not found: %v %v", month, employee)
}

func (data *codaData) GetInvoices(month string, with With) (dto.Invoices, error) {
	params := goda.ListRowsParams{}.WithQuery(dto.Ids.Invoices.Cols.Month, month)
	resp, err := data.doc.ListRows(dto.Ids.Invoices.Id, &params)

	if err != nil {
		return nil, err
	}

	invoices := make(dto.Invoices, len(resp.Items))

	var corrs map[string][]*dto.Correction
	var employees map[string]*dto.Employee

	if len(invoices) > 0 && with.Corrections {
		corrs, err = data.getCorrectionsIndexedByInvoice(month)
		if err != nil {
			return nil, errors.Wrap(err, "unable to find corrections")
		}
	}

	if len(invoices) > 0 && with.Employees {
		employees, err = data.GetAllEmployees()
		if err != nil {
			return nil, errors.Wrap(err, "unable to find employees")
		}
	}

	thisMonth, prevMonth, monthErr := data.getMonthsData(month)
	if monthErr != nil {
		log.Println(err)
	}

	var prevInvoicesById map[string]*dto.Invoice

	if len(invoices) > 0 && with.PrevInvoice {
		if monthErr == nil {
			prevInvoices, prevErr := data.GetInvoices(prevMonth.ID, With{})
			if prevErr != nil {
				return nil, errors.Wrap(prevErr, "unable to find previous invoices")
			}
			prevInvoicesById = uniqueInvoiceById(prevInvoices)
		}
	}

	for i, row := range resp.Items {
		invoices[i] = dto.NewInvoiceFromRow(&row)
		if with.Corrections {
			invoices[i].Corrections = corrs[invoices[i].Id]
		}
		if with.PrevInvoice {
			invoices[i].PrevInvoice = prevInvoicesById[invoices[i].PreviousInvoiceId]
		}
		if with.Employees {
			invoices[i].Employee = employees[invoices[i].EmployeeName]
		}

		invoices[i].MonthData = thisMonth
	}

	return invoices, nil
}

func (data *codaData) getMonthsData(month string) (*dto.Month, *dto.Month, error) {
	var err error
	months, err := data.GetMonths()
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

func (data *codaData) getCorrectionsIndexedByInvoice(month string) (map[string][]*dto.Correction, error) {
	resp, err := data.doc.ListRows(dto.Ids.Corrections.Id, &goda.ListRowsParams{})

	if err != nil {
		return nil, err
	}

	result := make(map[string][]*dto.Correction)

	for _, row := range resp.Items {
		corr := dto.NewCorrectionFromRow(&row)
		if strings.Contains(corr.PaymentInvoice, month) {
			result[corr.PaymentInvoice] = append(result[corr.PaymentInvoice], corr)
		}
	}

	return result, nil
}

func (data *codaData) GetAllEmployees() (map[string]*dto.Employee, error) {
	resp, err := data.doc.ListRows(dto.Ids.Employees.Id, &goda.ListRowsParams{})

	if err != nil {
		return nil, err
	}

	result := make(map[string]*dto.Employee)

	for _, row := range resp.Items {
		empl := dto.NewEmployeeFromRow(&row)
		result[empl.Name] = empl
	}

	return result, nil
}
