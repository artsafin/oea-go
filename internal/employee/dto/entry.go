package dto

import (
	"fmt"
	"github.com/artsafin/go-coda"
	"html/template"
	"oea-go/internal/codatypes"
	"oea-go/internal/common"
	"oea-go/internal/employee/codaschema"
	"strings"
)

type Entry struct {
	Invoice   string
	Type      string
	Comment   string
	RUBAmount codatypes.MoneyRub
	EURAmount codatypes.MoneyEur
	Sort      uint16
}

func (e *Entry) LongComment() template.HTML {
	htmlComment := strings.Replace(e.Comment, "\n", "<br>", -1)
	return template.HTML(fmt.Sprintf("<code>%s</code><br>%s", e.Type, htmlComment))
}

func NewEntryFromRow(row *coda.Row) (*Entry, error) {
	entry := Entry{}
	errs := codatypes.NewErrorContainer()
	var err error

	if entry.Invoice, err = codatypes.ToString(codaschema.ID.Table.Entries.Cols.Invoice.ID, row); err != nil {
		errs.AddError(err)
	}
	if entry.Type, err = codatypes.ToString(codaschema.ID.Table.Entries.Cols.Type.ID, row); err != nil {
		errs.AddError(err)
	}
	if entry.Comment, err = codatypes.ToString(codaschema.ID.Table.Entries.Cols.Comment.ID, row); err != nil {
		errs.AddError(err)
	}
	if entry.RUBAmount, err = codatypes.ToRub(codaschema.ID.Table.Entries.Cols.RUBAmount.ID, row); err != nil {
		errs.AddError(err)
	}
	if entry.EURAmount, err = codatypes.ToEur(codaschema.ID.Table.Entries.Cols.EURAmount.ID, row); err != nil {
		errs.AddError(err)
	}
	if entry.Sort, err = codatypes.ToUint16(codaschema.ID.Table.Entries.Cols.Sort.ID, row); err != nil {
		errs.AddError(err)
	}

	return &entry, common.JoinErrors(errs)
}
