package interceptor

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	business_errors2 "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
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
		logParams := contexts.GetLogParamsGetterFunc(ctx)()

		resp, err := handler(ctx, req)
		if err != nil {
			return resp, convertError(err, info.FullMethod, options.Logger, logParams)
		}
		return resp, nil
	}
}

func convertError(
	err error,
	method string,
	logger Logger,
	additionLogParams []zap.Field,
) error {
	if businessErr := business_errors2.GetBusinessError(err); businessErr != nil {
		grpcStatus := business_errors2.ConvertBusinessErrToGRPCStatus(businessErr)
		logBusinessError(businessErr, method, logger, grpcStatus, additionLogParams)

		return grpcStatus.Err()
	}

	if statusErr, ok := status.FromError(err); ok {
		businessErr := business_errors2.ConvertGRPCStatusToBusinessError(statusErr)
		logBusinessError(businessErr, method, logger, statusErr, additionLogParams)
		return err
	}

	businessErr := business_errors2.NewInternalError(err)
	logBusinessError(businessErr, method, logger, nil, additionLogParams)

	return status.Error(codes.Internal, internalServerErrorMessage)
}

func logBusinessError(
	err *business_errors2.BusinessError,
	method string,
	logger Logger,
	grpcStatus *status.Status,
	additionLogParams []zap.Field,
) {
	logParams := []zap.Field{
		zap.Int64(logCodeKey, int64(err.Code)),
		zap.Int64(logGRPCCodeKey, int64(grpcStatus.Code())),
		zap.String(logMessageKey, grpcStatus.Message()),
		zap.String(logTypeKey, err.ErrType()),
		zap.String(logErrMessageKey, err.Error()),
		zap.String(logMethodKey, method),
	}

	if len(additionLogParams) > 0 {
		logParams = append(logParams, additionLogParams...)
	}

	logger.Error(
		fmt.Sprintf("BusinessError: %s", err.Error()),
		logParams...,
	)
}
