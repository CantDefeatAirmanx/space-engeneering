package platform_kafka_producer

import (
	"context"

	"github.com/IBM/sarama"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

func (p *KafkaProducerImpl) ProduceSync(
	ctx context.Context,
	message platform_kafka.ProducerMessage,
) (partition int32, offset int64, err error) {
	select {
	case result := <-workerProduceSync(message, p.syncProducer):
		return result.partition, result.offset, result.err
	case <-ctx.Done():
		return 0, 0, ctx.Err()
	}
}

type resultVal struct {
	partition int32
	offset    int64
	err       error
}

func workerProduceSync(
	message platform_kafka.ProducerMessage,
	producer sarama.SyncProducer,
) <-chan resultVal {
	resultCh := make(chan resultVal)

	go func() {
		partition, offset, err := producer.SendMessage(
			newSaramaMessage(message),
		)
		resultCh <- resultVal{
			partition: partition,
			offset:    offset,
			err:       err,
		}
	}()

	return resultCh
}
