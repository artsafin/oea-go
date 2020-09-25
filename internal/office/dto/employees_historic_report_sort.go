package dto

import "strings"

type EmployeesMonthlyReportWithPercentCollection []EmployeesMonthlyReportWithPercent

func (o EmployeesMonthlyReportWithPercentCollection) Len() int {
	return len(o)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (o EmployeesMonthlyReportWithPercentCollection) Less(i, j int) bool {
	return strings.Compare(o[i].Month, o[j].Month) < 0
}

// Swap swaps the elements with indexes i and j.
func (o EmployeesMonthlyReportWithPercentCollection) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
