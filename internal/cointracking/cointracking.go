package cointracking

import (
	"cointracking-export-converter/internal/interfaces"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ct struct {
	appCtx interfaces.AppContext
}

func (c *ct) OpenExportFile() (string, error) {
	filename, err := runtime.OpenFileDialog(c.appCtx.Context(), runtime.OpenDialogOptions{
		DefaultDirectory: c.appCtx.LastSelectedFileDir(),
		Title:            "Select Cointracking export file",
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

	c.appCtx.SetLastSelectedFileDirFromFile(filename)

	return filename, nil
}

func New(appCtx interfaces.AppContext) interfaces.CoinTrackingBackend {
	return &ct{
		appCtx: appCtx,
	}
}
