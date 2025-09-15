package platform_kafka_producer

import "context"

type KafkaProducer interface {
	Produce(ctx context.Context, message []byte) error
}
