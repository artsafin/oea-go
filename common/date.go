package common

import (
	"time"
)

func isWorkDay(date time.Time) bool {
	return date.Weekday() != time.Sunday && date.Weekday() != time.Saturday
}

func skipDirection(days int) int {
	if days < 0 {
		return -1
	}
	return 1
}

func daysToSkip(days int) int {
	if days < 0 {
		return -1 * days
	}
	return days
}

func AddWorkingDate(date time.Time, years, months, days int) time.Time {
	date = date.AddDate(years, months, 0)

	dir := skipDirection(days)
	toSkip := daysToSkip(days)
	if !isWorkDay(date) && toSkip == 0 {
		toSkip = 1
	}
	for skippedDays := 0; skippedDays < toSkip; {
		date = date.AddDate(0, 0, dir)

		if isWorkDay(date) {
			skippedDays++
		}
	}

	return date
}
