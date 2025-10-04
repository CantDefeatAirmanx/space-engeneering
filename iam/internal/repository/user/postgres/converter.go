package repository_user_postgres

import (
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

func convertRepoUserToModelUserInfoWithHashPwd(user *User) *model_user.UserInfoWithHashPwd {
	return &model_user.UserInfoWithHashPwd{
		UUID:         user.UUID.String(),
		PasswordHash: user.PasswordHash,

		UserShortInfo: model_user.UserShortInfo{
			Login: user.Login,
			Email: user.Email,
		},

		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
}
