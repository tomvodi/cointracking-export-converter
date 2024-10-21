package wailsruntime

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Logger struct {
}

func (w *Logger) LogPrint(ctx context.Context, message string) {
	runtime.LogPrint(ctx, message)
}

func (w *Logger) LogPrintf(ctx context.Context, format string, args ...any) {
	runtime.LogPrintf(ctx, format, args...)
}

func (w *Logger) LogTrace(ctx context.Context, message string) {
	runtime.LogTrace(ctx, message)
}

func (w *Logger) LogTracef(ctx context.Context, format string, args ...any) {
	runtime.LogTracef(ctx, format, args...)
}

func (w *Logger) LogDebug(ctx context.Context, message string) {
	runtime.LogDebug(ctx, message)
}

func (w *Logger) LogDebugf(ctx context.Context, format string, args ...any) {
	runtime.LogDebugf(ctx, format, args...)
}

func (w *Logger) LogInfo(ctx context.Context, message string) {
	runtime.LogInfo(ctx, message)
}

func (w *Logger) LogInfof(ctx context.Context, format string, args ...any) {
	runtime.LogInfof(ctx, format, args...)
}

func (w *Logger) LogWarning(ctx context.Context, message string) {
	runtime.LogWarning(ctx, message)
}

func (w *Logger) LogWarningf(ctx context.Context, format string, args ...any) {
	runtime.LogWarningf(ctx, format, args...)
}

func (w *Logger) LogError(ctx context.Context, message string) {
	runtime.LogError(ctx, message)
}

func (w *Logger) LogErrorf(ctx context.Context, format string, args ...any) {
	runtime.LogErrorf(ctx, format, args...)
}

func (w *Logger) LogFatal(ctx context.Context, message string) {
	runtime.LogFatal(ctx, message)
}

func (w *Logger) LogFatalf(ctx context.Context, format string, args ...any) {
	runtime.LogFatalf(ctx, format, args...)
}

func NewLog() *Logger {
	return &Logger{}
}
