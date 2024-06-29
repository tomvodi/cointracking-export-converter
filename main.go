package main

import (
	"embed"
	"github.com/spf13/afero"
	"github.com/tomvodi/cointracking-export-converter/internal/app"
	"github.com/tomvodi/cointracking-export-converter/internal/blockpit"
	"github.com/tomvodi/cointracking-export-converter/internal/cointracking"
	"github.com/tomvodi/cointracking-export-converter/internal/config"
	"github.com/tomvodi/cointracking-export-converter/internal/wails_runtime"
	"github.com/tomvodi/cointracking-export-converter/internal/xml"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var appName = "cointracking-export-converter"

func main() {
	appCtx := app.NewAppContext()
	wailsLog := wails_runtime.NewWailsLog()
	fs := afero.NewOsFs()

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("failed getting user config dir: %s", err.Error())
	}

	appConfigDir := filepath.Join(configDir, appName)
	appConfigFileHandler := config.NewConfigFileHandler(fs, appConfigDir)
	err = appConfigFileHandler.Init()
	if err != nil {
		log.Fatalf("failed initializing config: %s", err.Error())
	}

	txManager := config.NewTxTypeManagerInitializer()
	if err = txManager.Init(); err != nil {
		log.Fatalf("failed initializing TX manager: %s", err.Error())
	}

	appInstance := app.NewApp(appCtx, logger.INFO)
	csvReader := cointracking.NewCsvReader()

	appConfig := config.NewAppConfig(appCtx, txManager, wailsLog)

	xmlFactory := xml.NewXmlFileFactory()
	txConverter := blockpit.NewTxConvert(appConfig, txManager)

	bpXmlWriter := blockpit.NewTxXmlFileWriter(xmlFactory, txConverter)
	bp := blockpit.New(appCtx, bpXmlWriter)
	ct := cointracking.New(appCtx, csvReader)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "CoinTracking Export Converter",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        appInstance.Startup,
		Bind: []interface{}{
			ct,
			bp,
			appConfig,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
