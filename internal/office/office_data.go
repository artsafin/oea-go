package office

import (
	"github.com/artsafin/goda"
	"github.com/pkg/errors"
	"log"
	"sort"
	"time"

	"oea-go/internal/common"
	"oea-go/internal/office/dto"
)

type Data interface {
	WaitForInvoice() string
	GetLastInvoice() (string, error)
	GetInvoice(invoiceID string) (*dto.Invoice, error)
	GetAllInvoices() (dto.Invoices, error)
	GetExpenses(invoiceID string) ([]*dto.Expense, error)
	GetHistory() (*dto.History, error)
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

func (cd *codaData) WaitForInvoice() string {
	var invoiceID string
	log.Printf("Waiting for invoice...")
	timerChan := time.After(2 * time.Minute)
	for {
		var err error
		invoiceID, err = cd.GetLastInvoice()
		if err == nil {
			log.Println("Found planned invoice:", invoiceID)
			break
		}

		select {
		case t := <-timerChan:
			log.Fatalf("Stopped waiting at %v\n", t)
		default:
		}

		time.Sleep(500 * time.Millisecond)
		log.Print(".")
	}

	log.Print("\n")

	return invoiceID
}

func (cd *codaData) GetLastInvoice() (string, error) {
	lastInvoice, lastInvoiceErr := cd.doc.GetFormula(dto.Ids.CodaFormulas.LastInvoice)

	if lastInvoiceErr != nil {
		return "", lastInvoiceErr
	}

	formulaVal, ok := lastInvoice.Value.(string)

	if !ok || formulaVal == "" {
		return "", errors.New("last invoice not found")
	}

	return formulaVal, nil
}

func (cd *codaData) GetInvoice(invoiceID string) (*dto.Invoice, error) {
	params := goda.ListRowsParams{}.WithQuery(dto.Ids.Invoices.Cols.No, invoiceID)
	resp, err := cd.doc.ListRows(dto.Ids.Invoices.Id, &params)
	if err != nil {
		panic(err)
	}

	if len(resp.Items) == 0 {
		return nil, errors.Errorf("Invoice %s is empty", invoiceID)
	}

	firstRow := resp.Items[0]

	return dto.NewInvoiceFromRow(&firstRow), nil
}

func (cd *codaData) GetAllInvoices() (dto.Invoices, error) {
	params := goda.ListRowsParams{}.WithSortBy(goda.RowsSortBy_natural)
	resp, err := cd.doc.ListRows(dto.Ids.Invoices.Id, &params)
	if err != nil {
		return nil, err
	}

	invoices := make(dto.Invoices, len(resp.Items))
	for i, row := range resp.Items {
		invoices[i] = dto.NewInvoiceFromRow(&row)
	}

	sort.Sort(invoices)

	return invoices, nil
}

func (cd *codaData) GetExpenses(invoiceID string) ([]*dto.Expense, error) {
	resp, err := cd.doc.ListRows(dto.Ids.Expenses.Id, &goda.ListRowsParams{})
	if err != nil {
		return nil, err
	}

	expenses := make([]*dto.Expense, 0, len(resp.Items))

	for _, row := range resp.Items {
		if row.Values[dto.Ids.Expenses.Cols.Invoice] == invoiceID {
			expenses = append(expenses, dto.NewExpenseFromRow(&row))
		}
	}

	return expenses, nil
}

func (cd *codaData) GetHistory() (*dto.History, error) {
	sentInvoices := make([]*dto.Invoice, 0)
	grandTotal := dto.GrandTotal{}
	allInvoices, invErr := cd.GetAllInvoices()
	if invErr != nil {
		return nil, invErr
	}
	for _, inv := range allInvoices {
		if inv.IsHistory() {
			sentInvoices = append(sentInvoices, inv)
			grandTotal.AddInvoice(inv)
		}
	}

	if len(sentInvoices) == 0 {
		return nil, errors.New("history is empty")
	}

	lastInvoice := sentInvoices[len(sentInvoices)-1]
	expenses, expErr := cd.GetExpenses(lastInvoice.No)
	if expErr != nil {
		return nil, expErr
	}
	lastInvoice.Expenses = expenses

	return &dto.History{
		FirstInvoice: sentInvoices[0],
		Invoices:     sentInvoices[:len(sentInvoices)-1],
		LastInvoice:  lastInvoice,
		GrandTotal:   grandTotal,
	}, nil
}
