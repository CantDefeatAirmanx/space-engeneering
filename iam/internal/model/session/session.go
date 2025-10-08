package model_session

import "time"

type Session struct {
	UUID     string
	UserUUID string

	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}

type CreateUserSessionParams struct {
	UserUUID  string
	ExpiresAt time.Time
}

type LoginWithPasswordData struct {
	Login    string
	Email    string
	Password string
}
