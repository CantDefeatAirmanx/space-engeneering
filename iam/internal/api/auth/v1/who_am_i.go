package api_auth_v1

import (
	"context"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
	service_auth "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/service/auth"
	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

func (a *Api) WhoAmI(
	ctx context.Context,
	req *auth_v1.WhoAmIRequest,
) (*auth_v1.WhoAmIResponse, error) {
	result, err := a.authService.WhoAmI(ctx, service_auth.WhoAmIParams{
		SessionUUID: req.SessionUuid,
	})
	if err != nil {
		return nil, err
	}

	sessionModel := result.Session
	return &auth_v1.WhoAmIResponse{
		Session: model_session.ConvertSessionToProto(sessionModel),
		User:    model_user.ConvertUserToProto(result.User),
	}, nil
}
