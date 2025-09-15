package closer

import "go.uber.org/zap"

type Logger interface {
	Info(message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
}
