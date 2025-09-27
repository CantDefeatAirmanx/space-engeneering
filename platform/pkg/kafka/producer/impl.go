package platform_kafka_producer

import (
	"context"
	"sync"

	"github.com/IBM/sarama"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

var _ platform_kafka.Producer = (*KafkaProducerImpl)(nil)

type KafkaProducerImpl struct {
	syncProducer  sarama.SyncProducer
	asyncProducer sarama.AsyncProducer

	asyncSuccessesCallbacks []func(success platform_kafka.ProducerMessage)
	asyncErrorsCallbacks    []func(msg platform_kafka.ProducerMessage, error error)

	mu     sync.Mutex
	cancel context.CancelFunc
}

func NewKafkaProducer(
	ctx context.Context,
	brokers []string,
	opts ...ProducerConfigOption,
) (*KafkaProducerImpl, error) {
	cfg := NewProducerConfig(brokers, opts...)

	saramaCfg := sarama.NewConfig()

	saramaCfg.Version = platform_kafka.KafkaVersion
	saramaCfg.Metadata.AllowAutoTopicCreation = false

	saramaCfg.Producer.Compression = compressionTypes[cfg.CompressionType]

	saramaCfg.Producer.Flush.Frequency = cfg.FlushFrequency
	saramaCfg.Producer.Flush.Bytes = cfg.BatchSize
	saramaCfg.Producer.MaxMessageBytes = cfg.MaxMessageBytes

	saramaCfg.Producer.Retry.Max = cfg.RetryMax
	saramaCfg.Producer.RequiredAcks = sarama.RequiredAcks(cfg.RequiredAcks)
	saramaCfg.Producer.Idempotent = cfg.IdempotentWrites

	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.Return.Errors = true

	// Устанавливаем MaxOpenRequests
	// Для идемпотентности рекомендуется 1 для строгого порядка
	saramaCfg.Net.MaxOpenRequests = cfg.MaxOpenRequests

	syncProducer, err := sarama.NewSyncProducer(
		cfg.Brokers,
		saramaCfg,
	)
	if err != nil {
		return nil, err
	}

	asyncProducer, err := sarama.NewAsyncProducer(
		cfg.Brokers,
		saramaCfg,
	)
	if err != nil {
		return nil, err
	}

	ctxWithCancel, cancel := context.WithCancel(ctx)

	producer := KafkaProducerImpl{
		syncProducer:  syncProducer,
		asyncProducer: asyncProducer,

		asyncSuccessesCallbacks: cfg.asyncSuccessesCallbacks,
		asyncErrorsCallbacks:    cfg.asyncErrorsCallbacks,

		cancel: cancel,
	}

	go func() {
		producer.watchAsyncSuccesses(ctxWithCancel)
	}()
	go func() {
		producer.watchAsyncErrors(ctxWithCancel)
	}()

	return &producer, nil
}

var compressionTypes = map[CompressionType]sarama.CompressionCodec{
	CompressionTypeGZIP:   sarama.CompressionGZIP,
	CompressionTypeSnappy: sarama.CompressionSnappy,
	CompressionTypeLZ4:    sarama.CompressionLZ4,
	CompressionTypeZSTD:   sarama.CompressionZSTD,
}
