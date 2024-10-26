package wailsruntime

import (
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type WailsRuntime struct {
	appCtx interfaces.AppContext
}

func (w *WailsRuntime) OpenFileDialog(
	dialogOptions runtime.OpenDialogOptions,
) (string, error) {
	return runtime.OpenFileDialog(
		w.appCtx.Context(),
		dialogOptions,
	)
}

func (w *WailsRuntime) SaveFileDialog(
	dialogOptions runtime.SaveDialogOptions,
) (string, error) {
	return runtime.SaveFileDialog(
		w.appCtx.Context(),
		dialogOptions,
	)
}

func (w *WailsRuntime) EventsEmit(
	eventName string,
	optionalData ...any,
) {
	runtime.EventsEmit(
		w.appCtx.Context(),
		eventName,
		optionalData...,
	)
}

func NewWailsRuntime(
	appCtx interfaces.AppContext,
) *WailsRuntime {
	return &WailsRuntime{
		appCtx: appCtx,
	}
}
