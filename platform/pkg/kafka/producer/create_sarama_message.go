package platform_kafka_producer

import (
	"github.com/IBM/sarama"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

func newSaramaMessage(message platform_kafka.ProducerMessage) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic:   message.Topic,
		Key:     sarama.ByteEncoder(message.Key),
		Value:   sarama.ByteEncoder(message.Value),
		Headers: getSaramaHeaders(message.Headers),
	}
}

func getSaramaHeaders(headers map[string][]byte) []sarama.RecordHeader {
	saramaHeaders := make([]sarama.RecordHeader, 0, len(headers))
	for key := range headers {
		value := headers[key]
		saramaHeaders = append(saramaHeaders, sarama.RecordHeader{
			Key:   []byte(key),
			Value: value,
		})
	}
	return saramaHeaders
}
