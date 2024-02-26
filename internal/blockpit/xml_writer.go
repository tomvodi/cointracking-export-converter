package blockpit

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/xuri/excelize/v2"
)

type xmlWriter struct {
	txTypeManager interfaces.TxTypeManager
}

func (x *xmlWriter) WriteTransactionsToXmlFile(
	filePath string,
	transactions []*common.CointrackingTx,
) error {
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

	err := f.SetSheetRow("Sheet1", "A1", &header)
	if err != nil {
		return err
	}

	rowNr := 2
	for _, tx := range transactions {
		excelDate := tx.DateTime.Time.UTC().Format("02.01.2006 15:04:05")
		excelTxType, err := x.txTypeManager.BlockpitTxType(tx.Type.TxType)
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

	if err = f.SaveAs(filePath); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewXmlWriter(
	txTypeMgr interfaces.TxTypeManager,
) interfaces.XmlWriter {
	return &xmlWriter{
		txTypeManager: txTypeMgr,
	}
}
