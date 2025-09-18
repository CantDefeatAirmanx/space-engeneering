package platform_kafka_producer

import platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"

func (p *KafkaProducerImpl) AddAsyncErrorCb(
	cb func(msg platform_kafka.ProducerMessage, error error),
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.asyncErrorsCallbacks = append(p.asyncErrorsCallbacks, cb)
}
