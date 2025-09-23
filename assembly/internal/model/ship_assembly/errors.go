package model_ship_assembly

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/business_errors"
)

var (
	ErrAssemblyNotFound         = business_errors.NewNotFoundError(errors.New("assembly not found"))
	ErrAssemblyInvalidArguments = business_errors.NewBadRequestError(errors.New("assembly invalid arguments"))
	ErrAssemblyConflict         = business_errors.NewConflictError(errors.New("assembly conflict"))
)
