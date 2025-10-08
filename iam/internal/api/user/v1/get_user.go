package api_user_v1

import (
	"context"

	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
	user_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/user/v1"
)

func (a *Api) GetUser(
	ctx context.Context,
	req *user_v1.GetUserRequest,
) (*user_v1.GetUserResponse, error) {
	user, err := a.userService.GetUser(ctx, req.UserUuid)
	if err != nil {
		return nil, err
	}

	protoUser := model_user.ConvertUserToProto(user)

	return &user_v1.GetUserResponse{
		User: protoUser.Info,
	}, nil
}
