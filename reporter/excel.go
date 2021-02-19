package reporter

import (
	"log"
	"sort"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/srgio-es/tcvolumeutils/model"
	"github.com/srgio-es/tcvolumeutils/utils"
)

const (
	sheetHeaderStyle = `{
		"font":
		{
			"bold": true
		},
		"alignment":
		{
			"horizontal": "center"
		}

	}`

	summarysheet string = "Summary"

	summaryFilter string = `{"column":"B","expression":"x > 0"}`
)

type ExcelReporter struct {
	reportFile	string
}

func NewExcelReporter(reportFile string) *ExcelReporter {
	return &ExcelReporter{reportFile: reportFile}
}

func (e *ExcelReporter) GenerateMissingFilesReport(data map[string][]*model.MissingFile) {
	f := excelize.NewFile()

	if (f.GetSheetName(0) != summarysheet) {
		f.SetSheetName("Sheet1", summarysheet)
	}

	headerStyle, err := f.NewStyle(sheetHeaderStyle)

	if err != nil {
		log.Fatal(err)
	}

	f.SetCellValue(summarysheet, "A1", "Volume Name")
	f.SetCellValue(summarysheet, "B1", "Missing Files")

	f.SetCellStyle(summarysheet, "A1", "B1", headerStyle)

	volumes := make([]string, 0, len(data))

	for volumeName := range data {
		volumes = append(volumes, volumeName)
	}

	sort.Strings(volumes)

	for i, k := range volumes {
		f.NewSheet(k)
		printVolumeSheetHeader(f, k)
		printVolumeSheetValues(f, k, data[k])
		printSummaryLine(f, i, k, len(data[k]))

		utils.SetExcelColumnsWidth(f, k, map[string]float64{
			"A": 84,
			"B": 7,
			"C": 17,
			"D": 20,
			"E": 5.5,
			"F": 95,
		})
	}

	utils.SetExcelColumnsWidth(f, summarysheet, map[string]float64 {
		"A": 30,
		"B": 12,
	})

	filterSummary(f)

	if err := f.SaveAs(e.reportFile); err != nil {
		log.Fatal(err)
	}

}

func printSummaryLine(file *excelize.File, order int, volume string, missingFilesQty int) {
	r := int64(order) + 2
	row := strconv.FormatInt(r , 10)
	file.SetCellValue(summarysheet, "A" + row, volume)
	file.SetCellValue(summarysheet, "B" + row, missingFilesQty)

	if missingFilesQty > 0 {
		file.SetCellHyperLink(summarysheet, "A" + row, "'"+volume+"'"+"!A1", "Location")
	}
}

func printVolumeSheetValues(file *excelize.File, sheet string, data []*model.MissingFile) {

	dateStyle, err := file.NewStyle(`{
		"number_format": 22
	}`)

	if err != nil {
		log.Fatal(err)
	}

	if len(data) > 0 {
		for i, v := range data {
			row := strconv.FormatInt(int64(i) + 2, 10)
			file.SetCellValue(sheet, "A" + row, v.DatasetName)
			file.SetCellValue(sheet, "B" + row, v.Version)
			file.SetCellValue(sheet, "C" + row, v.UID)
			file.SetCellValue(sheet, "D" + row, v.ModifiedDate)
			file.SetCellStyle(sheet, "D" + row, "D"+row, dateStyle)
			file.SetCellValue(sheet, "E" + row, v.Site)
			file.SetCellValue(sheet, "F" + row, v.FileLocation)
		}
	} else {
		file.SetSheetVisible(sheet, false)
	}
}


func printVolumeSheetHeader(file *excelize.File, sheet string) {
	headerStyle, err := file.NewStyle(sheetHeaderStyle)

	if err != nil {
		log.Fatal(err)
	}

	file.SetCellValue(sheet, "A1", "Dataset Name")
	file.SetCellValue(sheet, "B1", "Version")
	file.SetCellValue(sheet, "C1", "UID")
	file.SetCellValue(sheet, "D1", "Last Modification Date")
	file.SetCellValue(sheet, "E1", "Site")
	file.SetCellValue(sheet, "F1", "File Location")

	file.SetCellStyle(sheet, "A1", "F1", headerStyle)
}


func filterSummary(file *excelize.File) {
	err := file.AutoFilter(summarysheet, "A1", "B500", summaryFilter)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := file.GetRows(summarysheet)

	if(err != nil) {
		log.Fatal(err)
	}

	for i, row := range rows {
		missingFiles := row[1]

		if i > 0 {
			missingfilesInt, err := strconv.ParseInt(missingFiles, 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			if missingfilesInt == 0 {
				file.SetRowVisible(summarysheet, i+1, false)
			}
		}
	}

}
