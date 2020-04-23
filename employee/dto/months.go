package dto

import (
	"errors"
	"oea-go/common"
	"time"

	"github.com/artsafin/go-coda"
)

type Months []*Month

func (m Months) Len() int {
	return len(m)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (m Months) Less(i, j int) bool {
	return m[i].LastMonthDay.Unix() < m[j].LastMonthDay.Unix()
}

// Swap swaps the elements with indexes i and j.
func (m Months) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Months) FindByName(name string) (*Month, error) {
	for _, v := range m {
		if v.ID == name {
			return v, nil
		}
	}

	return nil, errors.New("not found")
}

func (m Months) IndexOfYearMonth(of time.Time) (int, error) {
	for k, v := range m {
		if v.LastMonthDay.Month() == of.Month() && v.LastMonthDay.Year() == of.Year() {
			return k, nil
		}
	}

	return 0, errors.New("not found")
}

func (m Months) IndexOfTime(t time.Time) int {
	var curMonthIndex int
	var err error

	if curMonthIndex, err = m.IndexOfYearMonth(t); err != nil {
		curMonthIndex = 0
	}

	return curMonthIndex
}

type Month struct {
	LastMonthDay      *time.Time
	WorkDays          uint16
	Year              uint16
	ID                string
	PreviousMonthLink string
	PreviousMonth     *time.Time
	IsCurrent         bool
}

func (m Month) String() string {
	return m.ID
}

func NewMonthFromRow(row *coda.Row) *Month {
	month := Month{}
	errs := make([]common.UnexpectedFieldTypeErr, 0)
	var err *common.UnexpectedFieldTypeErr

	if month.ID, err = common.ToString(Ids.Months.Cols.ID, row); err != nil {
		errs = append(errs, *err)
	}

	if month.LastMonthDay, err = common.ToDate(Ids.Months.Cols.Month, row); err != nil {
		errs = append(errs, *err)
	}
	if month.WorkDays, err = common.ToUint16(Ids.Months.Cols.WorkDays, row); err != nil {
		errs = append(errs, *err)
	}
	if month.Year, err = common.ToUint16(Ids.Months.Cols.Year, row); err != nil {
		errs = append(errs, *err)
	}
	if month.PreviousMonthLink, err = common.ToString(Ids.Months.Cols.PreviousMonthLink, row); err != nil {
		errs = append(errs, *err)
	}
	if month.PreviousMonth, err = common.ToDate(Ids.Months.Cols.PreviousMonth, row); err != nil {
		errs = append(errs, *err)
	}
	if month.IsCurrent, err = common.ToBool(Ids.Months.Cols.IsCurrent, row); err != nil {
		errs = append(errs, *err)
	}

	if len(errs) > 0 {
		stringErr := ""
		for _, err := range errs {
			stringErr += err.Error() + "; "
		}

		panic(stringErr)
	}

	return &month
}
