package app

import (
	"cointracking-export-converter/internal/interfaces"
	"context"
	"os"
	"path/filepath"
)

type appCtx struct {
	ctx                 context.Context
	lastSelectedFileDir string
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
	return &appCtx{}
}
