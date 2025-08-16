package business_errors

import (
	"errors"
	"fmt"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/helper_structs"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

var businessCodesToGRPCCodes = map[ErrCode]codes.Code{
	ErrCodeUnknown:       codes.Unknown,
	ErrCodeInternal:      codes.Internal,
	ErrCodeNotFound:      codes.NotFound,
	ErrCodeBadRequest:    codes.InvalidArgument,
	ErrCodeConflict:      codes.AlreadyExists,
	ErrCodeUnauthorized:  codes.Unauthenticated,
	ErrCodeForbidden:     codes.PermissionDenied,
	ErrCodeAlreadyExists: codes.AlreadyExists,
}

var businessCodesToHttpCodes = map[ErrCode]int{
	ErrCodeUnknown:       http.StatusInternalServerError,
	ErrCodeInternal:      http.StatusInternalServerError,
	ErrCodeNotFound:      http.StatusNotFound,
	ErrCodeBadRequest:    http.StatusBadRequest,
	ErrCodeConflict:      http.StatusConflict,
	ErrCodeUnauthorized:  http.StatusUnauthorized,
	ErrCodeForbidden:     http.StatusForbidden,
	ErrCodeAlreadyExists: http.StatusConflict,
}

var businessCodesToMessages = map[ErrCode]string{
	ErrCodeUnknown:       "Unknown business error",
	ErrCodeInternal:      "Internal server error",
	ErrCodeNotFound:      "Not found",
	ErrCodeBadRequest:    "Bad request",
	ErrCodeConflict:      "Conflict",
	ErrCodeUnauthorized:  "Unauthorized",
	ErrCodeForbidden:     "Forbidden",
	ErrCodeAlreadyExists: "Already exists",
}

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
	Code     ErrCode
	Message  string
	HttpCode int
}

func (e *BusinessError) Error() string {
	message, ok := businessCodesToMessages[e.Code]
	if !ok {
		message = "Unknown business error"
	}

	if e.Err != nil {
		return fmt.Sprintf("%s: %s", message, e.Err.Error())
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

func ConvertBusinessErrToGRPCStatus(err *BusinessError) *status.Status {
	grpcCode, isSet := err.GRPCCode.GetValue(), err.GRPCCode.IsSet()

	if isSet {
		return status.New(grpcCode, err.Error())
	}

	grpcCode, ok := businessCodesToGRPCCodes[err.Code]
	if !ok {
		grpcCode = codes.Unknown
	}

	return status.New(grpcCode, err.Error())
}

func ConvertBusinessErrToHttpResponse(err *BusinessError) *BusinessHttpErrResponse {
	httpCode, isSet := err.HttpCode.GetValue(), err.HttpCode.IsSet()

	if isSet {
		return &BusinessHttpErrResponse{
			Code:     err.Code,
			Message:  err.Error(),
			HttpCode: httpCode,
		}
	}

	httpCode, ok := businessCodesToHttpCodes[err.Code]
	if !ok {
		httpCode = http.StatusInternalServerError
	}

	return &BusinessHttpErrResponse{
		Code:     err.Code,
		Message:  err.Error(),
		HttpCode: httpCode,
	}
}
