package interceptor

type Options struct {
	Logger Logger
}

type InterceptopOpt func(o *Options)

func WithLogger(logger Logger) InterceptopOpt {
	return func(o *Options) {
		o.Logger = logger
	}
}
