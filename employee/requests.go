package employee

import (
	"errors"
	"fmt"
	"github.com/artsafin/go-coda"
	"oea-go/common"
	"oea-go/employee/dto"
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

func (requests *Requests) GetLastMonths() *dto.Months {
	params := coda.ListRowsParameters{}
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Months.Id, params)
	if err != nil {
		panic(err)
	}

	months := make(dto.Months, len(resp.Rows))

	for k, v := range resp.Rows {
		months[k] = dto.NewMonthFromRow(&v)
	}

	sort.Sort(sort.Reverse(months))

	return &months
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

func (requests *Requests) GetInvoices(month string) *dto.Invoices {
	params := coda.ListRowsParameters{
		Query: fmt.Sprintf("%s:\"%s\"", dto.Ids.Invoices.Cols.Month, month),
	}
	resp, err := requests.Client.ListTableRows(requests.DocId, dto.Ids.Invoices.Id, params)

	if err != nil {
		panic(err)
	}

	return dto.NewInvoicesFromRows(&resp)
}

//func (requests *Requests) GetCorrections(month string) (string, error) {
//}
//
//func (requests *Requests) GetTaxes(month string) (string, error) {
//}
//
//func (requests *Requests) GetPatents(month string) (string, error) {
//}
