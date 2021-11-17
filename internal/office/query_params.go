package office

import (
	"github.com/artsafin/go-coda"
	"oea-go/internal/common"
	"oea-go/internal/office/codaschema"
)

var InvoiceIsRecent common.QueryParam

func init() {
	InvoiceIsRecent = func(p *coda.ListRowsParameters) {
		p.Query = common.Query(codaschema.ID.Table.Invoices.Cols.IsRecent.ID, "true")
	}
}
