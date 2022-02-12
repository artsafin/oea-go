package excel

import (
	"fmt"
	x "github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"log"
	"oea-go/internal/employee/codaschema"
	"sort"
	"strings"
)

const payrollReportSheetName = "Sheet1"

type columnGroup struct {
	name  string
	color string
	sort  uint8
}

var (
	groupPrelude = columnGroup{
		name:  "",
		color: "#ffffff",
		sort:  0,
	}
	groupEarnings = columnGroup{
		name:  "EARNINGS",
		color: "#fce4d6",
		sort:  1,
	}
	groupDeductions = columnGroup{
		name:  "DEDUCTIONS",
		color: "#d9e1f2",
		sort:  2,
	}
	groupSubcontractorTax = columnGroup{
		name:  "SUBCONTRACTOR'S TAX",
		color: "#ededed",
		sort:  3,
	}
	groupEmployerContributions = columnGroup{
		name:  "EMPLOYER'S CONTRIBUTIONS",
		color: "#e2efda",
		sort:  4,
	}
	groupTotals = columnGroup{
		name:  "TOTALS",
		color: "#fff2cc",
		sort:  5,
	}
)

type columnValues map[int]interface{}
type column struct {
	header string
	group  columnGroup
	sort   uint16
	vals   columnValues
}

func (c column) GetSort() uint32 {
	return uint32(c.group.sort)<<16 + uint32(c.sort)
}

func getGroupByEntry(e codaschema.Entries) columnGroup {
	special := map[string]columnGroup{
		"pension fund fixed":   groupEmployerContributions,
		"pension fund percent": groupEmployerContributions,
		"social insurance":     groupEmployerContributions,

		"org payments":               groupSubcontractorTax,
		"individual entrepreneur 6%": groupSubcontractorTax,
		"self-employed tax":          groupSubcontractorTax,
		"patent":                     groupSubcontractorTax,
	}

	if group, ok := special[strings.ToLower(e.Type.String())]; ok {
		return group
	}

	if e.EURTotal() < 0 {
		return groupDeductions
	}

	return groupEarnings
}

func columnsFromInvoices(invoices []codaschema.Invoice) (cols []column) {
	cm := make(map[string]*column)

	// Prelude
	cm["#"] = &column{"", groupPrelude, 0, map[int]interface{}{}}
	cm["Location"] = &column{"Location", groupPrelude, 1, map[int]interface{}{}}
	cm["Entity name"] = &column{"Entity name", groupPrelude, 2, map[int]interface{}{}}
	cm["Employee"] = &column{"Employee", groupPrelude, 3, map[int]interface{}{}}

	// Totals
	cm["Net salaries"] = &column{"Net salaries", groupTotals, 0, map[int]interface{}{}}
	cm["Company cost"] = &column{"Company cost", groupTotals, 1, map[int]interface{}{}}
	cm["Rounding"] = &column{"Rounding", groupTotals, 2, map[int]interface{}{}}
	cm["Total"] = &column{"Total", groupTotals, 3, map[int]interface{}{}}

	for invoiceIndex, inv := range invoices {
		cm["#"].vals[invoiceIndex] = cell{value: invoiceIndex + 1}
		cm["Location"].vals[invoiceIndex] = cell{value: inv.Employee.FirstMaybe().Location.String()}
		cm["Entity name"].vals[invoiceIndex] = cell{value: inv.SenderDetails.String()}
		cm["Employee"].vals[invoiceIndex] = cell{value: inv.Employee.String()}

		var netSalariesEUR float64
		var companyCostsEUR float64

		for _, entry := range inv.EntriesSorted() {
			var ok bool

			group := getGroupByEntry(entry)

			if _, ok = cm[entry.Type.String()]; !ok {
				cm[entry.Type.String()] = &column{
					header: entry.Type.String(),
					group:  group,
					sort:   uint16(entry.Sort),
					vals:   make(columnValues),
				}
			}

			comment := entry.Comment
			if entry.RUBAmount != 0 {
				comment = fmt.Sprintf("\nOriginal amount: %s\n%s", entry.RUBAmountMoney().String(), entry.Comment)
			}

			var cellVal numberCell
			if cellVal, ok = cm[entry.Type.String()].vals[invoiceIndex].(numberCell); !ok {
				cellVal = numberCell{value: entry.EURTotal(), comment: comment}
			} else {
				cellVal.value += entry.EURTotal()
				cellVal.comment += "\n" + comment
			}

			cm[entry.Type.String()].vals[invoiceIndex] = cellVal

			if group == groupDeductions || group == groupEarnings {
				netSalariesEUR += entry.EURTotal()
			} else {
				companyCostsEUR += entry.EURTotal()
			}
		}

		excelRoundingError := inv.EURTotal - inv.EURRounding - netSalariesEUR - companyCostsEUR
		cm["Net salaries"].vals[invoiceIndex] = numberCell{value: netSalariesEUR}
		cm["Company cost"].vals[invoiceIndex] = numberCell{value: companyCostsEUR}
		cm["Rounding"].vals[invoiceIndex] = numberCell{value: inv.EURRounding + excelRoundingError}
		cm["Total"].vals[invoiceIndex] = numberCell{value: inv.EURTotal}
	}

	for _, col := range cm {
		cols = append(cols, *col)
	}

	sort.Slice(cols, func(i, j int) bool {
		return cols[i].GetSort() < cols[j].GetSort()
	})

	return
}

type columnColorRange struct {
	from, to int
	color    string
}

func styleReport(sheet *sheetRef, columnRanges []columnColorRange, totalColumns int, totalRows int) error {
	for _, r := range columnRanges {
		wraperr := func(step string, err error) error {
			return fmt.Errorf("range style error (%s) %v: %w", step, r, err)
		}

		styleIDContent, _ := sheet.file.NewStyle(&x.Style{
			Fill: x.Fill{Type: "pattern", Pattern: 1, Color: []string{r.color}},
			Border: []x.Border{
				{
					Type:  "right",
					Color: "AAAAAA",
					Style: 1,
				},
				{
					Type:  "bottom",
					Color: "AAAAAA",
					Style: 1,
				},
			},
			NumFmt: 4,
		})
		styleIDHeader, _ := sheet.file.NewStyle(&x.Style{
			Fill: x.Fill{Type: "pattern", Pattern: 1, Color: []string{r.color}},
			Font: &x.Font{Bold: true},
			Border: []x.Border{
				{
					Type:  "right",
					Color: "999999",
					Style: 1,
				},
				{
					Type:  "bottom",
					Color: "999999",
					Style: 1,
				},
			},
		})

		// Header styles
		pHeaderTopLeft, err := x.CoordinatesToCellName(r.from+1, 1)
		if err != nil {
			return wraperr("pHeaderTopLeft", err)
		}
		pHeaderTopRight, err := x.CoordinatesToCellName(r.to+1, 1)
		if err != nil {
			return wraperr("pHeaderTopRight", err)
		}
		pHeaderBottomRight, err := x.CoordinatesToCellName(r.to+1, 2)
		if err != nil {
			return wraperr("pHeaderBottomRight", err)
		}

		err = sheet.file.MergeCell(sheet.name, pHeaderTopLeft, pHeaderTopRight)
		if err != nil {
			return wraperr("MergeCell", err)
		}

		sheet.file.SetCellStyle(sheet.name, pHeaderTopLeft, pHeaderBottomRight, styleIDHeader)

		// Content styles
		pContentTopLeft, err := x.CoordinatesToCellName(r.from+1, 3)
		if err != nil {
			return wraperr("pContentTopLeft", err)
		}
		pContentBottomRight, err := x.CoordinatesToCellName(r.to+1, totalRows+2)
		if err != nil {
			return wraperr("pContentBottomRight", err)
		}

		sheet.file.SetCellStyle(sheet.name, pContentTopLeft, pContentBottomRight, styleIDContent)
	}

	firstColName, _ := x.ColumnNumberToName(1)
	lastColName, _ := x.ColumnNumberToName(totalColumns)
	sheet.file.SetColWidth(sheet.name, firstColName, lastColName, 18)

	sheet.file.SetColWidth(sheet.name, "A", "A", 8)
	sheet.file.SetColWidth(sheet.name, "B", "B", 15)
	sheet.file.SetColWidth(sheet.name, "C", "C", 30)

	return nil
}

func RenderPayrollReport(wr io.Writer, invoices []codaschema.Invoice) error {
	f := x.NewFile()
	//f.SetPanes(payrollReportSheetName, `{"freeze":true,"split":false,"x_split":1,"y_split":0,"top_left_cell":"B1","active_pane":"topRight","panes":[{"sqref":"B1","active_cell":"B1","pane":"topRight"}]}`)

	sheet := newSheetRef(f, payrollReportSheetName)

	columns := columnsFromInvoices(invoices)
	//log.Printf("RenderPayrollReport columns: %+v", columns)

	currentGroup := columnGroup{name: "_"}
	var columnRanges []columnColorRange

	var idx int
	var col column
	for idx, col = range columns {
		if currentGroup.name != col.group.name {
			err := sheet.setCellValue(rowCol{0, idx}, col.group.name)
			if err != nil {
				log.Printf("RenderPayrollReport currentGroup setCellValue: %v", err)
			}

			if len(columnRanges) > 0 {
				columnRanges[len(columnRanges)-1].to = idx - 1
				columnRanges[len(columnRanges)-1].color = currentGroup.color
			}
			columnRanges = append(columnRanges, columnColorRange{from: idx})

			currentGroup = col.group
		}

		err := sheet.setCellValue(rowCol{1, idx}, col.header)
		if err != nil {
			log.Printf("RenderPayrollReport header setCellValue: %v", err)
		}

		err = sheet.writeSparseColumn(rowCol{2, idx}, col.vals)
		if err != nil {
			return err
		}
	}
	columnRanges[len(columnRanges)-1].to = idx
	columnRanges[len(columnRanges)-1].color = currentGroup.color

	//log.Printf("RenderPayrollReport columnRanges: %+v", columnRanges)

	err := styleReport(&sheet, columnRanges, len(columns), len(invoices))
	if err != nil {
		log.Printf("RenderPayrollReport styleReport error: %v", err)
	}

	if writeErr := f.Write(wr); writeErr != nil {
		return writeErr
	}

	return nil
}
