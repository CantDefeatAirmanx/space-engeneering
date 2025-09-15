package logger

type Options struct {
	Level   Level
	Env     Env
	Encoder EncoderType
	Target  *logger
}

type OptionFunc func(opts *Options)

func WithLevel(level Level) OptionFunc {
	return func(opts *Options) {
		opts.Level = level
	}
}

func WithEnv(env Env) OptionFunc {
	return func(opts *Options) {
		opts.Env = env
	}
}

func WithEncoder(encoder EncoderType) OptionFunc {
	return func(opts *Options) {
		opts.Encoder = encoder
	}
}

func WithTarget(target *logger) OptionFunc {
	return func(opts *Options) {
		opts.Target = target
	}
}
