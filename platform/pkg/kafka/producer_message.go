package platform_kafka

type ProducerMessage struct {
	Headers ProducerHeaders

	Topic string
	Key   []byte
	Value []byte
}

type ProducerHeaders map[string][]byte

func NewProducerMessage(
	topic string,
	key []byte,
	value []byte,
	opts ...ProducerMessageOption,
) ProducerMessage {
	message := ProducerMessage{
		Topic: topic,
		Key:   key,
		Value: value,
	}

	for _, opt := range opts {
		opt(&message)
	}

	return message
}

type ProducerMessageOption func(opts *ProducerMessage)

func WithHeaders(headers map[string][]byte) ProducerMessageOption {
	return func(opts *ProducerMessage) {
		opts.Headers = headers
	}
}
