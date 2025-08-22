package interceptor

import (
	"context"

	"google.golang.org/grpc"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
)

func WithLogParamsInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx1 context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		ctxWithLogParams := contexts.NewCtxWithLogParams(ctx1)

		resp, err := handler(ctxWithLogParams, req)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}
