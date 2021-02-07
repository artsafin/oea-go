package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"oea-go/internal/common"
	"oea-go/internal/employee/dto"
)

const payrollReportSheetName = "Sheet1"

func rowColZeroIndexesToCellAddr(col, row int) string {
	if col > 25 {
		panic("too much columns in payroll report")
	}
	return fmt.Sprintf("%c%d", 65+col, row+1)
}

type styledValue struct {
	styleID int
	value   interface{}
}

type sheetRef struct {
	file        *excelize.File
	name        string
	boldStyleId int
}

func newSheetRef(file *excelize.File, name string) sheetRef {
	return sheetRef{
		file:        file,
		name:        name,
		boldStyleId: -1,
	}
}

func (sheet *sheetRef) writeRowAndIncr(rowZeroIndexed *int, values ...interface{}) {
	sheet.writeRow(*rowZeroIndexed, values...)
	*rowZeroIndexed++
}

func (sheet *sheetRef) writeRow(rowZeroIndexed int, values ...interface{}) {
	for idx, value := range values {
		addr := rowColZeroIndexesToCellAddr(idx, rowZeroIndexed)

		if styledVal, isStyled := value.(styledValue); isStyled {
			sheet.file.SetCellValue(sheet.name, addr, styledVal.value)
			sheet.file.SetCellStyle(sheet.name, addr, addr, styledVal.styleID)
		} else {
			sheet.file.SetCellValue(sheet.name, addr, value)
		}
	}
}
func (sheet *sheetRef) bold(value interface{}) styledValue {
	if sheet.boldStyleId == -1 {
		sheet.boldStyleId, _ = sheet.file.NewStyle(`{"font":{"bold":true}}`)
	}

	return styledValue{sheet.boldStyleId, value}
}

func RenderPayrollReport(wr io.Writer, invoices dto.Invoices) error {
	f := excelize.NewFile()
	sheet := newSheetRef(f, payrollReportSheetName)

	sheet.writeRow(0,
		sheet.bold("Employee name"),
		sheet.bold("Subject"),
		sheet.bold("Amount, RUB"),
		sheet.bold("Amount, EUR"),
		sheet.bold("Additional comments"),
	)

	f.SetColWidth(payrollReportSheetName, "A", "A", 35)
	f.SetColWidth(payrollReportSheetName, "B", "B", 50)
	f.SetColWidth(payrollReportSheetName, "C", "C", 18)
	f.SetColWidth(payrollReportSheetName, "D", "D", 18)
	f.SetColWidth(payrollReportSheetName, "E", "E", 50)
	f.SetPanes(payrollReportSheetName, `{"freeze":true,"split":false,"x_split":1,"y_split":0,"top_left_cell":"B1","active_pane":"topRight","panes":[{"sqref":"B1","active_cell":"B1","pane":"topRight"}]}`)

	var grandTotalRub common.MoneyRub
	var grandTotalEur common.MoneyEur
	rowNum := 1

	for _, invoice := range invoices {
		if invoice.BaseSalaryRub > 0 {
			sheet.writeRowAndIncr(&rowNum,
				invoice.EmployeeName,
				fmt.Sprintf("Salary %s", invoice.DateYm()),
				invoice.BaseSalaryRub.Number(),
			)
		} else {
			sheet.writeRowAndIncr(&rowNum,
				invoice.EmployeeName,
				fmt.Sprintf("Salary %s", invoice.DateYm()),
				"",
				invoice.BaseSalaryEur.Number(),
			)
		}
		if invoice.BankFees > 0 {
			sheet.writeRowAndIncr(&rowNum, invoice.EmployeeName, "Bank fees", invoice.BankFees.Number())
		}
		if invoice.PatentRub > 0 {
			sheet.writeRowAndIncr(&rowNum, invoice.EmployeeName, "Patent", invoice.PatentRub.Number())
		}
		if invoice.TaxesRub > 0 {
			sheet.writeRowAndIncr(&rowNum, invoice.EmployeeName, "Taxes", invoice.TaxesRub.Number())
		}
		if invoice.RateErrorPrevMon > 0 {
			sheet.writeRowAndIncr(
				&rowNum,
				invoice.EmployeeName,
				"Currency rate variance from previous month",
				invoice.RateErrorPrevMon.Number(),
				"",
				fmt.Sprintf(
					"Expected: %s, actual: %s",
					invoice.PrevInvoice.EurRubExpected.Number(),
					invoice.PrevInvoice.EurRubActual.Number(),
				),
			)
		}

		for _, corr := range invoice.Corrections {
			addComment := fmt.Sprintf("%s - %s", corr.Category, corr.SubCategory())
			if corr.AbsoluteCorrectionEur > 0 {
				addComment += fmt.Sprintf("\nFor reference: %s", corr.AbsoluteCorrectionEur)
			}
			sheet.writeRowAndIncr(
				&rowNum,
				invoice.EmployeeName,
				corr.Comment,
				corr.TotalCorrectionRub.Number(),
				"",
				addComment,
			)
		}

		sheet.writeRowAndIncr(
			&rowNum,
			sheet.bold(fmt.Sprintf("%s Subtotal", invoice.EmployeeName)),
			"",
			invoice.RequestedSubtotalRub.Number(),
			invoice.RequestedSubtotalEur.Number(),
			fmt.Sprintf("EURRUB %s", invoice.EurRubExpected),
		)

		if invoice.RoundingPrevMonEur > 0 {
			sheet.writeRowAndIncr(
				&rowNum,
				invoice.EmployeeName,
				"Rounding in previous month",
				"",
				invoice.RoundingPrevMonEur.Neg().Number(),
			)
		}

		sheet.writeRowAndIncr(
			&rowNum,
			invoice.EmployeeName,
			"Rounding",
			"",
			invoice.Rounding.Number(),
		)

		sheet.writeRowAndIncr(
			&rowNum,
			sheet.bold(fmt.Sprintf("%s Total", invoice.EmployeeName)),
			"",
			"",
			invoice.AmountRequestedEur.Number(),
		)

		grandTotalRub += invoice.AmountRequestedEur.ToRub(invoice.EurRubExpected)
		grandTotalEur += invoice.AmountRequestedEur
	}

	sheet.writeRow(
		rowNum,
		sheet.bold("Grand Total"),
		"",
		grandTotalRub.Number(),
		grandTotalEur.Number(),
		fmt.Sprintf("Total %v employees", len(invoices)),
	)

	wrapTextStyleID, _ := sheet.file.NewStyle(`{"alignment":{"wrap_text":true}}`)
	f.SetCellStyle(payrollReportSheetName, "B2", fmt.Sprintf("E%d", rowNum), wrapTextStyleID)

	if writeErr := f.Write(wr); writeErr != nil {
		return writeErr
	}

	return nil
}
