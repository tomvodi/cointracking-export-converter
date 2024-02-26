package blockpit

import (
	"github.com/tomvodi/cointracking-export-converter/internal/common"
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
	return b.blockpitXmlWriter.WriteTransactionsToXmlFile(filename, allTxs)
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
