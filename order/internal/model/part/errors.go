package model_part

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
)

var (
	ErrPartInvalidArguments = business_errors.NewBadRequestError(errors.New("part invalid arguments"))
	ErrPartNotFound         = business_errors.NewNotFoundError(errors.New("part not found"))
)
