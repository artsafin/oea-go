package office

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/internal/common"
	"oea-go/internal/office/codaschema"
	"oea-go/internal/office/dto"
	"sort"
)

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

func (requests *Requests) GetInvoice(invoiceID string) (*dto.Invoice, error) {
	params := coda.ListRowsParameters{
		Query: fmt.Sprintf("\"%s\":\"%s\"", codaschema.ID.Table.Invoices.Cols.No.ID, invoiceID),
	}
	resp, err := requests.Client.ListTableRows(requests.DocId, codaschema.ID.Table.Invoices.ID, params)
	if err != nil {
		return nil, err
	}

	if len(resp.Rows) == 0 {
		return nil, fmt.Errorf("invoice %s not found", invoiceID)
	}

	firstRow := resp.Rows[0]

	return dto.NewInvoiceFromRow(&firstRow), nil
}

func (requests *Requests) GetInvoices(query ...common.QueryParam) (dto.Invoices, error) {
	params := coda.ListRowsParameters{
		SortBy: "natural",
	}
	for _, q := range query {
		q.Apply(&params)
	}
	resp, err := requests.Client.ListTableRows(requests.DocId, codaschema.ID.Table.Invoices.ID, params)
	if err != nil {
		return nil, err
	}

	invoices := make(dto.Invoices, len(resp.Rows))
	for i, row := range resp.Rows {
		invoices[i] = dto.NewInvoiceFromRow(&row)
	}

	sort.Sort(invoices)

	return invoices, nil
}

func (requests *Requests) GetExpenses(invoiceID string) ([]*dto.Expense, error) {
	params := coda.ListRowsParameters{}
	resp, err := requests.Client.ListTableRows(requests.DocId, codaschema.ID.Table.Expenses.ID, params)
	if err != nil {
		return nil, err
	}

	expenses := make([]*dto.Expense, 0)

	for _, row := range resp.Rows {
		if row.Values[codaschema.ID.Table.Expenses.Cols.Invoice.ID] == invoiceID {
			expenses = append(expenses, dto.NewExpenseFromRow(&row))
		}
	}

	return expenses, nil
}

func (requests *Requests) GetHistory() (*dto.History, error) {
	sentInvoices := make([]*dto.Invoice, 0)
	grandTotal := dto.GrandTotal{}
	invoices, err := requests.GetInvoices()
	if err != nil {
		return nil, err
	}
	for _, inv := range invoices {
		if inv.Status != "" {
			sentInvoices = append(sentInvoices, inv)
			grandTotal.AddInvoice(inv)
		}
	}

	lastInvoice := sentInvoices[len(sentInvoices)-1]
	lastInvoice.Expenses, err = requests.GetExpenses(lastInvoice.No)
	if err != nil {
		return nil, err
	}

	return &dto.History{
		FirstInvoice: sentInvoices[0],
		Invoices:     sentInvoices[:len(sentInvoices)-1],
		LastInvoice:  lastInvoice,
		GrandTotal:   grandTotal,
	}, nil
}
