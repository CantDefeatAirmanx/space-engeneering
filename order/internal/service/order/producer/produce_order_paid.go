package service_order_producer

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"

	"github.com/CantDefeatAirmanx/space-engeneering/order/config"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	kafka_events_order "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/order/v1"
)

func (o *OrderProducerImpl) ProduceOrderPaid(
	ctx context.Context,
	order kafka_events_order.OrderPaidEvent,
) error {
	eventUUID := uuid.Must(uuid.NewV7()).String()
	protoPayload := kafka_events_order.ConvertOrderPaidModelToProto(
		&order,
	)
	protoBytes, err := proto.Marshal(&protoPayload)
	if err != nil {
		return fmt.Errorf("%w: %s", model_order.ErrOrderProducer, err.Error())
	}

	_, _, err = o.producer.ProduceSync(ctx, platform_kafka.NewProducerMessage(
		config.Config.Kafka().OrderTopic(),
		[]byte(eventUUID),
		protoBytes,
	))
	if err != nil {
		return fmt.Errorf("%w: %s", model_order.ErrOrderProducer, err.Error())
	}

	return nil
}
