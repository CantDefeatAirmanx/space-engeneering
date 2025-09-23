package service_order_producer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	"github.com/CantDefeatAirmanx/space-engeneering/order/config"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	kafka_events_order "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/order/v1"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

func (o *OrderProducerImpl) ProduceOrderPaid(
	ctx context.Context,
	order kafka_events_order.OrderPaidEvent,
) (returnErr error) {
	eventUUID := uuid.Must(uuid.NewV7()).String()
	orderPaidProto := kafka_events_order.ConvertOrderPaidModelToProto(
		&order,
	)
	protoOrderPaidEnvelope := order_events_v1.OrderEventEnvelope{
		EventType: order_events_v1.OrderEventType_ORDER_EVENT_TYPE_PAID,
		EventUuid: eventUUID,
		Event: &order_events_v1.OrderEventEnvelope_OrderPaid{
			OrderPaid: &orderPaidProto,
		},
	}

	protoBytes, err := proto.Marshal(&protoOrderPaidEnvelope)
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
