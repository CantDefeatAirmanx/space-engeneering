package platform_telegram

type MessageOptions struct {
	ThreadId int
}

type MessageOption func(opts *MessageOptions)

func WithThreadId(threadId int) MessageOption {
	return func(opts *MessageOptions) {
		opts.ThreadId = threadId
	}
}
