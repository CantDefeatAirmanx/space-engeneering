package platform_kafka_producer

import (
	"github.com/IBM/sarama"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

func convertHeaders(headers []sarama.RecordHeader) platform_kafka.ProducerHeaders {
	producerHeaders := make(platform_kafka.ProducerHeaders)

	for idx := range headers {
		header := headers[idx]
		producerHeaders[string(header.Key)] = header.Value
	}

	return producerHeaders
}

func convertSaramaMessageToProducerMessage(
	saramaMessage *sarama.ProducerMessage,
) (*platform_kafka.ProducerMessage, error) {
	key, err := saramaMessage.Key.Encode()
	if err != nil {
		return nil, err
	}

	value, err := saramaMessage.Value.Encode()
	if err != nil {
		return nil, err
	}

	mess := platform_kafka.NewProducerMessage(
		saramaMessage.Topic,
		key,
		value,
		platform_kafka.WithHeaders(convertHeaders(saramaMessage.Headers)),
	)

	return &mess, nil
}
