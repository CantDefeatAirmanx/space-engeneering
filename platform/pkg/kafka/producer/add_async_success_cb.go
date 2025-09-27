package platform_kafka_producer

import platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"

func (p *KafkaProducerImpl) AddAsyncSuccessCb(
	cb func(success platform_kafka.ProducerMessage),
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.asyncSuccessesCallbacks = append(p.asyncSuccessesCallbacks, cb)
}
