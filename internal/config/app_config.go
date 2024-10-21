package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	bptype "github.com/tomvodi/cointracking-export-converter/internal/common/blockpittxtype"
	cttype "github.com/tomvodi/cointracking-export-converter/internal/common/cointrackingtxtype"
	"github.com/tomvodi/cointracking-export-converter/internal/common/swaphandling"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
)

type AppConfig struct {
	wailsLog      interfaces.WailsLogger
	appCtx        interfaces.AppContext
	txTypeManager interfaces.TxTypeManager
}

func (a *AppConfig) SetTimezone(tz string) error {
	viper.Set("timezone", tz)
	return a.writeConfig()
}

func (a *AppConfig) Timezone() string {
	tz := viper.GetString("timezone")

	return tz
}

func (a *AppConfig) AllTimezones() []common.TimezoneData {
	return AllTimezones
}

func (a *AppConfig) BlockpitTxTypes() ([]common.TxDisplayName, error) {
	return a.txTypeManager.BlockpitTxTypes()
}

func (a *AppConfig) TxTypeMappings() ([]common.Ct2BpTxMapping, error) {
	return a.txTypeManager.GetMapping()
}

func (a *AppConfig) SwapHandling() string {
	viper.SetDefault("swap_handling", swaphandling.SwapNonTaxable)
	handling, err := swaphandling.SwapHandlingString(viper.GetString("swap_handling"))
	if err != nil {
		return swaphandling.NoSwapHandling.String()
	}

	return handling.String()
}

func (a *AppConfig) SetSwapHandling(handling string) error {
	swapHandling, err := swaphandling.SwapHandlingString(handling)
	if err != nil {
		return fmt.Errorf("swap handling %s is no valid handling", handling)
	}
	viper.Set("swap_handling", swapHandling)
	return a.writeConfig()
}

func (a *AppConfig) SetCointracking2BlockpitMapping(
	ctTxType string,
	bpTxType string,
) error {
	ctType, err := cttype.CtTxTypeString(ctTxType)
	if err != nil {
		return fmt.Errorf("cointracking tx type %s is no valid type", ctTxType)
	}
	bpType, err := bptype.BpTxTypeString(bpTxType)
	if err != nil {
		return fmt.Errorf("blockpit tx type %s is no valid type", bpTxType)
	}

	a.wailsLog.LogTracef(a.appCtx.Context(), "set cointracking tx mapping for '%s' to Blockpit Tx type '%s'",
		ctTxType, bpTxType)
	return a.txTypeManager.SetMapping(ctType, bpType)
}

func (a *AppConfig) writeConfig() error {
	return viper.WriteConfig()
}

func NewAppConfig(
	appCtx interfaces.AppContext,
	txTypeManager interfaces.TxTypeManager,
	wailsLog interfaces.WailsLogger,
) *AppConfig {
	return &AppConfig{
		wailsLog:      wailsLog,
		appCtx:        appCtx,
		txTypeManager: txTypeManager,
	}
}
