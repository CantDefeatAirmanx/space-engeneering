package model_user

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
)

var (
	ErrUserNotFound         = business_errors.NewNotFoundError(errors.New("user not found"))
	ErrUserAlreadyExists    = business_errors.NewConflictError(errors.New("user already exists"))
	ErrUserInvalidArguments = business_errors.NewBadRequestError(errors.New("user invalid arguments"))
)
