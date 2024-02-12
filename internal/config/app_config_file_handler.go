package config

import (
	"cointracking-export-converter/internal/interfaces"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path/filepath"
)

var configFileName = "config"

type fileHandler struct {
	configDir string
}

func (f *fileHandler) Init() error {
	viper.AddConfigPath(f.configDir)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		} else {
			return err
		}
	}

	err := os.MkdirAll(f.configDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed creating config dir: %s", err.Error())
	}

	configPath := filepath.Join(f.configDir, configFileName+".yaml")
	_, err = os.Stat(configPath)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		if _, err := os.Create(configPath); err != nil {
			return fmt.Errorf("failed creating config file: %s", err.Error())
		}
	}

	return nil
}

func NewConfigFileHandler(configDir string) interfaces.Initializer {
	return &fileHandler{
		configDir: configDir,
	}
}
