package platform_kafka_consumer

import (
	"github.com/IBM/sarama"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	platform_kafka_converter "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/converter"
)

var _ sarama.ConsumerGroupHandler = (*groupHandler)(nil)

type groupHandler struct {
	topics       []string
	handler      platform_kafka.MessageHandler
	handlerErrCh chan error
}

func NewGroupHandler(
	topics []string,
	handler platform_kafka.MessageHandler,
	handlerErrCh chan error,
) sarama.ConsumerGroupHandler {
	return &groupHandler{topics: topics, handler: handler, handlerErrCh: handlerErrCh}
}

func (g *groupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (g *groupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
loop:
	for {
		select {
		case <-session.Context().Done():
			return session.Context().Err()
		case message, isOpen := <-claim.Messages():
			if !isOpen {
				return nil
			}

			session.MarkMessage(message, "")

			if err := g.handler(
				session.Context(),
				platform_kafka_converter.SaramaMessageToPlatformMessage(message),
			); err != nil {
				session.Commit()
				g.handlerErrCh <- err
				continue loop
			}
			session.Commit()
		}
	}
}

func (g *groupHandler) Setup(
	session sarama.ConsumerGroupSession,
) error {
	return nil
}
