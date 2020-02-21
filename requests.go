package main

import (
	"fmt"
	"sort"

	"github.com/artsafin/ofa-go/dto"
	"github.com/phouse512/go-coda"
)

func (ccl *CodaClient) getLastInvoice() string {
	lastInvoice, err := ccl.GetFormula(dto.Ids.OfficeAccounting.Id, dto.Ids.CodaFormulas.LastInvoice)

	if err != nil {
		panic(err)
	}
	if lastInvoice.Formula.Value == "" {
		panic("Last invoice is empty")
	}

	str, ok := lastInvoice.Formula.Value.(string)

	if !ok {
		panic("Last invoice is not a string")
	}

	return str
}

func (ccl *CodaClient) getInvoice(invoiceID string) *dto.Invoice {
	params := coda.ListRowsParameters{
		Query: fmt.Sprintf("\"%s\":\"%s\"", dto.Ids.Invoices.Cols.No, invoiceID),
	}
	resp, err := ccl.ListTableRows(dto.Ids.OfficeAccounting.Id, dto.Ids.Invoices.Id, params)
	if err != nil {
		panic(err)
	}

	if len(resp.Rows) == 0 {
		panic(fmt.Sprintf("Invoice %s is empty", invoiceID))
	}

	firstRow := resp.Rows[0]

	return dto.NewInvoiceFromRow(&firstRow)
}

func (ccl *CodaClient) getAllInvoices() dto.Invoices {
	params := coda.ListRowsParameters{
		SortBy: "natural",
	}
	resp, err := ccl.ListTableRows(dto.Ids.OfficeAccounting.Id, dto.Ids.Invoices.Id, params)
	if err != nil {
		panic(err)
	}

	invoices := make(dto.Invoices, len(resp.Rows))
	for i, row := range resp.Rows {
		invoices[i] = dto.NewInvoiceFromRow(&row)
	}

	sort.Sort(invoices)

	return invoices
}

func (ccl *CodaClient) getExpenses(invoiceID string) []*dto.Expense {
	params := coda.ListRowsParameters{}
	resp, err := ccl.ListTableRows(dto.Ids.OfficeAccounting.Id, dto.Ids.Expenses.Id, params)
	if err != nil {
		panic(err)
	}

	expenses := make([]*dto.Expense, 0)

	for _, row := range resp.Rows {
		if row.Values[dto.Ids.Expenses.Cols.Invoice] == invoiceID {
			expenses = append(expenses, dto.NewExpenseFromRow(&row))
		}
	}

	return expenses
}

func (ccl *CodaClient) getHistory() *dto.History {
	sentInvoices := make([]*dto.Invoice, 0)
	grandTotal := dto.GrandTotal{}
	for _, inv := range ccl.getAllInvoices() {
		if inv.Status != "" {
			sentInvoices = append(sentInvoices, inv)
			grandTotal.AddInvoice(inv)
		}
	}

	lastInvoice := sentInvoices[len(sentInvoices)-1]
	lastInvoice.Expenses = ccl.getExpenses(lastInvoice.No)

	return &dto.History{
		FirstInvoice: sentInvoices[0],
		Invoices:     sentInvoices[:len(sentInvoices)-1],
		LastInvoice:  lastInvoice,
		GrandTotal:   grandTotal,
	}
}
