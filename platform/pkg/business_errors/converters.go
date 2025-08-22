package business_errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func ConvertGRPCStatusToBusinessError(status *status.Status) *BusinessError {
	grpcCode := status.Code()
	businessErr, ok := grpcCodesToBusinessErrors[grpcCode]
	if !ok {
		businessErr = ErrUnknownZero
	}

	return businessErr
}
