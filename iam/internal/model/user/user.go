package model_user

import (
	"time"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
)

type User struct {
	UUID string
	Info UserFullInfo

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFilter struct {
	UUID  string
	Login string
	Email string
}

type UserFullInfo struct {
	UserShortInfo
	NotificationMethods []model_notification_method.NotificationMethod
}

type UserShortInfo struct {
	Login string
	Email string
}

type UserInfoWithHashPwd struct {
	UUID         string
	PasswordHash string

	UserShortInfo

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegisterData struct {
	Login               string
	Email               string
	Password            string
	NotificationMethods []model_notification_method.NotificationMethod
}
