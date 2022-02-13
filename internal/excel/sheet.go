package excel

import (
	"encoding/json"
	"fmt"
	x "github.com/xuri/excelize/v2"
)

type rowCol struct {
	row, col int
}

type commenter interface {
	Comment() string
}

type sheetRef struct {
	file *x.File
	name string
}

func newSheetRef(file *x.File, name string) sheetRef {
	return sheetRef{
		file: file,
		name: name,
	}
}

func (sheet *sheetRef) setCellValue(addr rowCol, value interface{}) error {
	axis, err := x.CoordinatesToCellName(addr.col+1, addr.row+1)
	if err != nil {
		return err
	}

	if v, ok := value.(floater); ok {
		sheet.file.SetCellValue(sheet.name, axis, v.AsFloat64())
	} else if v, ok := value.(fmt.Stringer); ok {
		sheet.file.SetCellValue(sheet.name, axis, v.String())
	} else {
		sheet.file.SetCellValue(sheet.name, axis, value)
	}

	if val, ok := value.(commenter); ok && len(val.Comment()) > 0 {
		jsonComment, _ := json.Marshal(val.Comment())
		// jsonComment will be a safe json-string: special chars quoted and wrapped with double-quotes
		sheet.file.AddComment(sheet.name, axis, fmt.Sprintf(`{"author":"Comment: ","text":%s}`, jsonComment))
	}

	return nil
}

func (sheet *sheetRef) writeSparseColumn(startFrom rowCol, cells columnValues) error {
	for row, value := range cells {
		err := sheet.setCellValue(rowCol{
			row: startFrom.row + row,
			col: startFrom.col,
		}, value)

		if err != nil {
			return err
		}
	}

	return nil
}
