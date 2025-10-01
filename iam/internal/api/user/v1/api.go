package api_user_v1

import (
	user_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/user/v1"
)

var _ user_v1.UserServiceServer = (*Api)(nil)

type Api struct {
	user_v1.UnimplementedUserServiceServer
}

func NewApi() *Api {
	return &Api{}
}
