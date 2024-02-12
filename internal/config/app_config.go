package config

import (
	"cointracking-export-converter/internal/common"
	"cointracking-export-converter/internal/interfaces"
	"github.com/spf13/viper"
)

type appConfig struct {
	appCtx interfaces.AppContext
}

func (a *appConfig) SetTimezone(tz string) error {
	viper.Set("timezone", tz)
	return a.writeConfig()
}

func (a *appConfig) Timezone() string {
	// Initial timezone from local machine
	//locTime := time.Now()
	//zone, offset := locTime.Zone()
	//fmt.Println(zone, offset)
	tz := viper.GetString("timezone")

	return tz
}

func (a *appConfig) AllTimezones() []common.TimezoneData {
	return AllTimezones
}

func (a *appConfig) writeConfig() error {
	return viper.WriteConfig()
}

func NewAppConfig(appCtx interfaces.AppContext) interfaces.AppConfig {
	return &appConfig{
		appCtx: appCtx,
	}
}
