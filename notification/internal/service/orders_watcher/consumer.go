package service_orders_watcher

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/CantDefeatAirmanx/space-engeneering/notification/config"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	kafka_events_order "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/order/v1"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

type OrdersWatcherConsumer struct {
	ordersConsumer platform_kafka.Consumer

	handleOrderPaidMessage handleOrderPaid
}

type handleOrderPaid func(
	ctx context.Context,
	message kafka_events_order.OrderPaidEvent,
) error

func NewOrdersWatcherConsumer(
	ordersConsumer platform_kafka.Consumer,
	handleOrderPaidMessage handleOrderPaid,
) *OrdersWatcherConsumer {
	ordersConsumer.SetProcessMessageErrHandlers([]func(err error){
		processKafkaOrdersErrorsHandler,
	})

	return &OrdersWatcherConsumer{
		ordersConsumer:         ordersConsumer,
		handleOrderPaidMessage: handleOrderPaidMessage,
	}
}

func (o *OrdersWatcherConsumer) ConsumeOrders(ctx context.Context) error {
	return o.ordersConsumer.Consume(
		ctx,
		[]string{config.Config.Kafka().OrderTopic()},
		o.handleOrderMessage,
	)
}

func (o *OrdersWatcherConsumer) handleOrderMessage(
	ctx context.Context,
	message platform_kafka.Message,
) error {
	payload := message.Value

	var orderEventEnvelope order_events_v1.OrderEventEnvelope
	if err := proto.Unmarshal(payload, &orderEventEnvelope); err != nil {
		return fmt.Errorf("%w: %s", platform_kafka.ErrEncodeDecode, err.Error())
	}

	switch orderEventEnvelope.EventType {
	case order_events_v1.OrderEventType_ORDER_EVENT_TYPE_PAID:
		orderPaidEvent, ok := orderEventEnvelope.Event.(*order_events_v1.OrderEventEnvelope_OrderPaid)
		if !ok {
			return fmt.Errorf(
				"failed to process order paid event, eventType: %v,\n payload\n :%v",
				orderEventEnvelope.EventType,
				orderEventEnvelope.Event,
			)
		}

		return o.handleOrderPaidMessage(
			ctx,
			kafka_events_order.ConvertOrderPaidProtoToModel(orderPaidEvent.OrderPaid),
		)

	default:
		return nil
	}
}

func processKafkaOrdersErrorsHandler(err error) {
	logger.Logger().Error(
		"Orders Watcher Consumer: Kafka errors handler",
		zap.Error(err),
	)
}
