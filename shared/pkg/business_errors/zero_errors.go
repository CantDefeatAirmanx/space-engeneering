package business_errors

import "errors"

var errZero = errors.New("")

var (
	ErrInternalZero     = NewInternalError(errZero)
	ErrUnknownZero      = NewUnknownError(errZero)
	ErrNotFoundZero     = NewNotFoundError(errZero)
	ErrBadRequestZero   = NewBadRequestError(errZero)
	ErrConflictZero     = NewConflictError(errZero)
	ErrUnauthorizedZero = NewUnauthorizedError(errZero)
	ErrForbiddenZero    = NewForbiddenError(errZero)
)
