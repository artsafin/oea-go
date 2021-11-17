package employee

import (
	"errors"
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
	"oea-go/internal/employee/codaschema"
	"oea-go/internal/employee/dto"
	"sort"
	"strings"
	"time"
)

type With struct {
	Corrections   bool
	PrevInvoice   bool
	Employees     bool
	BankDetails   bool
	LegalEntities bool
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
	resp, err := requests.Client.ListTableRows(requests.DocId, codaschema.ID.Table.Months.ID, params)
	if err != nil {
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
	currMonth, err := requests.Client.GetFormula(requests.DocId, codaschema.ID.Formula.CurrentMonth.ID)

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

func (requests *Requests) GetInvoiceForMonthAndEmployee(month, employee string) (*dto.Invoice, error) {
	invoices, err := requests.GetInvoices(month, With{Employees: true})
	if err != nil {
		return nil, err
	}

	for _, invoice := range invoices {
		if invoice.EmployeeName == employee {
			return invoice, nil
		}
	}

	return nil, nil
}

func (requests *Requests) GetInvoices(month string, with With) (invoices dto.Invoices, err error) {
	params := coda.ListRowsParameters{
		Query: fmt.Sprintf("%s:\"%s\"", codaschema.ID.Table.Invoice.Cols.Month.ID, month),
	}
	resp, err := requests.Client.ListTableRows(requests.DocId, codaschema.ID.Table.Invoice.ID, params)

	if err != nil {
		return nil, err
	}

	numRows := len(resp.Rows)
	invoices = make(dto.Invoices, 0, numRows)

	var corrs map[string][]*dto.Correction
	var employees map[string]*dto.Employee
	var bankDetails map[string]dto.BankDetails

	if numRows > 0 && with.Corrections {
		corrs, err = requests.getCorrectionsIndexedByInvoice(month)
		if err != nil {
			return nil, err
		}
	}

	if numRows > 0 && with.Employees {
		employees, err = requests.GetAllEmployees(with)
		if err != nil {
			return nil, err
		}
	}

	thisMonth, prevMonth, err := requests.getMonthsData(month)
	if err != nil {
		return nil, err
	}

	var prevInvoices map[string]*dto.Invoice

	if numRows > 0 && with.PrevInvoice {
		prevInvoicesList, err := requests.GetInvoices(prevMonth.ID, With{})
		if err != nil {
			return nil, err
		}
		prevInvoices = uniqueInvoiceById(prevInvoicesList)
	}

	if numRows > 0 && with.BankDetails {
		bankDetails, err = requests.GetAllBankDetails()
		if err != nil {
			return nil, err
		}
	}

	for _, row := range resp.Rows {
		invoice := dto.NewInvoiceFromRow(&row)
		if with.Corrections {
			invoice.Corrections = corrs[invoice.Id]
		}
		if with.PrevInvoice {
			invoice.PrevInvoice = prevInvoices[invoice.PreviousInvoiceId]
		}
		if with.Employees {
			invoice.Employee = employees[invoice.EmployeeName]
		}

		if invoiceBankDetails, detailsOk := bankDetails[invoice.BankDetailsID]; with.BankDetails && detailsOk {
			invoice.BankDetails = &invoiceBankDetails
		}

		invoice.MonthData = thisMonth

		invoices = append(invoices, invoice)
	}

	sort.Sort(invoices)

	return invoices, nil
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

func (requests *Requests) getCorrectionsIndexedByInvoice(month string) (map[string][]*dto.Correction, error) {
	resp, err := requests.Client.ListTableRows(requests.DocId, codaschema.ID.Table.Corrections.ID, coda.ListRowsParameters{})

	if err != nil {
		return nil, err
	}

	result := make(map[string][]*dto.Correction)

	for _, row := range resp.Rows {
		corr := dto.NewCorrectionFromRow(&row)
		if strings.Contains(corr.PaymentInvoice, month) {
			result[corr.PaymentInvoice] = append(result[corr.PaymentInvoice], corr)
		}
	}

	return result, nil
}

func (requests *Requests) GetAllEmployees(with With) (map[string]*dto.Employee, error) {
	resp, err := requests.Client.ListTableRows(requests.DocId, codaschema.ID.Table.AllEmployees.ID, coda.ListRowsParameters{})

	if err != nil {
		return nil, err
	}

	var legalEntities map[string]dto.LegalEntity
	if with.LegalEntities {
		legalEntities, err = requests.GetAllLegalEntities()
		if err != nil {
			return nil, err
		}
	}

	result := make(map[string]*dto.Employee)

	for _, row := range resp.Rows {
		empl := dto.NewEmployeeFromRow(&row)

		if legalEntity, found := legalEntities[empl.LegalEntityName]; found && with.LegalEntities {
			empl.LegalEntity = &legalEntity
		}

		result[empl.Name] = empl
	}

	return result, nil
}

func (requests *Requests) GetAllBankDetails() (map[string]dto.BankDetails, error) {
	bankDetailsResp, err := requests.Client.ListViewRows(requests.DocId, codaschema.ID.Table.BankDetails.ID, coda.ListViewRowsParameters{
		ValueFormat: "rich",
	})

	if err != nil {
		return nil, err
	}

	banks, err := requests.GetBeneficiaryBanksByRowID()

	if err != nil {
		return nil, err
	}

	result := make(map[string]dto.BankDetails)

	for _, row := range bankDetailsResp.Rows {
		d := dto.NewBankDetailsFromRow(&row)

		if bank, ok := banks[d.Bank.RowID]; ok {
			d.Bank = bank
		}
		result[d.ID] = d
	}

	return result, nil
}

func (requests *Requests) GetBeneficiaryBanksByRowID() (map[string]dto.BeneficiaryBank, error) {
	resp, err := requests.Client.ListViewRows(requests.DocId, codaschema.ID.Table.BeneficiaryBank.ID, coda.ListViewRowsParameters{
		ValueFormat: "rich",
	})

	if err != nil {
		return nil, err
	}
	result := make(map[string]dto.BeneficiaryBank)

	for _, row := range resp.Rows {
		d := dto.NewBeneficiaryBankFromRow(&row)
		result[row.Id] = d
	}

	return result, nil
}

func (requests *Requests) GetAllLegalEntities() (map[string]dto.LegalEntity, error) {
	resp, err := requests.Client.ListViewRows(requests.DocId, codaschema.ID.Table.LegalEntity.ID, coda.ListViewRowsParameters{})

	if err != nil {
		return nil, err
	}
	result := make(map[string]dto.LegalEntity)

	for _, row := range resp.Rows {
		d := dto.NewLegalEntityFromRow(&row)
		result[d.EntityName] = d
	}

	return result, nil
}

func (requests *Requests) GetPayrollScheduleByMonth(month string) (*time.Time, error) {
	params := coda.ListViewRowsParameters{
		Query: fmt.Sprintf("%s:\"%s\"", codaschema.ID.Table.PayrollSchedule.Cols.Month.ID, month),
	}
	resp, err := requests.Client.ListViewRows(requests.DocId, codaschema.ID.Table.PayrollSchedule.ID, params)

	if err != nil {
		return nil, err
	}

	if len(resp.Rows) != 1 {
		return nil, errors.New("schedule not found for month " + month)
	}

	res, err := codatypes.ToDate(codaschema.ID.Table.PayrollSchedule.Cols.ExecutionDate.ID, &resp.Rows[0])
	if err != nil {
		return nil, err
	}

	return res, nil
}
