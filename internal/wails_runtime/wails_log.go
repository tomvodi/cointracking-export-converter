package wails_runtime

import (
	"context"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type wLog struct {
}

func (w *wLog) LogPrint(ctx context.Context, message string) {
	runtime.LogPrint(ctx, message)
}

func (w *wLog) LogPrintf(ctx context.Context, format string, args ...interface{}) {
	runtime.LogPrintf(ctx, format, args...)
}

func (w *wLog) LogTrace(ctx context.Context, message string) {
	runtime.LogTrace(ctx, message)
}

func (w *wLog) LogTracef(ctx context.Context, format string, args ...interface{}) {
	runtime.LogTracef(ctx, format, args...)
}

func (w *wLog) LogDebug(ctx context.Context, message string) {
	runtime.LogDebug(ctx, message)
}

func (w *wLog) LogDebugf(ctx context.Context, format string, args ...interface{}) {
	runtime.LogDebugf(ctx, format, args...)
}

func (w *wLog) LogInfo(ctx context.Context, message string) {
	runtime.LogInfo(ctx, message)
}

func (w *wLog) LogInfof(ctx context.Context, format string, args ...interface{}) {
	runtime.LogInfof(ctx, format, args...)
}

func (w *wLog) LogWarning(ctx context.Context, message string) {
	runtime.LogWarning(ctx, message)
}

func (w *wLog) LogWarningf(ctx context.Context, format string, args ...interface{}) {
	runtime.LogWarningf(ctx, format, args...)
}

func (w *wLog) LogError(ctx context.Context, message string) {
	runtime.LogError(ctx, message)
}

func (w *wLog) LogErrorf(ctx context.Context, format string, args ...interface{}) {
	runtime.LogErrorf(ctx, format, args...)
}

func (w *wLog) LogFatal(ctx context.Context, message string) {
	runtime.LogFatal(ctx, message)
}

func (w *wLog) LogFatalf(ctx context.Context, format string, args ...interface{}) {
	runtime.LogFatalf(ctx, format, args...)
}

func NewWailsLog() interfaces.WailsLog {
	return &wLog{}
}
