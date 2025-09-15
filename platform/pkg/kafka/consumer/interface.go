package platform_kafka_consumer

import "context"

type KafkaConsumer interface {
	Consume(ctx context.Context) error
}
