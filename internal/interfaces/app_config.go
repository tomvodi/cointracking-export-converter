package interfaces

import (
	"github.com/tomvodi/cointracking-export-converter/internal/common"
)

type AppConfig interface {
	SetTimezone(tz string) error
	Timezone() string
	TxTypeMappings() ([]common.Ct2BpTxMapping, error)
	BlockpitTxTypes() ([]common.TxDisplayName, error)
	SetCointracking2BlockpitMapping(ctTxType string, bpTxType string) error
	SwapHandling() string
	SetSwapHandling(handling string) error
}
