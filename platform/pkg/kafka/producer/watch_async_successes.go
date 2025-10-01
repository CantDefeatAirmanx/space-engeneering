package platform_kafka_producer

import (
	"context"

	platform_kafka_converter "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/converter"
)

func (p *KafkaProducerImpl) watchAsyncSuccesses(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case success, open := <-p.asyncProducer.Successes():
			if !open {
				return
			}
			for _, callback := range p.asyncSuccessesCallbacks {
				message, err := platform_kafka_converter.SaramaMessageToProducerMessage(
					success,
				)
				if err != nil {
					continue
				}

				callback(*message)
			}
		}
	}
}
