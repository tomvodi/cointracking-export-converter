package interfaces

import "io"

type XMLFile interface {
	io.Closer
	SetSheetHeader(sheetNr int, headers []string) error
	SetSheetRow(sheetNr int, rowNr int, data []any) error
	SaveAs(filePath string) error
}
