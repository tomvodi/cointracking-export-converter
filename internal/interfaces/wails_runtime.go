package interfaces

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type WailsRuntime interface {
	OpenFileDialog(dialogOptions runtime.OpenDialogOptions) (string, error)
	SaveFileDialog(dialogOptions runtime.SaveDialogOptions) (string, error)
	EventsEmit(eventName string, optionalData ...any)
}
