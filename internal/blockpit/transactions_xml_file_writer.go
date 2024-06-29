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
	txTypeManager interfaces.TxTypeManager
	appCfg        interfaces.AppConfig
	xmlFFactory   interfaces.XmlFileFactory
	txConverter   interfaces.BlockpitTxConverter
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

	rowNr := 2
	for _, tx := range transactions {
		convertedTxs, err := t.txConverter.FromCointrackingTx(tx)
		if err != nil {
			return err
		}

		for _, tx = range convertedTxs {
			excelDate := tx.DateTime.Time.UTC().Format("02.01.2006 15:04:05")

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

			err = f.SetSheetRow(1, rowNr, excelRow)
			if err != nil {
				return fmt.Errorf("failed settings row in excel sheet: %s", err.Error())
			}

			rowNr++
		}
	}

	if err = f.SaveAs(filePath); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewTxXmlFileWriter(

	xmlFactory interfaces.XmlFileFactory,
) interfaces.TransactionsFileWriter {
	return &txXmlFWriter{
		txTypeManager: txTypeMgr,
		appCfg:        appCfg,
		xmlFFactory:   xmlFactory,
	}
}
