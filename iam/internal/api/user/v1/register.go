package api_user_v1

import (
	"context"

	user_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/user/v1"
)

func (a *Api) Register(
	ctx context.Context, req *user_v1.RegisterRequest,
) (*user_v1.RegisterResponse, error) {
	panic("unimplemented")
}
