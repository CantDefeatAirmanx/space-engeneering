package interceptor

import (
	"context"
	"fmt"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/business_errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryErrorInterceptor(opts ...InterceptopOpt) grpc.UnaryServerInterceptor {
	options := &Options{
		Logger: logger.NoopLogger(),
	}
	for _, opt := range opts {
		opt(options)
	}

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return resp, convertError(ctx, err, info.FullMethod, options.Logger)
		}
		return resp, nil
	}
}

func convertError(
	ctx context.Context,
	err error,
	method string,
	logger Logger,
) error {
	if businessErr := business_errors.GetBusinessError(err); businessErr != nil {
		grpcStatus := business_errors.ConvertBusinessErrToGRPCStatus(businessErr)
		logBusinessError(ctx, businessErr, method, logger, grpcStatus)

		return grpcStatus.Err()
	}

	if statusErr, ok := status.FromError(err); ok {
		businessErr := business_errors.ConvertGRPCStatusToBusinessError(statusErr)
		logBusinessError(ctx, businessErr, method, logger, statusErr)
		return err
	}

	businessErr := business_errors.NewInternalError(err)
	logBusinessError(ctx, businessErr, method, logger, nil)

	return status.Error(codes.Internal, internalServerErrorMessage)
}

func logBusinessError(
	ctx context.Context,
	err *business_errors.BusinessError,
	method string,
	logger Logger,
	grpcStatus *status.Status,
) {
	logParams := []zap.Field{
		zap.Int64(logCodeKey, int64(err.Code)),
		zap.Int64(logGRPCCodeKey, int64(grpcStatus.Code())),
		zap.String(logMessageKey, grpcStatus.Message()),
		zap.String(logTypeKey, err.ErrType()),
		zap.String(logErrMessageKey, err.Error()),
		zap.String(logMethodKey, method),
	}

	if ctxLogParams, ok := ctx.Value(LogParamsKey).([]zap.Field); ok {
		logParams = append(logParams, ctxLogParams...)
	}

	logger.Error(
		fmt.Sprintf("BusinessError: %s", err.Error()),
		logParams...,
	)
}
