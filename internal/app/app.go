package app

import (
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"context"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	appCtx      interfaces.AppContext
	logLevel    logger.LogLevel
	timezoneLoc string
}

// NewApp creates a new App application struct
func NewApp(appCtx interfaces.AppContext, logLevel logger.LogLevel) *App {
	return &App{
		appCtx:   appCtx,
		logLevel: logLevel,
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.appCtx.SetContext(ctx)
	runtime.LogSetLogLevel(ctx, a.logLevel)
}

func (a *App) GetTimezoneLocation() (string, error) {
	return a.timezoneLoc, nil
}

func (a *App) SetTimezoneLocation(loc string) error {
	a.timezoneLoc = loc
	return nil
}
