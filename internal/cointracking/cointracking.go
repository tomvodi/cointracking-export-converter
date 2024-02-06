package cointracking

import (
	"cointracking-export-converter/internal/common"
	bp_type "cointracking-export-converter/internal/common/blockpit_tx_type"
	ct_type "cointracking-export-converter/internal/common/cointracking_tx_type"
	"cointracking-export-converter/internal/interfaces"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
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

func (c *ct) GetExportFiles() ([]*common.ExportFileInfo, error) {
	return c.exportFiles, nil
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

func (c *ct) ExportToBlockpitXlsx() error {
	filename, err := runtime.SaveFileDialog(c.appCtx.Context(), runtime.SaveDialogOptions{
		DefaultDirectory: c.appCtx.LastSelectedFileDir(),
		DefaultFilename:  "blockpit-import.xlsx",
		Title:            "Save Blockpit manual import file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Blockpit import files (.xlsx)",
				Pattern:     "*.xlsx",
			},
		},
	})
	if err != nil {
		return err
	}
	if filename == "" {
		return nil
	}

	return c.writeTransactionsToXmlFile(filename)
}

func (c *ct) writeTransactionsToXmlFile(filePath string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	header := []string{
		"Date (UTC)",
		"Integration Name",
		"Label",
		"Outgoing Asset",
		"Outgoing Amount",
		"Incoming Asset",
		"Incoming Amount",
		"Fee Asset (optional)",
		"Fee Amount (optional)",
		"Comment (optional)",
		"Trx. ID (optional)",
	}

	// Set value of a cell.
	err := f.SetSheetRow("Sheet1", "A1", &header)
	if err != nil {
		return err
	}

	rowNr := 2
	for _, file := range c.exportFiles {
		for _, tx := range file.Transactions {
			excelDate := tx.DateTime.Time.UTC().Format("02.01.2006 15:04:05")
			excelTxType, err := c.txTypeManager.BlockpitTxType(tx.Type.TxType)
			if err != nil {
				return err
			}

			excelRow := []interface{}{
				excelDate,
				tx.Exchange,
				excelTxType.Title,
				tx.SellCurrency,
				tx.SellValue,
				tx.BuyCurrency,
				tx.BuyValue,
				tx.FeeCurrency,
				tx.FeeValue,
				tx.Comment,
				tx.ID,
			}

			err = f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", rowNr), &excelRow)
			if err != nil {
				return fmt.Errorf("failed settings row in excel sheet: %s", err.Error())
			}

			rowNr++
		}
	}

	// Save spreadsheet by the given path.
	if err = f.SaveAs(filePath); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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
