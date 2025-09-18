package platform_kafka_producer

import "context"

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
				message, err := convertSaramaMessageToProducerMessage(
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
