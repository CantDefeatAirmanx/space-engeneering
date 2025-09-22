package platform_kafka_consumer

import (
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	platform_kafka_converter "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/converter"
	"github.com/IBM/sarama"
)

var _ sarama.ConsumerGroupHandler = (*groupHandler)(nil)

type groupHandler struct {
	topics  []string
	handler platform_kafka.MessageHandler
}

func NewGroupHandler(
	topics []string,
	handler platform_kafka.MessageHandler,
) sarama.ConsumerGroupHandler {
	return &groupHandler{topics: topics, handler: handler}
}

func (g *groupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (g *groupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
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
				return err
			}
		}
	}
}

func (g *groupHandler) Setup(
	session sarama.ConsumerGroupSession,
) error {
	return nil
}
