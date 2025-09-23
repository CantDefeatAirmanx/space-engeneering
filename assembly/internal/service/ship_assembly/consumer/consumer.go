package service_ship_assembly_consumer

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

var _ ShipAssemblyConsumer = (*ShipAssemblyConsumerImpl)(nil)

type ShipAssemblyConsumerImpl struct {
	consumer              platform_kafka.Consumer
	processOrderPaidEvent ProcessOrderPaidEvent
}
type ProcessOrderPaidEvent func(ctx context.Context, orderPaidEvent *order_events_v1.OrderPaidEvent) error

func NewShipAssemblyConsumer(
	consumer platform_kafka.Consumer,
	processOrderPaidEvent ProcessOrderPaidEvent,
) *ShipAssemblyConsumerImpl {
	consumer.SetProcessMessageErrHandlers([]func(err error){
		processKafkaErrorsHandler,
	})

	return &ShipAssemblyConsumerImpl{
		consumer:              consumer,
		processOrderPaidEvent: processOrderPaidEvent,
	}
}

func (s *ShipAssemblyConsumerImpl) WatchOrderPaidEvent(ctx context.Context) {
	err := s.consumer.Consume(
		ctx,
		[]string{config.Config.Kafka().OrderTopic()},
		s.handleOrderMessage,
	)

	if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
		logger.Logger().Error("Failed to consume order paid events", zap.Error(err))
	}
}

func (s *ShipAssemblyConsumerImpl) handleOrderMessage(ctx context.Context, message platform_kafka.Message) (returnErr error) {
	payload := message.Value

	var orderEventEnvelope order_events_v1.OrderEventEnvelope
	if err := proto.Unmarshal(payload, &orderEventEnvelope); err != nil {
		return fmt.Errorf(
			"%w: %s\n payload:\n %s",
			platform_kafka.ErrEncodeDecode,
			err.Error(),
			string(payload),
		)
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
		return s.processOrderPaidEvent(ctx, orderPaidEvent.OrderPaid)

	default:
		return nil
	}
}

func processKafkaErrorsHandler(err error) {
	logger.Logger().Error(
		"Ship Assembly Service: Consumer Error processing message",
		zap.Error(err),
	)
}
