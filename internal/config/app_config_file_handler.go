package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path/filepath"
)

var configFileName = "config"

type FileHandler struct {
	configDir string
	afs       afero.Fs
}

func (f *FileHandler) Init() error {
	err := f.initConfigDir()
	if err != nil {
		return err
	}

	viper.SetFs(f.afs)
	viper.AddConfigPath(f.configDir)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil
		}
		return err
	}

	return nil
}

func (f *FileHandler) initConfigDir() error {
	err := f.afs.MkdirAll(f.configDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed creating config dir: %s", err.Error())
	}

	configPath := filepath.Join(f.configDir, configFileName+".yaml")
	_, err = f.afs.Stat(configPath)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		if _, err := f.afs.Create(configPath); err != nil {
			return fmt.Errorf("failed creating config file: %s", err.Error())
		}
	}

	return nil
}

func NewConfigFileHandler(
	afs afero.Fs,
	configDir string,
) *FileHandler {
	return &FileHandler{
		configDir: configDir,
		afs:       afs,
	}
}
