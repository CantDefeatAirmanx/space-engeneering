package platform_kafka_consumer

import "time"

type ConsumerConfig struct {
	Brokers          []string
	GroupID          string
	AutoOffsetReset  string
	EnableAutoCommit bool
	SessionTimeout   time.Duration
	MaxPollRecords   int
}
