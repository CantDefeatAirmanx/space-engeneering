package contexts

import (
	"context"

	"go.uber.org/zap"
)

type logParamsKey string

const (
	LogSetParamsKey logParamsKey = "logSetParams"
	LogGetParamsKey logParamsKey = "logGetParams"
)

func NewCtxWithLogParams(ctx context.Context) context.Context {
	logParams := []zap.Field{}

	ctxWithLogParamsSet := context.WithValue(
		ctx,
		LogSetParamsKey,
		func(params []zap.Field) {
			logParams = append(logParams, params...)
		},
	)
	ctxWithLogParamsGet := context.WithValue(
		ctxWithLogParamsSet,
		LogGetParamsKey,
		func() []zap.Field {
			return logParams
		},
	)

	return ctxWithLogParamsGet
}

func GetLogParamsSetterFunc(ctx context.Context) func(params []zap.Field) {
	if setLogParams, ok := ctx.Value(LogSetParamsKey).(func(params []zap.Field)); ok {
		return setLogParams
	}
	return func(params []zap.Field) {}
}

func GetLogParamsGetterFunc(ctx context.Context) func() []zap.Field {
	if getLogParams, ok := ctx.Value(LogGetParamsKey).(func() []zap.Field); ok {
		return getLogParams
	}
	return func() []zap.Field {
		return []zap.Field{}
	}
}
