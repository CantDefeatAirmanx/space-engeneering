package model_part

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/business_errors"
)

var (
	ErrPartInternal         = business_errors.NewInternalError(errors.New("part internal"))
	ErrPartInvalidArguments = business_errors.NewBadRequestError(errors.New("part invalid arguments"))
	ErrPartNotFound         = business_errors.NewNotFoundError(errors.New("part not found"))
)
