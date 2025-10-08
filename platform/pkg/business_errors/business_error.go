package business_errors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/helper_structs"
)

type ErrCode int64

const (
	ErrCodeUnknown ErrCode = iota
	ErrCodeInternal
	ErrCodeNotFound
	ErrCodeBadRequest
	ErrCodeConflict
	ErrCodeUnauthorized
	ErrCodeForbidden
	ErrCodeAlreadyExists
)

type BusinessError struct {
	Code     ErrCode
	Err      error
	HttpCode helper_structs.OptionalInterface[int]
	GRPCCode helper_structs.OptionalInterface[codes.Code]
}

type newBusinessErrorParams struct {
	Code     ErrCode
	Err      error
	HttpCode helper_structs.OptionalInterface[int]
	GRPCCode helper_structs.OptionalInterface[codes.Code]
}

func newBusinessError(params newBusinessErrorParams) *BusinessError {
	err := BusinessError{
		Code: params.Code,
		Err:  params.Err,
	}

	if params.GRPCCode != nil {
		err.GRPCCode = params.GRPCCode
	} else {
		err.GRPCCode = helper_structs.NewOptionalEmpty[codes.Code]()
	}

	if params.HttpCode != nil {
		err.HttpCode = params.HttpCode
	} else {
		err.HttpCode = helper_structs.NewOptionalEmpty[int]()
	}

	return &err
}

type BusinessHttpErrResponse struct {
	Code     ErrCode `json:"code"`
	Message  string  `json:"message"`
	HttpCode int     `json:"http_code"`
}

func (e *BusinessError) Error() string {
	errType := e.ErrType()

	if e.Err != nil && e.Err.Error() != "" {
		return fmt.Sprintf("%s: %s", errType, e.Err.Error())
	}

	return errType
}

func (e *BusinessError) ErrType() string {
	message, ok := businessCodesToMessages[e.Code]
	if !ok {
		message = "Unknown business error"
	}

	return message
}

func (e *BusinessError) Unwrap() error {
	return e.Err
}

func (e *BusinessError) Is(target error) bool {
	_, ok := target.(*BusinessError)
	return ok
}

func GetBusinessError(err error) *BusinessError {
	var businessError *BusinessError
	if errors.As(err, &businessError) {
		return businessError
	}
	return nil
}
