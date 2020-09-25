package dto

import (
	"fmt"
	"sort"
)

type EmployeesHistoricReport struct {
	CurrentMonth *EmployeesMonthlyReport
	RestMonths   EmployeesMonthlyReportWithPercentCollection
}

type EmployeesMonthlyReportWithPercent struct {
	*EmployeesMonthlyReport
	CategoryPercents map[string]float32
}

func (mrwp EmployeesMonthlyReportWithPercent) GetPercentForCategory(cat string) string {
	if perc, found := mrwp.CategoryPercents[cat]; found {
		return fmt.Sprintf("%+.2f%%", 100 * (perc - 1))
	}
	return "0%"
}

func (rep *EmployeesHistoricReport) AppendHistoricReport(report *EmployeesMonthlyReport) {
	catPercents := make(map[string]float32)
	for k, _ := range report.Categories {
		catPercents[k] = 0
	}

	rm := EmployeesMonthlyReportWithPercent{
		EmployeesMonthlyReport: report,
		CategoryPercents:       catPercents,
	}
	rep.RestMonths = append(rep.RestMonths, rm)
}

func (rep *EmployeesHistoricReport) RunPostCalculations() {
	rep.calculatePercents()
	rep.sortRestMonths()
}

func (rep *EmployeesHistoricReport) calculatePercents() {
	if rep.CurrentMonth == nil {
		panic("unable to calculate percents of employees historic report: current month is not set")
	}

	for _, restRep := range rep.RestMonths {
		for catName, _ := range restRep.CategoryPercents {
			currentCat, currentFound := rep.CurrentMonth.Categories[catName]
			restCat, restFound := restRep.Categories[catName]
			if !currentFound || !restFound {
				continue
			}
			percent := float32(0)
			if currentCat.Total != 0 {
				percent = float32(restCat.Total) / float32(currentCat.Total)
			}
			restRep.CategoryPercents[catName] = percent
		}
	}
}

func (rep *EmployeesHistoricReport) sortRestMonths() {
	sort.Sort(sort.Reverse(rep.RestMonths))
}

