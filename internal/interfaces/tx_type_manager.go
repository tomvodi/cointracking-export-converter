package interfaces

import (
	"cointracking-export-converter/internal/common"
	bp "cointracking-export-converter/internal/common/blockpit_tx_type"
	ct "cointracking-export-converter/internal/common/cointracking_tx_type"
)

//go:generate mockgen -source tx_type_manager.go -destination ./mocks/mock_tx_type_manager.go

type TxTypeManager interface {
	GetMapping() ([]common.Ct2BpTxMapping, error)
	SetMapping(ctTxType ct.CtTxType, bpTxType bp.BpTxType) error
	BlockpitTxTypes() ([]common.TxDisplayName, error)
	BlockpitTxType(ctTxType ct.CtTxType) (common.TxDisplayName, error)
}
