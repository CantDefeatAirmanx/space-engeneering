package logger

import (
	"context"

	"go.uber.org/zap"
)

type LoggerInterface interface {
	Debug(message string, fields ...zap.Field)
	DebugWithCtx(ctx context.Context, message string, fields ...zap.Field)
	Info(message string, fields ...zap.Field)
	InfoWithCtx(ctx context.Context, message string, fields ...zap.Field)
	Warn(message string, fields ...zap.Field)
	WarnWithCtx(ctx context.Context, message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
	ErrorWithCtx(ctx context.Context, message string, fields ...zap.Field)
	Fatal(message string, fields ...zap.Field)
	FatalWithCtx(ctx context.Context, message string, fields ...zap.Field)
}
