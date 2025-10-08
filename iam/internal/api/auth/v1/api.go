package api_auth_v1

import (
	service_auth "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/service/auth"
	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

var _ auth_v1.AuthServiceServer = (*Api)(nil)

type Api struct {
	auth_v1.UnimplementedAuthServiceServer
	authService service_auth.AuthService
}

func NewApi(
	authService service_auth.AuthService,
) *Api {
	return &Api{
		authService: authService,
	}
}
