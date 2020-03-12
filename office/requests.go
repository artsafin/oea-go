package office

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/phouse512/go-coda"
	"oea-go/common"
	"oea-go/office/dto"
)

type Requests struct {
	Client *common.CodaClient
	DocId  string
}

func (requests *Requests) WaitForInvoice() string {
	var invoiceID string
	fmt.Printf("Waiting for invoice in doc %s...", requests.DocId)
	timerChan := time.After(2 * time.Minute)
	for {
		var err error
		invoiceID, err = requests.GetLastInvoice()
		if err == nil {
			fmt.Println("Found planned invoice:", invoiceID)
			break
		}

		select {
		case t := <-timerChan:
			fmt.Printf("Stopped waiting at %v\n", t)
			os.Exit(1)
		default:
		}

		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}

	return invoiceID
}

func (requests *Requests) GetLastInvoice() (string, error) {
	lastInvoice, err := requests.Client.GetFormula(requests.DocId, dto.Ids.CodaFormulas.LastInvoice)

	if err != nil {
		panic(err)
	}
	if lastInvoice.Formula.Value == "" {
		return "", errors.New("Last invoice not found")
	}

	str, ok := lastInvoice.Formula.Value.(string)

	if !ok {
		panic("Last invoice is not a string")
	}

	return str, nil
}

func (requests *Requests) GetInvoice(invoiceID string) *dto.Invoice {
	params := coda.ListRowsParameters{
		Query: fmt.Sprintf("\"%s\":\"%s\"", dto.Ids.Invoices.Cols.No, invoiceID),
	}
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Invoices.Id, params)
	if err != nil {
		panic(err)
	}

	if len(resp.Rows) == 0 {
		panic(fmt.Sprintf("Invoice %s is empty", invoiceID))
	}

	firstRow := resp.Rows[0]

	return dto.NewInvoiceFromRow(&firstRow)
}

func (requests *Requests) GetAllInvoices() dto.Invoices {
	params := coda.ListRowsParameters{
		SortBy: "natural",
	}
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Invoices.Id, params)
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

func (requests *Requests) GetExpenses(invoiceID string) []*dto.Expense {
	params := coda.ListRowsParameters{}
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Expenses.Id, params)
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

func (requests *Requests) GetHistory() *dto.History {
	sentInvoices := make([]*dto.Invoice, 0)
	grandTotal := dto.GrandTotal{}
	for _, inv := range requests.GetAllInvoices() {
		if inv.Status != "" {
			sentInvoices = append(sentInvoices, inv)
			grandTotal.AddInvoice(inv)
		}
	}

	lastInvoice := sentInvoices[len(sentInvoices)-1]
	lastInvoice.Expenses = requests.GetExpenses(lastInvoice.No)

	return &dto.History{
		FirstInvoice: sentInvoices[0],
		Invoices:     sentInvoices[:len(sentInvoices)-1],
		LastInvoice:  lastInvoice,
		GrandTotal:   grandTotal,
	}
}
