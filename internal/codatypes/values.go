package codatypes

import (
	"errors"
	"github.com/artsafin/go-coda"
	"math"
	"strings"
	"time"
)

func ToStructuredValue(colID string, row *coda.Row) (sv StructuredValue, err error) {
	if row.Values[colID] == nil {
		return StructuredValue{}, errors.New("column missing")
	}
	var values map[string]interface{}
	var ok bool

	if values, ok = row.Values[colID].(map[string]interface{}); !ok {
		return StructuredValue{}, errors.New("column is not map")
	}

	if sv.Context, ok = values["@context"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast @context to string")
	}

	if sv.Type, ok = values["@type"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast @type to string")
	}

	if sv.AdditionalType, ok = values["additionalType"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast additionalType to string")
	}

	if sv.Name, ok = values["name"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast name to string")
	}

	if sv.Url, ok = values["url"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast url to string")
	}

	if sv.TableId, ok = values["tableId"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast tableId to string")
	}

	if sv.RowId, ok = values["rowId"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast rowId to string")
	}

	if sv.TableUrl, ok = values["tableUrl"].(string); !ok {
		return StructuredValue{}, errors.New("cannot cast tableUrl to string")
	}

	return
}

func ToString(colID string, row *coda.Row) (string, error) {
	if row.Values[colID] == nil {
		return "", nil
	}
	if value, ok := row.Values[colID].(string); ok {
		return strings.Trim(value, "`"), nil
	}

	return "", UnexpectedFieldTypeErr{colID, "string", row}
}

func ToDate(colID string, row *coda.Row) (*time.Time, error) {
	if row.Values[colID] == nil {
		return nil, nil
	}
	if value, ok := row.Values[colID].(string); ok {
		if value == "" {
			return nil, nil
		}
		if time, err := time.Parse(time.RFC3339, value); err == nil {
			time = time.In(DatesLocation)
			return &time, nil
		}
	}

	return nil, UnexpectedFieldTypeErr{colID, "RFC3339 date", row}
}

func ToFloat64(colID string, row *coda.Row) (float64, error) {
	if row.Values[colID] == nil {
		return 0, nil
	}
	switch v := row.Values[colID].(type) {
	case float64:
		return v, nil
	case string:
		return 0, nil
	default:
		return 0, UnexpectedFieldTypeErr{colID, "float64", row}
	}
}

func ToBool(colID string, row *coda.Row) (bool, error) {
	if row.Values[colID] == nil {
		return false, nil
	}
	switch v := row.Values[colID].(type) {
	case bool:
		return v, nil
	case string:
		return false, nil
	default:
		return false, UnexpectedFieldTypeErr{colID, "bool", row}
	}
}

func ToUint16(colID string, row *coda.Row) (uint16, error) {
	if v, err := ToFloat64(colID, row); err == nil {
		return uint16(v), nil
	}
	return 0, UnexpectedFieldTypeErr{colID, "uint16", row}
}

func ToEur(colID string, row *coda.Row) (MoneyEur, error) {
	if v, err := ToFloat64(colID, row); err == nil {
		return MoneyEur(math.Round(v * 100)), nil
	}
	return 0, UnexpectedFieldTypeErr{colID, "MoneyEur", row}
}

func ToRub(colID string, row *coda.Row) (MoneyRub, error) {
	if v, err := ToFloat64(colID, row); err == nil {
		return MoneyRub(math.Round(v * 100)), nil
	}
	return 0, UnexpectedFieldTypeErr{colID, "MoneyRub", row}
}
