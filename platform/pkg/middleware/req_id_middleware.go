package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func CreateReqIdMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId, err := uuid.NewV7()
			if err != nil {
				requestId = uuid.Must(uuid.NewV7())
			}

			ctxWithReqId := context.WithValue(r.Context(), RequestIDCtxKey, requestId.String())
			reqWithReqId := r.WithContext(ctxWithReqId)

			next.ServeHTTP(w, reqWithReqId)
		})
	}
}
