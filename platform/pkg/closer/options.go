package closer

import "os"

type CloserOpts struct {
	Logger  Logger
	Signals []os.Signal
}

type CloserOptsFunc func(opts *CloserOpts)

func WithLogger(logger Logger) CloserOptsFunc {
	return func(opts *CloserOpts) {
		opts.Logger = logger
	}
}

func WithSignals(signals []os.Signal) CloserOptsFunc {
	return func(opts *CloserOpts) {
		opts.Signals = signals
	}
}
