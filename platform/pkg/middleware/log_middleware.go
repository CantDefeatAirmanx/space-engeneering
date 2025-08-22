package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
)

func CreateLogMiddleware(logger Logger) Middleware {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId := r.Context().Value(RequestIDCtxKey)
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)

			logParams := contexts.GetLogParamsGetterFunc(r.Context())()
			logger.Info(
				fmt.Sprintf(
					"request: %s %s reqId: %s took %s",
					r.Method,
					r.URL.Path,
					reqId,
					duration,
				),
				logParams...,
			)
		})
	}
}
