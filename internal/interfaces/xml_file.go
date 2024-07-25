package interfaces

import "io"

type XmlFile interface {
	io.Closer
	SetSheetHeader(sheetNr int, headers []string) error
	SetSheetRow(sheetNr int, rowNr int, data []interface{}) error
	SaveAs(filePath string) error
}
