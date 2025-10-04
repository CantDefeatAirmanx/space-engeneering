package model_session

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
)

var (
	ErrInvalidArguments   = business_errors.NewBadRequestError(errors.New("invalid arguments"))
	ErrInvalidCredentials = business_errors.NewUnauthorizedError(errors.New("invalid credentials"))
	ErrUnauthorized       = business_errors.NewUnauthorizedError(errors.New("unauthorized"))
	ErrSessionExpired     = business_errors.NewUnauthorizedError(errors.New("session expired"))
	ErrNotFound           = business_errors.NewNotFoundError(errors.New("not found"))
	ErrConflict           = business_errors.NewConflictError(errors.New("conflict"))
)
