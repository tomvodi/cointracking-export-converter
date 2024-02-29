package blockpit

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type bp struct {
	appCtx            interfaces.AppContext
	blockpitXmlWriter interfaces.XmlWriter
}

func (b *bp) ExportToBlockpitXlsx() error {
	filename, err := runtime.SaveFileDialog(b.appCtx.Context(), runtime.SaveDialogOptions{
		DefaultDirectory: b.appCtx.LastSelectedFileDir(),
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

	var allTxs []*common.CointrackingTx
	for _, info := range b.appCtx.ExportFiles() {
		allTxs = append(allTxs, info.Transactions...)
	}

	for i, tx := range allTxs {
		allTxs[i], err = getBlockpitTxFeeAdapted(tx)
		adaptTxTypeForTradesWith0Income(allTxs[i])
		if err != nil {
			return err
		}
	}

	return b.blockpitXmlWriter.WriteTransactionsToXmlFile(filename, allTxs)
}

// getBlockpitTxFeeAdapted returns a new CointrackingTx with the fee value adapted to Blockpit
// as CoinTracking does handle fee values in a transaction differently than Blockpit.
func getBlockpitTxFeeAdapted(ctTx *common.CointrackingTx) (*common.CointrackingTx, error) {
	bpTx := *ctTx
	if ctTx.FeeCurrency == ctTx.SellCurrency {
		bpTx.SellValue -= bpTx.FeeValue
		if bpTx.SellValue < 0 {
			return nil, fmt.Errorf(
				"fee value %.2f is higher than sell value %.2f",
				bpTx.FeeValue,
				bpTx.SellValue,
			)
		}
	}
	if ctTx.FeeCurrency == ctTx.BuyCurrency {
		bpTx.BuyValue += bpTx.FeeValue
		return &bpTx, nil
	}

	return &bpTx, nil
}

func adaptTxTypeForTradesWith0Income(ctTx *common.CointrackingTx) {
	if ctTx.Type.TxType != ctt.Trade {
		return
	}

	if ctTx.BuyValue == 0 {
		ctTx.Type.TxType = ctt.OtherExpense
	}

	if ctTx.SellValue == 0 {
		ctTx.Type.TxType = ctt.OtherIncome
	}
}

func New(
	appCtx interfaces.AppContext,
	blockpitXmlWriter interfaces.XmlWriter,
) interfaces.BlockpitBackend {
	return &bp{
		appCtx:            appCtx,
		blockpitXmlWriter: blockpitXmlWriter,
	}
}
