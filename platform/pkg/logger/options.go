package logger

type Options struct {
	Level   Level
	Env     Env
	Encoder EncoderType
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
