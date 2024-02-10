package config

import (
	"cointracking-export-converter/internal/interfaces"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type AppConfigInitializer interface {
	interfaces.AppConfig
	interfaces.Initializer
}

var configFileName = "config"

type appConfig struct {
	configDir string
	appCtx    interfaces.AppContext
}

func (a *appConfig) Init() error {
	viper.AddConfigPath(a.configDir)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		} else {
			return err
		}
	}

	err := os.MkdirAll(a.configDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed creating config dir: %s", err.Error())
	}

	return nil
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

	return viper.GetString("timezone")
}

func (a *appConfig) writeConfig() error {
	configPath := filepath.Join(a.configDir, configFileName+".yaml")
	_, err := os.Stat(configPath)
	if !os.IsExist(err) {
		if _, err := os.Create(configPath); err != nil {
			return fmt.Errorf("failed creating config file: %s", err.Error())
		}
	}

	return viper.WriteConfig()
}

func NewAppConfig(configDir string, appCtx interfaces.AppContext) AppConfigInitializer {
	return &appConfig{
		configDir: configDir,
		appCtx:    appCtx,
	}
}
