package codaschema

import (
	"oea-go/internal/codatypes"
	"sort"
)

func (inv Invoice) EntriesSorted() (data []Entries) {
	data = inv.InvoiceEntries.All()
	sort.Slice(data, func(i, j int) bool {
		return data[i].Sort < data[j].Sort
	})
	return
}

func (inv Invoice) RUBEntriesInEUR() codatypes.MoneyEur {
	return codatypes.MoneyEur(100 * (inv.EURSubtotal - inv.EUREntries))
}

func (inv Invoice) EUREntriesMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.EUREntries * 100)
}

func (inv Invoice) RUBEntriesMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.RUBEntries * 100)
}

func (inv Invoice) EURRUBExpectedMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.EURRUBExpected * 100)
}

func (inv Invoice) EURRUBActualMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.EURRUBActual * 100)
}

func (inv Invoice) EURSubtotalMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.EURSubtotal * 100)
}

func (inv Invoice) EURRoundingMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.EURRounding * 100)
}

func (inv Invoice) EURTotalMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.EURTotal * 100)
}

func (inv Invoice) RUBTotalMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.RUBTotal * 100)
}

func (inv Invoice) RUBActualMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.RUBActual * 100)
}

func (inv Invoice) RUBRateErrorMoney() codatypes.MoneyRub {
	return codatypes.MoneyRub(inv.RUBRateError * 100)
}

func (inv Invoice) HourRateMoney() codatypes.MoneyEur {
	return codatypes.MoneyEur(inv.HourRate * 100)
}
