package common

import (
	"time"
)

func isWorkDay(date time.Time) bool {
	return date.Weekday() != time.Sunday && date.Weekday() != time.Saturday
}

func AddWorkingDate(date *time.Time, years, months, days int) time.Time {
	result := date.AddDate(years, months, days)

	sign := 1
	if days < 0 {
		sign = -1
	}

	for !isWorkDay(result) {
		result = result.AddDate(0, 0, sign)
	}

	return result
}
