package app

import (
	"context"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"os"
	"path/filepath"
)

type appCtx struct {
	ctx                 context.Context
	lastSelectedFileDir string
	exportFiles         []*common.ExportFileInfo
	txIds               []string
}

func (a *appCtx) AllTxIds() []string {
	return a.txIds
}

func (a *appCtx) ExportFiles() []*common.ExportFileInfo {
	return a.exportFiles
}

func (a *appCtx) AddExportFile(file *common.ExportFileInfo) {
	for _, transaction := range file.Transactions {
		// Skip transaction ids that have already been added
		containsId := false
		for _, id := range a.txIds {
			if id == transaction.ID {
				containsId = true
				break
			}
		}

		if !containsId {
			a.txIds = append(a.txIds, transaction.ID)
		}
	}

	a.exportFiles = append(a.exportFiles, file)
}

func (a *appCtx) SetLastSelectedFileDirFromFile(file string) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return
	}

	if !fileInfo.IsDir() {
		a.lastSelectedFileDir = filepath.Dir(file)
	}
}

func (a *appCtx) LastSelectedFileDir() string {
	return a.lastSelectedFileDir
}

func (a *appCtx) Context() context.Context {
	return a.ctx
}

func (a *appCtx) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func NewAppContext() interfaces.AppContext {
	return &appCtx{
		exportFiles: make([]*common.ExportFileInfo, 0),
	}
}
