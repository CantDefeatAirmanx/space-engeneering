package api_auth_v1

import (
	"context"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

func (a *Api) Login(
	ctx context.Context,
	req *auth_v1.LoginRequest,
) (*auth_v1.LoginResponse, error) {
	params := model_session.ConvertLoginWithPasswordDataToModel(req)

	result, err := a.authService.Login(ctx, &params)
	if err != nil {
		return nil, err
	}

	return &auth_v1.LoginResponse{
		SessionUuid: result.SessionUUID,
	}, nil
}
