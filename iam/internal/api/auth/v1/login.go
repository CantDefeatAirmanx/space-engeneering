package api_auth_v1

import (
	"context"

	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

func (a *Api) Login(
	ctx context.Context,
	req *auth_v1.LoginRequest,
) (*auth_v1.LoginResponse, error) {
	panic("unimplemented")
}
