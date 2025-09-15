package logger

import (
	"context"

	"go.uber.org/zap"
)

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *logger) DebugWithCtx(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	fields = append(fields, getFieldsFromCtx(ctx)...)
	l.zapLogger.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *logger) InfoWithCtx(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	fields = append(fields, getFieldsFromCtx(ctx)...)
	l.zapLogger.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *logger) WarnWithCtx(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	fields = append(fields, getFieldsFromCtx(ctx)...)
	l.zapLogger.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *logger) ErrorWithCtx(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	fields = append(fields, getFieldsFromCtx(ctx)...)
	l.zapLogger.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.zapLogger.Fatal(msg, fields...)
}

func (l *logger) FatalWithCtx(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	fields = append(fields, getFieldsFromCtx(ctx)...)
	l.zapLogger.Fatal(msg, fields...)
}

func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.zapLogger.Panic(msg, fields...)
}

func (l *logger) PanicWithCtx(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	fields = append(fields, getFieldsFromCtx(ctx)...)
	l.zapLogger.Panic(msg, fields...)
}

func getFieldsFromCtx(ctx context.Context) []zap.Field {
	fields := []zap.Field{}

	if traceID, ok := ctx.Value(CtxLogFieldKeyTraceID).(string); ok {
		fields = append(fields, zap.String(string(CtxLogFieldKeyTraceID), traceID))
	}

	if userID, ok := ctx.Value(CtxLogFieldKeyUserID).(string); ok {
		fields = append(fields, zap.String(string(CtxLogFieldKeyUserID), userID))
	}

	return fields
}
