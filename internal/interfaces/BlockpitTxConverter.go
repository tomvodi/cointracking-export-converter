package interfaces

import "github.com/tomvodi/cointracking-export-converter/internal/common"

type BlockpitTx struct {
	TxType common.TxDisplayName
	CtTx   *common.CointrackingTx
}

type BlockpitTxConverter interface {
	// FromCointrackingTx converts a CointrackingTx to a BlockpitTx
	// it may return multiple BlockpitTx if the conversion requires it depending
	// on the swap handling setting for example
	FromCointrackingTx(ctTx *common.CointrackingTx) ([]*BlockpitTx, error)
}
