package common

import (
	"fmt"
	"oea-go/internal/codatypes"
	"testing"
	"time"
)

type timeTesting struct {
	*testing.T
}

func (test *timeTesting) timeMustParse(t time.Time, err error) time.Time {
	if err != nil {
		test.Fatal(err)
	}
	return t
}

func (test *timeTesting) parse(val string) time.Time {
	ttt, err := time.Parse(time.RFC3339, val)
	ttt = test.timeMustParse(ttt, err).In(codatypes.DatesLocation)

	return ttt
}

func TestIsWorkDay(t *testing.T) {
	test := timeTesting{t}
	dates := map[time.Time]bool{
		test.parse("2020-05-22T00:00:00Z"): true,
		test.parse("2020-05-23T00:00:00Z"): false,
		test.parse("2020-05-24T00:00:00Z"): false,
		test.parse("2020-05-25T00:00:00Z"): true,
		test.parse("2020-05-26T00:00:00Z"): true,
		test.parse("2020-05-27T00:00:00Z"): true,
		test.parse("2020-05-28T00:00:00Z"): true,
		test.parse("2020-05-29T00:00:00Z"): true,
		test.parse("2020-05-30T00:00:00Z"): false,
		test.parse("2020-05-31T00:00:00Z"): false,
		test.parse("2020-06-01T00:00:00Z"): true,
	}

	for t, result := range dates {
		actual := isWorkDay(t)
		if actual != result {
			test.Errorf("case %v: expected %v, actual %v", t, result, actual)
		}
	}
}

type DatesTestTable struct {
	date     time.Time
	days     int
	expected time.Time
}

func (t DatesTestTable) String() string {
	return fmt.Sprintf("{date: %v, days: %v, expected: %v}", t.date.Format(time.RFC822Z), t.days, t.expected.Format(time.RFC822Z))
}

func TestAddWorkingDate(t *testing.T) {
	test := timeTesting{t}

	dates := []DatesTestTable{
		{test.parse("2020-05-25T00:00:00Z"), 0, test.parse("2020-05-25T00:00:00Z")},
		{test.parse("2020-05-25T00:00:00Z"), 1, test.parse("2020-05-26T00:00:00Z")},
		{test.parse("2020-05-25T00:00:00Z"), 2, test.parse("2020-05-27T00:00:00Z")},
		{test.parse("2020-05-25T00:00:00Z"), -1, test.parse("2020-05-22T00:00:00Z")},
		{test.parse("2020-05-25T00:00:00Z"), -2, test.parse("2020-05-21T00:00:00Z")},

		{test.parse("2020-05-29T00:00:00Z"), 0, test.parse("2020-05-29T00:00:00Z")},
		{test.parse("2020-05-29T00:00:00Z"), 1, test.parse("2020-06-01T00:00:00Z")},
		{test.parse("2020-05-29T00:00:00Z"), 2, test.parse("2020-06-02T00:00:00Z")},
		{test.parse("2020-05-29T00:00:00Z"), -1, test.parse("2020-05-28T00:00:00Z")},

		{test.parse("2020-05-30T00:00:00Z"), 1, test.parse("2020-06-01T00:00:00Z")},
		{test.parse("2020-05-30T00:00:00Z"), 0, test.parse("2020-06-01T00:00:00Z")},
		{test.parse("2020-05-30T00:00:00Z"), -1, test.parse("2020-05-29T00:00:00Z")},
		{test.parse("2020-05-30T00:00:00Z"), -2, test.parse("2020-05-28T00:00:00Z")},

		{test.parse("2020-05-31T00:00:00Z"), 1, test.parse("2020-06-01T00:00:00Z")},
		{test.parse("2020-05-31T00:00:00Z"), 0, test.parse("2020-06-01T00:00:00Z")},
		{test.parse("2020-05-31T00:00:00Z"), -1, test.parse("2020-05-29T00:00:00Z")},
		{test.parse("2020-05-31T00:00:00Z"), -2, test.parse("2020-05-28T00:00:00Z")},

		{test.parse("2020-05-28T00:00:00Z"), -10, test.parse("2020-05-14T00:00:00Z")},
		{test.parse("2020-05-28T00:00:00Z"), -8, test.parse("2020-05-18T00:00:00Z")},
	}

	for idx, d := range dates {
		actual := AddWorkingDate(d.date, 0, 0, d.days)
		if actual != d.expected {
			test.Errorf("case #%v %v: actual %v", idx, d, actual)
		}
	}
}
