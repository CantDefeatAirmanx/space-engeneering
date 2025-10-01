package api_auth_v1

import (
	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

var _ auth_v1.AuthServiceServer = (*Api)(nil)

type Api struct {
	auth_v1.UnimplementedAuthServiceServer
}

func NewApi() *Api {
	return &Api{}
}
