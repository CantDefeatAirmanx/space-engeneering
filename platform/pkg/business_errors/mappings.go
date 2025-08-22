package business_errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
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

var grpcCodesToBusinessErrors = map[codes.Code]*BusinessError{
	codes.Unknown:          ErrUnknownZero,
	codes.Internal:         ErrInternalZero,
	codes.NotFound:         ErrNotFoundZero,
	codes.InvalidArgument:  ErrBadRequestZero,
	codes.AlreadyExists:    ErrConflictZero,
	codes.Unauthenticated:  ErrUnauthorizedZero,
	codes.PermissionDenied: ErrForbiddenZero,
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
