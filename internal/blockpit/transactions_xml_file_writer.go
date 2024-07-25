package blockpit

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
)

var header = []string{
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

type txXmlFWriter struct {
	xmlFFactory interfaces.XmlFileFactory
	txConverter interfaces.BlockpitTxConverter
}

func (t *txXmlFWriter) WriteTransactionsToFile(
	filePath string,
	transactions []*common.CointrackingTx,
) error {
	f := t.xmlFFactory.NewXmlFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	err := f.SetSheetHeader(1, header)
	if err != nil {
		return err
	}

	var bpTxs []*interfaces.BlockpitTx

	for _, tx := range transactions {
		bpTx, err := t.txConverter.FromCointrackingTx(tx)
		if err != nil {
			return err
		}

		bpTxs = append(bpTxs, bpTx...)
	}

	rowNr := 2
	for _, bpTx := range bpTxs {
		ctTx := bpTx.CtTx
		excelDate := ctTx.DateTime.Time.UTC().Format("02.01.2006 15:04:05")

		excelRow := []interface{}{
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

		err = f.SetSheetRow(1, rowNr, excelRow)
		if err != nil {
			return fmt.Errorf("failed settings row in excel sheet: %s", err.Error())
		}

		rowNr++
	}

	if err = f.SaveAs(filePath); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewTxXmlFileWriter(
	xmlFactory interfaces.XmlFileFactory,
	txConverter interfaces.BlockpitTxConverter,
) interfaces.TransactionsFileWriter {
	return &txXmlFWriter{
		xmlFFactory: xmlFactory,
		txConverter: txConverter,
	}
}
