package platform_kafka

import "time"

type Message struct {
	Headers        map[string][]byte
	Timestamp      time.Time
	BlockTimestamp time.Time

	Topic     string
	Key       []byte
	Value     []byte
	Partition int32
	Offset    int64
}
