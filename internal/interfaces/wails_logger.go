package interfaces

import "context"

type WailsLogger interface {
	LogPrint(ctx context.Context, message string)
	LogPrintf(ctx context.Context, format string, args ...any)
	LogTrace(ctx context.Context, message string)
	LogTracef(ctx context.Context, format string, args ...any)
	LogDebug(ctx context.Context, message string)
	LogDebugf(ctx context.Context, format string, args ...any)
	LogInfo(ctx context.Context, message string)
	LogInfof(ctx context.Context, format string, args ...any)
	LogWarning(ctx context.Context, message string)
	LogWarningf(ctx context.Context, format string, args ...any)
	LogError(ctx context.Context, message string)
	LogErrorf(ctx context.Context, format string, args ...any)
	LogFatal(ctx context.Context, message string)
	LogFatalf(ctx context.Context, format string, args ...any)
}
