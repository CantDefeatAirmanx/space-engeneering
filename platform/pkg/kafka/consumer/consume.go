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
	handlerErrCh := make(chan error)
	groupHandler := NewGroupHandler(topics, handler, handlerErrCh)

	for {
		kafkaErrCh := consumeWorker(
			func(handlerErrCh chan error) error {
				return k.consumerGroup.Consume(ctx, topics, groupHandler)
			},
			handlerErrCh,
		)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case kafkaErr := <-kafkaErrCh:
			for _, errHandler := range k.kafkaErrorsHandlers {
				errHandler(platform_kafka_converter.ConvertSaramaError(kafkaErr))
			}
		case handlerErr := <-handlerErrCh:
			for _, errHandler := range k.processMessageErrHandlers {
				errHandler(handlerErr)
			}
		}
	}
}

func consumeWorker(
	action func(handlerErrCh chan error) error,
	handlerErrCh chan error,
) <-chan error {
	errCh := make(chan error)

	go func() {
		err := action(handlerErrCh)
		if err != nil {
			errCh <- err
		}
	}()

	return errCh
}
