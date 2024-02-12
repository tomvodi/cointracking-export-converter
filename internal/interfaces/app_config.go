package interfaces

import "cointracking-export-converter/internal/common"

//go:generate mockgen -source app_config.go -destination ./mocks/mock_app_config.go

type AppConfig interface {
	SetTimezone(tz string) error
	Timezone() string
	TxTypeMappings() ([]common.Ct2BpTxMapping, error)
	BlockpitTxTypes() ([]common.TxDisplayName, error)
	SetCointracking2BlockpitMapping(ctTxType string, bpTxType string) error
}
