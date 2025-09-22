package platform_kafka_consumer

import (
	"time"
)

type ConsumerConfig struct {
	ConsumeErrHandlers []func(err error)

	Brokers []string
	GroupID string

	InitialOffsetStart InitialOffsetStart
	SessionTimeout     time.Duration
	MaxPollRecords     int

	HeartbeatInterval time.Duration
	RequestTimeout    time.Duration
	MaxOpenRequests   int

	FetchMinBytes int
	FetchMaxBytes int
	FetchMaxWait  time.Duration

	RetryMax     int
	RetryBackoff time.Duration

	RebalanceTimeout  time.Duration
	PartitionStrategy PartitionStrategy

	EnableAutoCommit bool
}

type InitialOffsetStart string

const (
	InitialOffsetStartEarliest InitialOffsetStart = "earliest"
	InitialOffsetStartLatest   InitialOffsetStart = "latest"
	InitialOffsetStartNone     InitialOffsetStart = "none"
)

type PartitionStrategy string

const (
	PartitionStrategyRoundRobin PartitionStrategy = "round-robin"
	PartitionStrategySticky     PartitionStrategy = "sticky"
	PartitionStrategyRange      PartitionStrategy = "range"
)

func NewConsumerConfig(
	brokers []string,
	groupID string,
	opts ...ConsumerConfigOption,
) ConsumerConfig {
	cfg := ConsumerConfig{
		ConsumeErrHandlers: make([]func(err error), 0),

		Brokers:            brokers,
		GroupID:            groupID,
		InitialOffsetStart: InitialOffsetStartEarliest,

		EnableAutoCommit: false,

		SessionTimeout: 30 * time.Second,
		MaxPollRecords: 500,

		HeartbeatInterval: 3 * time.Second,
		RequestTimeout:    30 * time.Second,
		MaxOpenRequests:   5,

		FetchMinBytes: 1,
		FetchMaxBytes: 50 * 1024 * 1024,
		FetchMaxWait:  500 * time.Millisecond,

		RetryMax:     3,
		RetryBackoff: 100 * time.Millisecond,

		RebalanceTimeout:  60 * time.Second,
		PartitionStrategy: "range",
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

type ConsumerConfigOption func(opts *ConsumerConfig)

func WithAutoOffsetReset(autoOffsetReset InitialOffsetStart) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.InitialOffsetStart = autoOffsetReset
	}
}

func WithEnableAutoCommit(enableAutoCommit bool) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.EnableAutoCommit = enableAutoCommit
	}
}

func WithSessionTimeout(sessionTimeout time.Duration) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.SessionTimeout = sessionTimeout
	}
}

func WithMaxPollRecords(maxPollRecords int) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.MaxPollRecords = maxPollRecords
	}
}

func WithHeartbeatInterval(heartbeatInterval time.Duration) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.HeartbeatInterval = heartbeatInterval
	}
}

func WithRequestTimeout(requestTimeout time.Duration) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.RequestTimeout = requestTimeout
	}
}

func WithMaxOpenRequests(maxOpenRequests int) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.MaxOpenRequests = maxOpenRequests
	}
}

func WithFetchMinBytes(fetchMinBytes int) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.FetchMinBytes = fetchMinBytes
	}
}

func WithFetchMaxBytes(fetchMaxBytes int) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.FetchMaxBytes = fetchMaxBytes
	}
}

func WithFetchMaxWait(fetchMaxWait time.Duration) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.FetchMaxWait = fetchMaxWait
	}
}

func WithRetryMax(retryMax int) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.RetryMax = retryMax
	}
}

func WithRetryBackoff(retryBackoff time.Duration) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.RetryBackoff = retryBackoff
	}
}

func WithRebalanceTimeout(rebalanceTimeout time.Duration) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.RebalanceTimeout = rebalanceTimeout
	}
}

func WithPartitionStrategy(partitionStrategy PartitionStrategy) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.PartitionStrategy = partitionStrategy
	}
}

func WithConsumeErrHandler(consumeErrHandler func(err error)) ConsumerConfigOption {
	return func(opts *ConsumerConfig) {
		opts.ConsumeErrHandlers = append(opts.ConsumeErrHandlers, consumeErrHandler)
	}
}
