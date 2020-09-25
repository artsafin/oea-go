package common

import (
	"fmt"
	"github.com/artsafin/goda"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"time"
)

var datesLocation *time.Location

var enPrinter *message.Printer

func init() {
	datesLocation = time.FixedZone("UTC+3", 3*60*60)
	enPrinter = message.NewPrinter(language.AmericanEnglish)
}

type MoneyRub int64
type MoneyEur int64

func (v MoneyRub) String() string {
	return enPrinter.Sprintf("%.2f ₽", float64(v)/100)
}

func (v MoneyEur) String() string {
	return enPrinter.Sprintf("€%.2f", float64(v)/100)
}

type UnexpectedFieldTypeErr struct {
	fieldID      string
	expectedType string
	row          *goda.Row
}

func (e *UnexpectedFieldTypeErr) Error() string {
	v := e.row.Values[e.fieldID]
	return fmt.Sprintf("Unexpected type for field %s: expected %s. Got value: %#v of type %T", e.fieldID, e.expectedType, v, v)
}

func ToString(colID string, row *goda.Row) (string, *UnexpectedFieldTypeErr) {
	if row.Values[colID] == nil {
		return "", nil
	}
	if value, ok := row.Values[colID].(string); ok {
		return value, nil
	}

	return "", &UnexpectedFieldTypeErr{colID, "string", row}
}

func ToDate(colID string, row *goda.Row) (*time.Time, *UnexpectedFieldTypeErr) {
	if row.Values[colID] == nil {
		return nil, nil
	}
	if value, ok := row.Values[colID].(string); ok {
		if value == "" {
			return nil, nil
		}
		if time, err := time.Parse(time.RFC3339, value); err == nil {
			time = time.In(datesLocation)
			return &time, nil
		}
	}

	return nil, &UnexpectedFieldTypeErr{colID, "RFC3339 date", row}
}

func ToFloat64(colID string, row *goda.Row) (float64, *UnexpectedFieldTypeErr) {
	if row.Values[colID] == nil {
		return 0, nil
	}
	switch v := row.Values[colID].(type) {
	case float64:
		return v, nil
	case string:
		return 0, nil
	default:
		return 0, &UnexpectedFieldTypeErr{colID, "float64", row}
	}
}

func ToBool(colID string, row *goda.Row) (bool, *UnexpectedFieldTypeErr) {
	if row.Values[colID] == nil {
		return false, nil
	}
	switch v := row.Values[colID].(type) {
	case bool:
		return v, nil
	case string:
		return false, nil
	default:
		return false, &UnexpectedFieldTypeErr{colID, "bool", row}
	}
}

func ToUint16(colID string, row *goda.Row) (uint16, *UnexpectedFieldTypeErr) {
	if v, err := ToFloat64(colID, row); err == nil {
		return uint16(v), nil
	}
	return 0, &UnexpectedFieldTypeErr{colID, "uint16", row}
}

func ToEur(colID string, row *goda.Row) (MoneyEur, *UnexpectedFieldTypeErr) {
	if v, err := ToFloat64(colID, row); err == nil {
		return MoneyEur(math.Round(v * 100)), nil
	}
	return 0, &UnexpectedFieldTypeErr{colID, "MoneyEur", row}
}

func ToRub(colID string, row *goda.Row) (MoneyRub, *UnexpectedFieldTypeErr) {
	if v, err := ToFloat64(colID, row); err == nil {
		return MoneyRub(math.Round(v * 100)), nil
	}
	return 0, &UnexpectedFieldTypeErr{colID, "MoneyRub", row}
}
