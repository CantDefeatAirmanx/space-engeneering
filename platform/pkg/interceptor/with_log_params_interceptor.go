package interceptor

import (
	"context"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	"google.golang.org/grpc"
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
