package blockpit

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
)

var XMLHeaderLabels = []string{
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

type TxXMLFileWriter struct {
	xmlFFactory interfaces.XMLFileFactory
	txConverter interfaces.BlockpitTxConverter
}

func (t *TxXMLFileWriter) WriteTransactionsToFile(
	filePath string,
	transactions []*common.CointrackingTx,
) error {
	f := t.xmlFFactory.NewXMLFile()
	defer f.Close()

	bpTxs, err := t.BpTxsFromCtTxs(transactions)
	if err != nil {
		return err
	}

	err = f.SetSheetHeader(1, XMLHeaderLabels)
	if err != nil {
		return err
	}

	rowNr := 2
	for _, bpTx := range bpTxs {
		err := setExcelSheetRowForTxs(f, rowNr, bpTx)
		if err != nil {
			return err
		}

		rowNr++
	}

	if err = f.SaveAs(filePath); err != nil {
		return err
	}

	return nil
}

func (t *TxXMLFileWriter) BpTxsFromCtTxs(
	transactions []*common.CointrackingTx,
) ([]*interfaces.BlockpitTx, error) {
	var bpTxs []*interfaces.BlockpitTx

	for _, tx := range transactions {
		bpTx, err := t.txConverter.FromCointrackingTx(tx)
		if err != nil {
			return nil, err
		}

		bpTxs = append(bpTxs, bpTx...)
	}

	return bpTxs, nil
}

func setExcelSheetRowForTxs(
	f interfaces.XMLFile,
	rowNr int,
	bpTx *interfaces.BlockpitTx,
) error {
	ctTx := bpTx.CtTx
	excelDate := ctTx.DateTime.Time.UTC().Format("02.01.2006 15:04:05")

	excelRow := []any{
		excelDate,
		ctTx.Exchange,
		bpTx.TxType.Title,
		ctTx.SellCurrency,
		ctTx.SellValue,
		ctTx.BuyCurrency,
		ctTx.BuyValue,
		ctTx.FeeCurrency,
		ctTx.FeeValue,
		ctTx.Comment,
		ctTx.ID,
	}

	err := f.SetSheetRow(1, rowNr, excelRow)
	if err != nil {
		return fmt.Errorf("failed settings row in excel sheet: %s", err.Error())
	}

	return nil
}

func NewTxXMLFileWriter(
	xmlFactory interfaces.XMLFileFactory,
	txConverter interfaces.BlockpitTxConverter,
) *TxXMLFileWriter {
	return &TxXMLFileWriter{
		xmlFFactory: xmlFactory,
		txConverter: txConverter,
	}
}
