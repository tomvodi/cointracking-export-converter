package xml

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type File struct {
	file *excelize.File
}

func (x *File) Close() error {
	return x.file.Close()
}

func (x *File) SetSheetHeader(sheetNr int, headers []string) error {
	sheetName := fmt.Sprintf("Sheet%d", sheetNr)
	return x.file.SetSheetRow(sheetName, "A1", &headers)
}

func (x *File) SetSheetRow(sheetNr, rowNr int, data []any) error {
	sheetName := fmt.Sprintf("Sheet%d", sheetNr)
	cell := fmt.Sprintf("A%d", rowNr)
	return x.file.SetSheetRow(sheetName, cell, &data)
}

func (x *File) SaveAs(filePath string) error {
	return x.file.SaveAs(filePath)
}

func newXMLFile() *File {
	return &File{
		file: excelize.NewFile(),
	}
}
