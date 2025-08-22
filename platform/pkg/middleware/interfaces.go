package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type Middleware func(next http.Handler) http.Handler

type Logger interface {
	Info(message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
}
