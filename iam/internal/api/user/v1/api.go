package api_user_v1

import (
	service_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/service/user"
	user_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/user/v1"
)

var _ user_v1.UserServiceServer = (*Api)(nil)

type Api struct {
	user_v1.UnimplementedUserServiceServer
	userService service_user.UserService
}

func NewApi(service service_user.UserService) *Api {
	return &Api{
		userService: service,
	}
}
