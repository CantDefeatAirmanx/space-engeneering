package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	Validate() error
}

func ValidateInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if validator, ok := req.(Validator); ok {
			if err := validator.Validate(); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
			}
		}

		return handler(ctx, req)
	}
}
