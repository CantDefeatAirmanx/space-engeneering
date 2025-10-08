package model_session

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
)

func ConvertSessionToProto(
	model *Session,
) *common_v1.Session {
	return &common_v1.Session{
		Uuid:      model.UUID,
		CreatedAt: timestamppb.New(model.CreatedAt),
		UpdatedAt: timestamppb.New(model.UpdatedAt),
		ExpiresAt: timestamppb.New(model.ExpiresAt),
	}
}

func ConvertLoginWithPasswordDataToModel(
	proto *auth_v1.LoginRequest,
) LoginWithPasswordData {
	return LoginWithPasswordData{
		Login:    proto.Login,
		Email:    proto.Email,
		Password: proto.Password,
	}
}
