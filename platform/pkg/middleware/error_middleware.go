package middleware

import (
	"fmt"
	"net/http"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	"go.uber.org/zap"
)

func CreateErrorLoggingMiddleware(logger Logger) Middleware {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrappedWriter := newWriterWrapper(w)
			next.ServeHTTP(wrappedWriter, r)
			statusCode := wrappedWriter.StatusCode()

			logParams := []zap.Field{
				zap.Int(StatusCodeLogKey, statusCode),
				zap.String(RequestIdLogKey, r.Context().Value(RequestIDCtxKey).(string)),
			}
			logParamsFromCtx := contexts.GetLogParamsGetterFunc(r.Context())()
			logParams = append(logParams, logParamsFromCtx...)

			if statusCode >= 400 && statusCode < 500 {
				logger.Error(
					fmt.Sprintf("Client Error: %s %s", r.Method, r.URL.Path),
					logParams...,
				)
			}

			if statusCode >= 500 {
				logger.Error(
					fmt.Sprintf("Server Error: %s %s", r.Method, r.URL.Path),
					logParams...,
				)
			}
		})
	}
}

var _ http.ResponseWriter = (*writerWrapper)(nil)

type writerWrapper struct {
	writer     http.ResponseWriter
	statusCode int
	written    bool
}

func newWriterWrapper(writer http.ResponseWriter) *writerWrapper {
	return &writerWrapper{
		writer:     writer,
		statusCode: http.StatusOK,
		written:    false,
	}
}

func (w *writerWrapper) Header() http.Header {
	return w.writer.Header()
}

func (w *writerWrapper) WriteHeader(statusCode int) {
	if w.written {
		return
	}
	w.statusCode = statusCode
	w.written = true
	w.writer.WriteHeader(statusCode)
}

func (w *writerWrapper) Write(data []byte) (int, error) {
	if !w.written {
		w.WriteHeader(http.StatusOK)
		w.written = true
	}
	return w.writer.Write(data)
}

func (w *writerWrapper) StatusCode() int {
	return w.statusCode
}
