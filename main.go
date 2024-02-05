package main

import (
	"cointracking-export-converter/internal/app"
	"cointracking-export-converter/internal/cointracking"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	appCtx := app.NewAppContext()

	appInstance := app.NewApp(appCtx, logger.INFO)
	csvReader := cointracking.NewCsvReader()
	txManager := cointracking.NewTxTypeManagerInitializer()
	if err := txManager.Init(); err != nil {
		log.Fatal(fmt.Sprintf("failed initializing TX manager: %s", err.Error()))
	}

	ct := cointracking.New(appCtx, csvReader, txManager)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "CoinTracking Export Converter",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        appInstance.Startup,
		Bind: []interface{}{
			appInstance,
			ct,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
