package employee

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/phouse512/go-coda"
	"oea-go/common"
	"oea-go/employee/dto"
)

type Requests struct {
	Client *common.CodaClient
	DocId  string
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
