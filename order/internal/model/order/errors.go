package model_order

import (
	"errors"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/business_errors"
)

var (
	ErrOrderInternal         = business_errors.NewInternalError(errors.New("order internal"))
	ErrOrderNotFound         = business_errors.NewNotFoundError(errors.New("order not found"))
	ErrOrderConflict         = business_errors.NewConflictError(errors.New("order conflict"))
	ErrOrderInvalidArguments = business_errors.NewBadRequestError(errors.New("order invalid arguments"))
)
