package interfaces

import (
	"context"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
)

type AppContext interface {
	// Context returns the application context used for wails runtime calls
	// or nil if wasn't set yet
	Context() context.Context
	SetContext(ctx context.Context)

	SetLastSelectedFileDirFromFile(filepath string)
	LastSelectedFileDir() string

	ExportFiles() []*common.ExportFileInfo
	AddExportFile(file *common.ExportFileInfo)
}
