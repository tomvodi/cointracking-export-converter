package interfaces

import "context"

type WailsLog interface {
	LogPrint(ctx context.Context, message string)
	LogPrintf(ctx context.Context, format string, args ...interface{})
	LogTrace(ctx context.Context, message string)
	LogTracef(ctx context.Context, format string, args ...interface{})
	LogDebug(ctx context.Context, message string)
	LogDebugf(ctx context.Context, format string, args ...interface{})
	LogInfo(ctx context.Context, message string)
	LogInfof(ctx context.Context, format string, args ...interface{})
	LogWarning(ctx context.Context, message string)
	LogWarningf(ctx context.Context, format string, args ...interface{})
	LogError(ctx context.Context, message string)
	LogErrorf(ctx context.Context, format string, args ...interface{})
	LogFatal(ctx context.Context, message string)
	LogFatalf(ctx context.Context, format string, args ...interface{})
}
