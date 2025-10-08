package middleware

import (
	"errors"
	"net/http"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
	platform_httpHelper "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/http_helper"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/interceptor"
	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

const (
	SessionUUIDHeaderKey = "X-Session-UUID"
)

var ErrUnauthorized = business_errors.NewUnauthorizedError(errors.New("session uuid is invalid"))

func CreateAuthMiddleware(authClient auth_v1.AuthServiceClient) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionUUID := r.Header.Get(SessionUUIDHeaderKey)
			if sessionUUID == "" {
				responseWithAuthError(w)
				return
			}

			res, err := authClient.WhoAmI(r.Context(), &auth_v1.WhoAmIRequest{
				SessionUuid: sessionUUID,
			})
			if err != nil {
				responseWithAuthError(w)
				return
			}

			withUser := interceptor.GetCtxWithUserInfo(r.Context(), res.User)
			withSessionUUID := interceptor.GetCtxWithSessionUUID(withUser, sessionUUID)

			enrichedCtx := withSessionUUID

			next.ServeHTTP(w, r.WithContext(enrichedCtx))
		})
	}
}

func responseWithAuthError(w http.ResponseWriter) {
	err := platform_httpHelper.WriteJsonResponse(
		w,
		http.StatusUnauthorized,
		business_errors.ConvertBusinessErrToHttpResponse(ErrUnauthorized),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
