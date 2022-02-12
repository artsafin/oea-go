package months

import (
	"context"
	"fmt"
	"github.com/artsafin/coda-go-client/codaapi"
	"oea-go/internal/employee/codaschema"
	"sort"
)

func GetNearest(doc *codaschema.CodaDocument) (list []codaschema.Months, err error) {
	months, err := doc.ListMonths(context.Background(), codaapi.ListRows.SortBy(codaapi.RowsSortByNatural))
	if err != nil {
		return
	}

	for _, m := range months {
		if m.Near {
			list = append(list, m)
		}
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("[Near?] field must be set correctly in the [Months] table")
	}

	sort.Slice(list, func(i, j int) bool {
		// Reverse sort
		return list[j].Month.Unix() < list[i].Month.Unix()
	})

	return
}
