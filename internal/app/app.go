package app

import (
	"cointracking-export-converter/internal/interfaces"
	"context"
)

// App struct
type App struct {
	appCtx interfaces.AppContext
}

// NewApp creates a new App application struct
func NewApp(appCtx interfaces.AppContext) *App {
	return &App{
		appCtx: appCtx,
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.appCtx.SetContext(ctx)
}
