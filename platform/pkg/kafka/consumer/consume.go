package platform_kafka_consumer

import (
	"context"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	platform_kafka_converter "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/converter"
)

func (k *KafkaConsumerImpl) Consume(
	ctx context.Context,
	topics []string,
	handler platform_kafka.MessageHandler,
) error {
	groupHandler := NewGroupHandler(topics, handler)

	for {
		errCh := consumeWorker(func() error {
			return k.consumerGroup.Consume(ctx, topics, groupHandler)
		})

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			for _, handler := range k.consumeErrHandlers {
				handler(platform_kafka_converter.ConvertSaramaError(err))
			}
		}
	}
}

func consumeWorker(action func() error) <-chan error {
	errCh := make(chan error)

	go func() {
		err := action()
		if err != nil {
			errCh <- err
		}
	}()

	return errCh
}
