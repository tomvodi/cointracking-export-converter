package app

import (
	"context"
	"github.com/spf13/afero"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"path/filepath"
)

type Ctx struct {
	afs                 afero.Fs
	ctx                 context.Context
	lastSelectedFileDir string
	exportFiles         []*common.ExportFileInfo
	txIDs               []string
}

func (a *Ctx) AllTxIDs() []string {
	return a.txIDs
}

func (a *Ctx) ExportFiles() []*common.ExportFileInfo {
	return a.exportFiles
}

// AddExportFile adds a new export file to the context.
// It also adds all transaction ids from the file to the list of known transaction ids.
func (a *Ctx) AddExportFile(file *common.ExportFileInfo) {
	for _, transaction := range file.Transactions {
		if !a.containsTxID(transaction.ID) {
			a.txIDs = append(a.txIDs, transaction.ID)
		}
	}

	a.exportFiles = append(a.exportFiles, file)
}

func (a *Ctx) containsTxID(txID string) bool {
	for _, id := range a.txIDs {
		if id == txID {
			return true
		}
	}

	return false
}

func (a *Ctx) SetLastSelectedFileDirFromFile(file string) {
	fileInfo, err := a.afs.Stat(file)
	if err != nil {
		return
	}

	if !fileInfo.IsDir() {
		a.lastSelectedFileDir = filepath.Dir(file)
	}
}

func (a *Ctx) LastSelectedFileDir() string {
	return a.lastSelectedFileDir
}

func (a *Ctx) Context() context.Context {
	return a.ctx
}

func (a *Ctx) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func NewAppContext(afs afero.Fs) *Ctx {
	return &Ctx{
		afs:         afs,
		exportFiles: make([]*common.ExportFileInfo, 0),
	}
}
