package platform_mongo

import "time"

type Options struct {
	URI            string
	ConnectTimeout time.Duration
	PingTimeout    time.Duration
}

type Option func(*Options)

func WithURI(uri string) Option {
	return func(o *Options) {
		o.URI = uri
	}
}

func WithConnectTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.ConnectTimeout = timeout
	}
}

func WithPingTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.PingTimeout = timeout
	}
}
