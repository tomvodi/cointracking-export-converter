package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	bp_type "github.com/tomvodi/cointracking-export-converter/internal/common/blockpit_tx_type"
	ct_type "github.com/tomvodi/cointracking-export-converter/internal/common/cointracking_tx_type"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
)

type appConfig struct {
	wailsLog      interfaces.WailsLog
	appCtx        interfaces.AppContext
	txTypeManager interfaces.TxTypeManager
}

func (a *appConfig) SetTimezone(tz string) error {
	viper.Set("timezone", tz)
	return a.writeConfig()
}

func (a *appConfig) Timezone() string {
	tz := viper.GetString("timezone")

	return tz
}

func (a *appConfig) AllTimezones() []common.TimezoneData {
	return AllTimezones
}

func (a *appConfig) BlockpitTxTypes() ([]common.TxDisplayName, error) {
	return a.txTypeManager.BlockpitTxTypes()
}

func (a *appConfig) TxTypeMappings() ([]common.Ct2BpTxMapping, error) {
	return a.txTypeManager.GetMapping()
}

func (a *appConfig) SetCointracking2BlockpitMapping(
	ctTxType string,
	bpTxType string,
) error {
	ctType, err := ct_type.CtTxTypeString(ctTxType)
	if err != nil {
		return fmt.Errorf("cointracking tx type %s is no valid type", ctTxType)
	}
	bpType, err := bp_type.BpTxTypeString(bpTxType)
	if err != nil {
		return fmt.Errorf("blockpit tx type %s is no valid type", bpTxType)
	}

	a.wailsLog.LogTracef(a.appCtx.Context(), "set cointracking tx mapping for '%s' to Blockpit Tx type '%s'",
		ctTxType, bpTxType)
	return a.txTypeManager.SetMapping(ctType, bpType)
}

func (a *appConfig) writeConfig() error {
	return viper.WriteConfig()
}

func NewAppConfig(
	appCtx interfaces.AppContext,
	txTypeManager interfaces.TxTypeManager,
	wailsLog interfaces.WailsLog,
) interfaces.AppConfig {
	return &appConfig{
		wailsLog:      wailsLog,
		appCtx:        appCtx,
		txTypeManager: txTypeManager,
	}
}
