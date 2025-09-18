package platform_kafka_producer

import (
	"time"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

type ProducerConfig struct {
	Brokers                 []string
	asyncSuccessesCallbacks []func(success platform_kafka.ProducerMessage)
	asyncErrorsCallbacks    []func(msg platform_kafka.ProducerMessage, error error)

	CompressionType  CompressionType
	FlushFrequency   time.Duration
	BatchSize        int
	MaxMessageBytes  int
	RetryMax         int
	RequiredAcks     int16
	IdempotentWrites bool
}

type CompressionType string

const (
	CompressionTypeGZIP   CompressionType = "gzip"
	CompressionTypeSnappy CompressionType = "snappy"
	CompressionTypeLZ4    CompressionType = "lz4"
	CompressionTypeZSTD   CompressionType = "zstd"
)

func NewProducerConfig(
	brokers []string, opts ...ProducerConfigOption,
) ProducerConfig {
	cfg := ProducerConfig{
		Brokers:          brokers,
		CompressionType:  "gzip",
		FlushFrequency:   10 * time.Millisecond,
		BatchSize:        1000,
		RetryMax:         10,
		RequiredAcks:     1,
		IdempotentWrites: true,

		asyncSuccessesCallbacks: make([]func(success platform_kafka.ProducerMessage), 0),
		asyncErrorsCallbacks:    make([]func(msg platform_kafka.ProducerMessage, error error), 0),
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

type ProducerConfigOption func(opts *ProducerConfig)

func WithCompressionType(compressionType CompressionType) ProducerConfigOption {
	return func(opts *ProducerConfig) {
		opts.CompressionType = compressionType
	}
}

func WithFlushFrequency(flushFrequency time.Duration) ProducerConfigOption {
	return func(opts *ProducerConfig) {
		opts.FlushFrequency = flushFrequency
	}
}

func WithBatchSize(batchSize int) ProducerConfigOption {
	return func(opts *ProducerConfig) {
		opts.BatchSize = batchSize
	}
}

func WithMaxMessageBytes(maxMessageBytes int) ProducerConfigOption {
	return func(opts *ProducerConfig) {
		opts.MaxMessageBytes = maxMessageBytes
	}
}

func WithRetryMax(retryMax int) ProducerConfigOption {
	return func(opts *ProducerConfig) {
		opts.RetryMax = retryMax
	}
}

func WithRequiredAcks(requiredAcks int16) ProducerConfigOption {
	return func(opts *ProducerConfig) {
		opts.RequiredAcks = requiredAcks
	}
}

func WithIdempotentWrites(idempotentWrites bool) ProducerConfigOption {
	return func(opts *ProducerConfig) {
		opts.IdempotentWrites = idempotentWrites
	}
}
