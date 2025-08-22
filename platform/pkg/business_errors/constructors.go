package business_errors

import (
	"google.golang.org/grpc/codes"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/helper_structs"
)

func NewUnknownError(err error, opts ...NewBusinessErrorOption) *BusinessError {
	params := newBusinessErrorParams{
		Code: ErrCodeUnknown,
		Err:  err,
	}
	applyOptions(&params, opts...)

	return newBusinessError(params)
}

func NewInternalError(err error, opts ...NewBusinessErrorOption) *BusinessError {
	params := newBusinessErrorParams{
		Code: ErrCodeInternal,
		Err:  err,
	}
	applyOptions(&params, opts...)

	return newBusinessError(params)
}

func NewNotFoundError(err error, opts ...NewBusinessErrorOption) *BusinessError {
	params := newBusinessErrorParams{
		Code: ErrCodeNotFound,
		Err:  err,
	}
	applyOptions(&params, opts...)

	return newBusinessError(params)
}

func NewBadRequestError(err error, opts ...NewBusinessErrorOption) *BusinessError {
	params := newBusinessErrorParams{
		Code: ErrCodeBadRequest,
		Err:  err,
	}
	applyOptions(&params, opts...)

	return newBusinessError(params)
}

func NewConflictError(err error, opts ...NewBusinessErrorOption) *BusinessError {
	params := newBusinessErrorParams{
		Code: ErrCodeConflict,
		Err:  err,
	}
	applyOptions(&params, opts...)

	return newBusinessError(params)
}

func NewUnauthorizedError(err error, opts ...NewBusinessErrorOption) *BusinessError {
	params := newBusinessErrorParams{
		Code: ErrCodeUnauthorized,
		Err:  err,
	}
	applyOptions(&params, opts...)

	return newBusinessError(params)
}

func NewForbiddenError(err error, opts ...NewBusinessErrorOption) *BusinessError {
	params := newBusinessErrorParams{
		Code: ErrCodeForbidden,
		Err:  err,
	}
	applyOptions(&params, opts...)

	return newBusinessError(params)
}

type NewBusinessErrorOption func(*newBusinessErrorParams)

func WithHttpCode(code int) NewBusinessErrorOption {
	return func(params *newBusinessErrorParams) {
		params.HttpCode = helper_structs.NewOptional(code)
	}
}

func WithGRPCCode(code codes.Code) NewBusinessErrorOption {
	return func(params *newBusinessErrorParams) {
		params.GRPCCode = helper_structs.NewOptional(code)
	}
}

func applyOptions(params *newBusinessErrorParams, opts ...NewBusinessErrorOption) {
	for _, opt := range opts {
		opt(params)
	}
}
