package model_part

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/business_errors"
)

var (
	ErrPartNotFound = business_errors.NewNotFoundError(errors.New("part not found"))
	ErrPartInternal = business_errors.NewInternalError(errors.New("part internal"))
)
