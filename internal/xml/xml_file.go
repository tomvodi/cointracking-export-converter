package xml

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/xuri/excelize/v2"
)

type xmlFile struct {
	file *excelize.File
}

func (x *xmlFile) Close() error {
	return x.file.Close()
}

func (x *xmlFile) SetSheetHeader(sheetNr int, headers []string) error {
	sheetName := fmt.Sprintf("Sheet%d", sheetNr)
	return x.file.SetSheetRow(sheetName, "A1", &headers)
}

func (x *xmlFile) SetSheetRow(sheetNr, rowNr int, data []interface{}) error {
	sheetName := fmt.Sprintf("Sheet%d", sheetNr)
	cell := fmt.Sprintf("A%d", rowNr)
	return x.file.SetSheetRow(sheetName, cell, &data)
}

func (x *xmlFile) SaveAs(filePath string) error {
	return x.file.SaveAs(filePath)
}

func newXmlFile() interfaces.XmlFile {
	return &xmlFile{
		file: excelize.NewFile(),
	}
}
