package dto

import (
	"errors"
	"oea-go/internal/codatypes"
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
	errs := codatypes.NewErrorContainer()
	var err error

	if month.ID, err = codatypes.ToString(Ids.Months.Cols.ID, row); err != nil {
		errs.AddError(err)
	}

	if month.LastMonthDay, err = codatypes.ToDate(Ids.Months.Cols.Month, row); err != nil {
		errs.AddError(err)
	}
	if month.WorkDays, err = codatypes.ToUint16(Ids.Months.Cols.WorkDays, row); err != nil {
		errs.AddError(err)
	}
	if month.Year, err = codatypes.ToUint16(Ids.Months.Cols.Year, row); err != nil {
		errs.AddError(err)
	}
	if month.PreviousMonthLink, err = codatypes.ToString(Ids.Months.Cols.PreviousMonthLink, row); err != nil {
		errs.AddError(err)
	}
	if month.PreviousMonth, err = codatypes.ToDate(Ids.Months.Cols.PreviousMonth, row); err != nil {
		errs.AddError(err)
	}
	if month.IsCurrent, err = codatypes.ToBool(Ids.Months.Cols.IsCurrent, row); err != nil {
		errs.AddError(err)
	}

	errs.PanicIfError()

	return &month
}
