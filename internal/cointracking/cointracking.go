package cointracking

import (
	"cointracking-export-converter/internal/common"
	bp_type "cointracking-export-converter/internal/common/blockpit_tx_type"
	ct_type "cointracking-export-converter/internal/common/cointracking_tx_type"
	"cointracking-export-converter/internal/interfaces"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

type ct struct {
	appCtx        interfaces.AppContext
	csvReader     interfaces.CointrackingCsvReader
	exportFiles   []*common.ExportFileInfo
	txTypeManager interfaces.TxTypeManager
}

func (c *ct) BlockpitTxTypes() ([]common.TxDisplayName, error) {
	return c.txTypeManager.BlockpitTxTypes()
}

func (c *ct) TxTypeMappings() ([]common.Ct2BpTxMapping, error) {
	return c.txTypeManager.GetMapping()
}

func (c *ct) SetCointracking2BlockpitMapping(
	ctTxType string,
	bpTxType string,
) error {
	ctType, err := ct_type.CtTxTypeString(ctTxType)
	if err != nil {
		return fmt.Errorf("cointracking tx type %s is no valid type", ctTxType)
	}
	bpType, err := bp_type.BpTxTypeString(bpTxType)
	if err != nil {
		return fmt.Errorf("blockpit tx type %s is no valid type", bpTxType)
	}

	runtime.LogTracef(c.appCtx.Context(), "set cointracking tx mapping for '%s' to Blockpit Tx type '%s'",
		ctTxType, bpTxType)
	return c.txTypeManager.SetMapping(ctType, bpType)
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

	for _, file := range c.exportFiles {
		if file.FilePath == filename {
			return "", fmt.Errorf("file already added")
		}
	}

	c.appCtx.SetLastSelectedFileDirFromFile(filename)

	fileInfo, err := c.csvReader.ReadFile(filename, loc)
	if err != nil {
		return "", fmt.Errorf("failed reading file %s: %s", filename, err.Error())
	}

	c.exportFiles = append(c.exportFiles, fileInfo)

	runtime.EventsEmit(c.appCtx.Context(), "ExportFilesChanged", c.exportFiles)

	return filename, nil
}

func New(
	appCtx interfaces.AppContext,
	csvReader interfaces.CointrackingCsvReader,
	txTypeManager interfaces.TxTypeManager,
) interfaces.CoinTrackingBackend {
	return &ct{
		appCtx:        appCtx,
		csvReader:     csvReader,
		txTypeManager: txTypeManager,
	}
}
