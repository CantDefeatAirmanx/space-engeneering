package logger

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelFatal Level = "fatal"
)

type Env string

const (
	EnvDev  Env = "dev"
	EnvProd Env = "prod"
)

type EncoderType string

const (
	EncoderTypeJSON    EncoderType = "json"
	EncoderTypeConsole EncoderType = "console"
)

type CtxLogFieldKey string

const (
	CtxLogFieldKeyTraceID CtxLogFieldKey = "trace_id"
	CtxLogFieldKeyUserID  CtxLogFieldKey = "user_id"
)
