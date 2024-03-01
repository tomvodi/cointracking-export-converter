package cointracking

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

type ct struct {
	appCtx    interfaces.AppContext
	csvReader interfaces.CointrackingCsvReader
}

func (c *ct) GetExportFiles() ([]*common.ExportFileInfo, error) {
	return c.appCtx.ExportFiles(), nil
}

func (c *ct) OpenExportFile(timezone string) (string, error) {
	filename, err := runtime.OpenFileDialog(c.appCtx.Context(), runtime.OpenDialogOptions{
		DefaultDirectory: c.appCtx.LastSelectedFileDir(),
		Title:            "Select CoinTracking export file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "CoinTracking export files (.csv)",
				Pattern:     "*.csv",
			},
		},
	})
	if err != nil {
		return "", err
	}
	if filename == "" {
		return "", nil
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", fmt.Errorf("failed getting timezone from string %s", timezone)
	}

	for _, file := range c.appCtx.ExportFiles() {
		if file.FilePath == filename {
			return "", fmt.Errorf("file already added")
		}
	}

	c.appCtx.SetLastSelectedFileDirFromFile(filename)
	fileInfo, err := c.csvReader.ReadFile(filename, loc, c.appCtx.AllTxIds())
	if err != nil {
		return "", fmt.Errorf("failed reading file %s: %s", filename, err.Error())
	}

	c.appCtx.AddExportFile(fileInfo)

	runtime.EventsEmit(c.appCtx.Context(), "ExportFilesChanged", c.appCtx.ExportFiles())

	return filename, nil
}

func New(
	appCtx interfaces.AppContext,
	csvReader interfaces.CointrackingCsvReader,
) interfaces.CoinTrackingBackend {
	return &ct{
		appCtx:    appCtx,
		csvReader: csvReader,
	}
}
