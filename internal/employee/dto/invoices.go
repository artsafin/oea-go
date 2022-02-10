package dto

import "strings"

type Invoices []*Invoice

func (invs Invoices) Len() int {
	return len(invs)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (invs Invoices) Less(i, j int) bool {
	monthCmp := strings.Compare(invs[i].MonthName, invs[j].MonthName)
	if monthCmp != 0 {
		return monthCmp < 0
	}

	return strings.Compare(invs[i].EmployeeName, invs[j].EmployeeName) < 0
}

// Swap swaps the elements with indexes i and j.
func (invs Invoices) Swap(i, j int) {
	invs[i], invs[j] = invs[j], invs[i]
}
