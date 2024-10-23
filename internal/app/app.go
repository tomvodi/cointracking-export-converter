package app

import (
	"context"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	appCtx   interfaces.AppContext
	logLevel logger.LogLevel
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.appCtx.SetContext(ctx)
	runtime.LogSetLogLevel(ctx, a.logLevel)
}

// NewApp creates a new App application struct
func NewApp(
	appCtx interfaces.AppContext,
	logLevel logger.LogLevel,
) *App {
	return &App{
		appCtx:   appCtx,
		logLevel: logLevel,
	}
}
