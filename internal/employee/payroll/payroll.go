package payroll

import (
	"context"
	"fmt"
	"oea-go/internal/employee/codaschema"
	"time"
)

func GetScheduleByMonth(doc *codaschema.CodaDocument, month string) (*time.Time, error) {
	schedule, err := doc.ListPayrollSchedule(context.Background())
	if err != nil {
		return nil, err
	}

	for _, s := range schedule {
		if s.Month.String() == month {
			return &s.ExecutionDate, nil
		}
	}

	return nil, fmt.Errorf("schedule not found for month " + month)
}
