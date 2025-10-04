package model_notification_method

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
)

var ErrInvalidArguments = business_errors.NewBadRequestError(errors.New("invalid arguments"))
