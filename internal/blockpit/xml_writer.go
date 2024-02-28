package blockpit

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"github.com/tomvodi/cointracking-export-converter/internal/common/swap_handling"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/xuri/excelize/v2"
)

type xmlWriter struct {
	txTypeManager interfaces.TxTypeManager
	appCfg        interfaces.AppConfig
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
		convertedTxs, err := x.convertCtTxToBlockpitTx(tx)
		if err != nil {
			return err
		}

		for _, tx = range convertedTxs {
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
	}

	if err = f.SaveAs(filePath); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (x *xmlWriter) convertCtTxToBlockpitTx(tx *common.CointrackingTx) ([]*common.CointrackingTx, error) {
	txs := make([]*common.CointrackingTx, 0)
	if tx.Type.TxType == ctt.SwapNonTaxable {
		swpHandling, err := swap_handling.SwapHandlingString(x.appCfg.SwapHandling())
		if err != nil {
			return nil, fmt.Errorf("failed to convert swap handling: %s", err.Error())
		}

		if swpHandling == swap_handling.SwapToTrade {
			tx.Type.TxType = ctt.Trade
			txs = append(txs, tx)
		} else {
			txOut := *tx
			txOut.Type = &common.TxType{TxType: ctt.ExpenseNonTaxable}
			txOut.BuyValue = 0.0
			txOut.BuyCurrency = ""
			if txOut.FeeCurrency != txOut.SellCurrency {
				txOut.FeeValue = 0.0
				txOut.FeeCurrency = ""
			}
			// When splitting up one transaction into two, the new transactions get a new ID
			err = common.SetIdForTransaction(&txOut)
			if err != nil {
				return txs, fmt.Errorf("failed to set id for transaction: %s", err.Error())
			}
			txs = append(txs, &txOut)

			txIn := *tx
			txIn.Type = &common.TxType{TxType: ctt.IncomeNonTaxable}
			txIn.SellValue = 0.0
			txIn.SellCurrency = ""
			if txIn.FeeCurrency != txIn.BuyCurrency {
				txIn.FeeValue = 0.0
				txIn.FeeCurrency = ""
			}
			err = common.SetIdForTransaction(&txIn)
			if err != nil {
				return txs, fmt.Errorf("failed to set id for transaction: %s", err.Error())
			}
			txs = append(txs, &txIn)
		}
	} else {
		txs = append(txs, tx)
	}

	return txs, nil
}

func NewXmlWriter(
	txTypeMgr interfaces.TxTypeManager,
	appCfg interfaces.AppConfig,
) interfaces.XmlWriter {
	return &xmlWriter{
		txTypeManager: txTypeMgr,
		appCfg:        appCfg,
	}
}
