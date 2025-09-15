package platform_kafka_producer

import (
	"context"
)

var _ KafkaProducer = (*KafkaProducerImpl)(nil)

type KafkaProducerImpl struct{}

func NewKafkaProducer() KafkaProducer {
	return &KafkaProducerImpl{}
}

func (p *KafkaProducerImpl) Produce(
	ctx context.Context,
	message []byte,
) error {
	return nil
}
