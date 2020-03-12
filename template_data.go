package main

import (
	officeDto "oea-go/office/dto"
	"os/user"
	"time"
)

type TemplateData struct {
	Timestamp time.Time
	Author    *user.User

	Office   OfficeTemplateData
	Employee EmployeeTemplateData
}

type OfficeTemplateData struct {
	Invoice       officeDto.Invoice
	ExpenseGroups officeDto.ExpenseGroupMap
	History       officeDto.History
}

type EmployeeTemplateData struct {
}
