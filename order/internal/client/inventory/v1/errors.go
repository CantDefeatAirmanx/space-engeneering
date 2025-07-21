package client_inventory_v1

import (
	"errors"
	"fmt"
)

var (
	ErrServiceUnavailable  = errors.New("service unavailable")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidArguments    = errors.New("invalid arguments")
	ErrTimeoutExceeded     = errors.New("timeout exceeded")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrPermissionDenied    = errors.New("permission denied")
)

type ErrPartNotFound struct {
	UUID string
	Err  error
}

func (e ErrPartNotFound) Error() string {
	return fmt.Sprintf("part %s not found", e.UUID)
}

func (e ErrPartNotFound) Unwrap() error {
	return e.Err
}

func (e ErrPartNotFound) Is(target error) bool {
	_, ok := target.(*ErrPartNotFound)
	return ok
}
