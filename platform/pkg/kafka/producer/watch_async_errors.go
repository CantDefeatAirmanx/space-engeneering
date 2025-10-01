package platform_kafka_producer

import (
	"context"

	platform_kafka_converter "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/converter"
)

func (p *KafkaProducerImpl) watchAsyncErrors(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case err, open := <-p.asyncProducer.Errors():
			if !open {
				return
			}
			for _, callback := range p.asyncErrorsCallbacks {
				message, err := platform_kafka_converter.SaramaMessageToProducerMessage(
					err.Msg,
				)
				if err != nil {
					continue
				}
				callback(*message, err)
			}
		}
	}
}
