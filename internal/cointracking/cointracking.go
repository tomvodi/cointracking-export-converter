package cointracking

import (
	"cointracking-export-converter/internal/common"
	bp "cointracking-export-converter/internal/common/blockpit_tx_type"
	ct_type "cointracking-export-converter/internal/common/cointracking_tx_type"
	"cointracking-export-converter/internal/interfaces"
	"cointracking-export-converter/internal/localization"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

type ct struct {
	appCtx      interfaces.AppContext
	csvReader   interfaces.CointrackingCsvReader
	exportFiles []*common.ExportFileInfo
}

func (c *ct) BlockpitTxTypes() (txNames []common.TxDisplayName, err error) {
	for _, txType := range bp.BpTxTypeValues() {
		if txType == bp.NoBpTxType {
			continue
		}

		txName := common.TxDisplayName{
			Value: txType.String(),
		}

		if translation, found := localization.BpTxTypeNames[txType]; found {
			txName.Title = translation
		} else {
			return nil,
				fmt.Errorf("no localization for Blockpit tx type %s found", txName.Value)
		}
		txNames = append(txNames, txName)
	}

	return txNames, nil
}

func (c *ct) TxTypeMappings() (mapping []common.Ct2BpTxMapping, err error) {
	for _, txType := range ct_type.CtTxTypeValues() {
		if txType == ct_type.NoCtTxType {
			continue
		}

		mapItem := common.Ct2BpTxMapping{
			Cointracking: common.TxDisplayName{
				Value: txType.String(),
			},
		}

		if translation, found := localization.CtTxTypeNames[txType]; found {
			mapItem.Cointracking.Title = translation
		} else {
			return nil,
				fmt.Errorf("no localization for CoinTracking tx type %s found", txType.String())
		}

		var bpType bp.BpTxType
		var found bool
		if bpType, found = common.Ct2BpMap[txType]; found {
			mapItem.Blockpit.Value = bpType.String()
		} else {
			return nil,
				fmt.Errorf("no blockpit tx type for CoinTracking tx type %s found", txType.String())
		}

		if translation, found := localization.BpTxTypeNames[bpType]; found {
			mapItem.Blockpit.Title = translation
		} else {
			return nil,
				fmt.Errorf("no localization for Blockpit tx type %s found", mapItem.Blockpit.Value)
		}

		mapping = append(mapping, mapItem)
	}

	return mapping, nil
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
) interfaces.CoinTrackingBackend {
	return &ct{
		appCtx:    appCtx,
		csvReader: csvReader,
	}
}
