package platform_kafka

import (
	"context"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
)

type MessageHandler func(ctx context.Context, message Message) error

type Consumer interface {
	Subscribe(topics []string) error
	Consume(ctx context.Context, handler MessageHandler) error
	interfaces.WithClose
}

type Producer interface {
	ProduceAsync(ctx context.Context, message ProducerMessage) error
	ProduceSync(ctx context.Context, message ProducerMessage) (partition int32, offset int64, err error)
	interfaces.WithClose
}

type ConsumerGroup interface {
	Subscribe(topics []string) error
	Consume(ctx context.Context, handler MessageHandler) error
	CommitOffsets(ctx context.Context, offsets map[string]map[int32]int64) error
	interfaces.WithClose
}
