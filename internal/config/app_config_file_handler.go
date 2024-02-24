package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"io/fs"
	"os"
	"path/filepath"
)

var configFileName = "config"

type fileHandler struct {
	configDir string
	fs        afero.Fs
}

func (f *fileHandler) Init() error {
	err := f.fs.MkdirAll(f.configDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed creating config dir: %s", err.Error())
	}

	configPath := filepath.Join(f.configDir, configFileName+".yaml")
	_, err = f.fs.Stat(configPath)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		if _, err := f.fs.Create(configPath); err != nil {
			return fmt.Errorf("failed creating config file: %s", err.Error())
		}
	}

	viper.SetFs(f.fs)
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

	return nil
}

func NewConfigFileHandler(fs afero.Fs, configDir string) interfaces.Initializer {
	return &fileHandler{
		configDir: configDir,
		fs:        fs,
	}
}
