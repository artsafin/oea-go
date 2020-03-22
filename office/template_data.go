package office

import (
	officeDto "oea-go/office/dto"
	"time"
)

type TemplateData struct {
	Timestamp time.Time

	Office    OfficeTemplateData
	Employees *officeDto.EmployeesPaymentCategories
}

type OfficeTemplateData struct {
	PrevInvoice   officeDto.Invoice
	Invoice       officeDto.Invoice
	ExpenseGroups officeDto.ExpenseGroupMap
	History       officeDto.History
}
