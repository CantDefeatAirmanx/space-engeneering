package middleware

import (
	"net/http"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
)

func CreateWithLogParamsCtxMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctxWithLogParams := contexts.NewCtxWithLogParams(r.Context())
			next.ServeHTTP(w, r.WithContext(ctxWithLogParams))
		})
	}
}
