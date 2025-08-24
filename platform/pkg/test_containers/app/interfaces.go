package test_containers_app

import "go.uber.org/zap"

type Logger interface {
	Info(msg string, args ...zap.Field)
	Error(msg string, args ...zap.Field)
}
