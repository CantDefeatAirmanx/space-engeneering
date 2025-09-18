package platform_kafka_producer

import "context"

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
				message, err := convertSaramaMessageToProducerMessage(
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
