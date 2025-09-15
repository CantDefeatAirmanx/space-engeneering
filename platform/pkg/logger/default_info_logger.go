package logger

var (
	defaultInfoLogger = &logger{}
	isInited          = false
)

func DefaultInfoLogger() *logger {
	if !isInited {
		_ = Init(WithTarget(defaultInfoLogger)) //nolint:gosec
		isInited = true
	}
	return defaultInfoLogger
}
