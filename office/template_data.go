package office

import (
	"oea-go/office/dto"
	"time"
)

type TemplateData struct {
	Timestamp time.Time

	Office    OfficeTemplateData
	Employees dto.EmployeesHistoricReport
}

type OfficeTemplateData struct {
	PrevInvoice   dto.Invoice
	Invoice       dto.Invoice
	ExpenseGroups dto.ExpenseGroupMap
	History       dto.History
}
