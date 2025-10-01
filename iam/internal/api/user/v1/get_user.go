package api_user_v1

import (
	"context"

	user_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/user/v1"
)

func (a *Api) GetUser(
	ctx context.Context,
	req *user_v1.GetUserRequest,
) (*user_v1.GetUserResponse, error) {
	panic("unimplemented")
}
