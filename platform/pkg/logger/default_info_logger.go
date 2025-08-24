package logger

var defaultInfoLogger = &logger{}
var isInited = false

func DefaultInfoLogger() *logger {
	if !isInited {
		Init(WithTarget(defaultInfoLogger))
		isInited = true
	}
	return defaultInfoLogger
}
