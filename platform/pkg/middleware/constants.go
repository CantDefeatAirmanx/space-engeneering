package middleware

const (
	RequestIDCtxKey reqIdKey = "reqId"

	RequestIdLogKey  = "req_id"
	StatusCodeLogKey = "status_code"
)

type reqIdKey string
