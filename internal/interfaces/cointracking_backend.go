package interfaces

import "cointracking-export-converter/internal/common"

//go:generate mockgen -source CoinTrackingBackend.go -destination ./mocks/mock_CointrackingBackend.go

type CoinTrackingBackend interface {
	OpenExportFile(timezone string) (string, error)
	TxTypeMappings() ([]common.Ct2BpTxMapping, error)
	BlockpitTxTypes() ([]common.TxDisplayName, error)
}
