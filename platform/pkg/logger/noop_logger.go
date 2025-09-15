package logger

import (
	"context"

	"go.uber.org/zap"
)

var _ LoggerInterface = (*noopLogger)(nil)

type noopLogger struct{}

var zeroNoopLogger = &noopLogger{}

func NoopLogger() LoggerInterface {
	return zeroNoopLogger
}

func (l *noopLogger) Fatal(message string, fields ...zap.Field)                             {}
func (l *noopLogger) FatalWithCtx(ctx context.Context, message string, fields ...zap.Field) {}
func (l *noopLogger) Debug(message string, fields ...zap.Field)                             {}
func (l *noopLogger) DebugWithCtx(ctx context.Context, message string, fields ...zap.Field) {}
func (l *noopLogger) Info(message string, fields ...zap.Field)                              {}
func (l *noopLogger) InfoWithCtx(ctx context.Context, message string, fields ...zap.Field)  {}
func (l *noopLogger) Warn(message string, fields ...zap.Field)                              {}
func (l *noopLogger) WarnWithCtx(ctx context.Context, message string, fields ...zap.Field)  {}
func (l *noopLogger) Error(message string, fields ...zap.Field)                             {}
func (l *noopLogger) ErrorWithCtx(ctx context.Context, message string, fields ...zap.Field) {}
