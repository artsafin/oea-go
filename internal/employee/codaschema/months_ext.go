package codaschema

import (
	"fmt"
	"time"
)

func (m Months) String() string {
	return m.ID
}

type MonthsList []*Months

func (m MonthsList) Len() int {
	return len(m)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (m MonthsList) Less(i, j int) bool {
	return m[i].LastMonthDay.Unix() < m[j].LastMonthDay.Unix()
}

// Swap swaps the elements with indexes i and j.
func (m MonthsList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MonthsList) FindByName(name string) (*Month, error) {
	for _, v := range m {
		if v.ID == name {
			return v, nil
		}
	}

	return nil, fmt.Errorf("month not found: %v", name)
}

func (m MonthsList) IndexOfYearMonth(of time.Time) (int, error) {
	for k, v := range m {
		if v.LastMonthDay.Month() == of.Month() && v.LastMonthDay.Year() == of.Year() {
			return k, nil
		}
	}

	return 0, fmt.Errorf("IndexOfYearMonth not found: %v", of)
}

func (m MonthsList) IndexOfTime(t time.Time) int {
	var curMonthIndex int
	var err error

	if curMonthIndex, err = m.IndexOfYearMonth(t); err != nil {
		curMonthIndex = 0
	}

	return curMonthIndex
}
