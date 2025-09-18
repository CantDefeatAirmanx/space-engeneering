package platform_kafka_producer

import (
	"context"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

func (p *KafkaProducerImpl) ProduceAsync(
	ctx context.Context,
	message platform_kafka.ProducerMessage,
) error {
	saramaMessage := newSaramaMessage(message)

	select {
	case p.asyncProducer.Input() <- saramaMessage:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
