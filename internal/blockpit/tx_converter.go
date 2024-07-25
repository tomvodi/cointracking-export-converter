package blockpit

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"github.com/tomvodi/cointracking-export-converter/internal/common/swap_handling"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
)

type txConvert struct {
	appCfg    interfaces.AppConfig
	txTypeMgr interfaces.TxTypeManager
}

func (t *txConvert) FromCointrackingTx(
	ctTx *common.CointrackingTx,
) ([]*interfaces.BlockpitTx, error) {
	txs := make([]*interfaces.BlockpitTx, 0)

	displayType, err := t.txTypeMgr.BlockpitTxType(ctTx.Type.TxType)
	if err != nil {
		return nil, err
	}

	bpTx := &interfaces.BlockpitTx{
		TxType: displayType,
		CtTx:   &common.CointrackingTx{},
	}

	if ctTx.Type.TxType == ctt.SwapNonTaxable {
		swpHandling, err := swap_handling.SwapHandlingString(t.appCfg.SwapHandling())
		if err != nil {
			return nil, fmt.Errorf("failed to convert swap handling: %s", err.Error())
		}

		if swpHandling == swap_handling.SwapToTrade {
			ctTx.Type.TxType = ctt.Trade
			bpTx.CtTx = ctTx
			txs = append(txs, bpTx)
		} else {
			txOut := *ctTx
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
			bpTx.CtTx = &txOut
			txs = append(txs, bpTx)

			txIn := *ctTx
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

			// Create new object for bpTx as CtTx member is a pointer and
			// we don't want to change the original object
			bpTx = &interfaces.BlockpitTx{
				TxType: displayType,
				CtTx:   &common.CointrackingTx{},
			}
			bpTx.CtTx = &txIn
			txs = append(txs, bpTx)
		}
	} else {
		bpTx.CtTx = ctTx
		txs = append(txs, bpTx)
	}

	return txs, nil
}

func NewTxConvert(
	appCfg interfaces.AppConfig,
	txTypeMgr interfaces.TxTypeManager,
) interfaces.BlockpitTxConverter {
	return &txConvert{
		appCfg:    appCfg,
		txTypeMgr: txTypeMgr,
	}
}
