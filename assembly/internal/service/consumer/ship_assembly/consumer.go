package consumer_ship_assembly

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	model_consumer_order "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/consumer/order"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

var _ model_consumer_order.OrderConsumer = (*ShipAssemblyConsumerImpl)(nil)

type ShipAssemblyConsumerImpl struct {
	consumer           platform_kafka.Consumer
	orderPaidProcessor model_consumer_order.WithProcessOrderPaidEvent
}

type InstanceParams struct {
	KafkaConsumer      platform_kafka.Consumer
	OrderPaidProcessor model_consumer_order.WithProcessOrderPaidEvent
}

func NewShipAssemblyConsumer(
	params InstanceParams,
) model_consumer_order.OrderConsumer {
	params.KafkaConsumer.SetProcessMessageErrHandlers([]func(err error){
		processKafkaErrorsHandler,
	})

	return &ShipAssemblyConsumerImpl{
		consumer:           params.KafkaConsumer,
		orderPaidProcessor: params.OrderPaidProcessor,
	}
}

func (s *ShipAssemblyConsumerImpl) SetOrderPaidProcessor(
	orderPaidProcessor model_consumer_order.WithProcessOrderPaidEvent,
) {
	s.orderPaidProcessor = orderPaidProcessor
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

func (s *ShipAssemblyConsumerImpl) handleOrderMessage(
	ctx context.Context,
	message platform_kafka.Message,
) (returnErr error) {
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
		return s.orderPaidProcessor.ProcessOrderPaidEvent(ctx, orderPaidEvent.OrderPaid)

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
