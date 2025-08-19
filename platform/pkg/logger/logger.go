package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *logger
	once         sync.Once
	dynamicLevel zap.AtomicLevel
)

var _ LoggerInterface = (*logger)(nil)

type logger struct {
	zapLogger *zap.Logger
}

func Logger() *logger {
	return globalLogger
}

func Init(opts ...OptionFunc) error {
	once.Do(func() {
		params := Options{
			Level:   LevelInfo,
			Env:     EnvProd,
			Encoder: EncoderTypeJSON,
		}
		for _, opt := range opts {
			opt(&params)
		}

		dynamicLevel = zap.NewAtomicLevelAt(getZapLevel(params.Level))
		encoder := getEncoder(params.Encoder, getEncoderConfig())

		core := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			dynamicLevel,
		)

		zLogger := zap.New(core)

		globalLogger = &logger{
			zapLogger: zLogger,
		}
	})

	return nil
}

func SetLevel(level Level) {
	dynamicLevel.SetLevel(getZapLevel(level))
}

func getEncoder(encoderType EncoderType, cfg zapcore.EncoderConfig) zapcore.Encoder {
	switch encoderType {
	case EncoderTypeJSON:
		return zapcore.NewJSONEncoder(cfg)
	case EncoderTypeConsole:
		return zapcore.NewConsoleEncoder(cfg)
	default:
		return zapcore.NewJSONEncoder(cfg)
	}
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func getZapLevel(level Level) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
