package utils

import "github.com/360EntSecGroup-Skylar/excelize/v2"

func SetExcelColumnsWidth(file *excelize.File, sheet string, columnDefs map[string]float64) {
	for k, v := range columnDefs {
		file.SetColWidth(sheet, k, k, v)
	}
}