package interceptor

type logParamsKey string

var (
	LogSetParamsKey logParamsKey = "logSetParams"

	logCodeKey       = "code"
	logGRPCCodeKey   = "grpc_code"
	logMessageKey    = "message"
	logErrMessageKey = "errMessage"
	logTypeKey       = "type"
	logMethodKey     = "method"

	internalServerErrorMessage = "internal server error"
)
