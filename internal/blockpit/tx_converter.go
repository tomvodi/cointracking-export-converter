package blockpit

import (
	"fmt"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	ctt "github.com/tomvodi/cointracking-export-converter/internal/common/cointrackingtxtype"
	"github.com/tomvodi/cointracking-export-converter/internal/common/swaphandling"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
)

type TxConverter struct {
	appCfg    interfaces.AppConfig
	txTypeMgr interfaces.TxTypeManager
}

func (t *TxConverter) FromCointrackingTx(
	ctTx *common.CointrackingTx,
) ([]*interfaces.BlockpitTx, error) {
	bpTx, err := t.newBlockpitTxWithType(ctTx.Type.TxType)
	if err != nil {
		return nil, err
	}

	if ctTx.Type.TxType != ctt.SwapNonTaxable {
		bpTx.CtTx = ctTx
		return []*interfaces.BlockpitTx{bpTx}, nil
	}

	// Swap handling non taxable
	swpHandling, err := swaphandling.SwapHandlingString(t.appCfg.SwapHandling())
	if err != nil {
		return nil, fmt.Errorf("failed to convert swap handling: %s", err.Error())
	}

	switch swpHandling {
	case swaphandling.SwapToTrade:
		ctTx.Type.TxType = ctt.Trade
		bpTx.CtTx = ctTx
		return []*interfaces.BlockpitTx{bpTx}, nil
	case swaphandling.SwapNonTaxable:
		return t.getNonTaxableSwap(ctTx)
	default:
		return nil, fmt.Errorf("unknown swap handling: %s", t.appCfg.SwapHandling())
	}
}

func (t *TxConverter) getNonTaxableSwap(
	ctTx *common.CointrackingTx,
) ([]*interfaces.BlockpitTx, error) {
	bpTx, err := t.newBlockpitTxWithType(ctTx.Type.TxType)
	if err != nil {
		return nil, err
	}

	var txs []*interfaces.BlockpitTx
	txOut := *ctTx
	txOut.Type = &common.TxType{TxType: ctt.ExpenseNonTaxable}
	txOut.BuyValue = 0.0
	txOut.BuyCurrency = ""
	if txOut.FeeCurrency != txOut.SellCurrency {
		txOut.FeeValue = 0.0
		txOut.FeeCurrency = ""
	}
	// When splitting up one transaction into two, the new transactions get a new ID
	err = common.SetIDForTransaction(&txOut)
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
	err = common.SetIDForTransaction(&txIn)
	if err != nil {
		return txs, fmt.Errorf("failed to set id for transaction: %s", err.Error())
	}

	// Create new object for bpTx as CtTx member is a pointer and
	// we don't want to change the original object
	bpTx, err = t.newBlockpitTxWithType(ctTx.Type.TxType)
	if err != nil {
		return nil, err
	}

	bpTx.CtTx = &txIn
	txs = append(txs, bpTx)

	return txs, nil
}

func (t *TxConverter) newBlockpitTxWithType(
	ctTxType ctt.CtTxType,
) (*interfaces.BlockpitTx, error) {
	displayType, err := t.txTypeMgr.BlockpitTxType(ctTxType)
	if err != nil {
		return nil, err
	}

	return &interfaces.BlockpitTx{
		TxType: displayType,
		CtTx:   &common.CointrackingTx{},
	}, nil
}

func NewTxConvert(
	appCfg interfaces.AppConfig,
	txTypeMgr interfaces.TxTypeManager,
) *TxConverter {
	return &TxConverter{
		appCfg:    appCfg,
		txTypeMgr: txTypeMgr,
	}
}
