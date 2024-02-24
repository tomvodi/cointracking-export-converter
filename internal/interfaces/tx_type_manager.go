package interfaces

import (
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	bp "github.com/tomvodi/cointracking-export-converter/internal/common/blockpit_tx_type"
	ct "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
)

//go:generate mockgen -source tx_type_manager.go -destination ./mocks/mock_tx_type_manager.go

type TxTypeManager interface {
	// GetMapping returns the list of mappings between CoinTracking and Blockpit
	GetMapping() ([]common.Ct2BpTxMapping, error)
	// SetMapping sets a mapping for a CoinTracking transaction type to a Blockpit transaction type
	SetMapping(ctTxType ct.CtTxType, bpTxType bp.BpTxType) error
	// BlockpitTxTypes returns all Blockpit transaction types
	BlockpitTxTypes() ([]common.TxDisplayName, error)
	// BlockpitTxType returns the mapped Blockpit type to a CoinTracking transaction type
	BlockpitTxType(ctTxType ct.CtTxType) (common.TxDisplayName, error)
}
