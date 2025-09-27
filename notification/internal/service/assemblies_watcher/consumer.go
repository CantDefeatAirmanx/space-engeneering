package service_assemblies_watcher

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/CantDefeatAirmanx/space-engeneering/notification/config"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	kafka_events_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/ship-assembly/v1"
	assembly_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/assembly/v1"
)

type AssembliesWatcherConsumer struct {
	assembliesConsumer           platform_kafka.Consumer
	handleAssemblyCompletedEvent handleAssemblyCompletedEvent
}

type handleAssemblyCompletedEvent func(
	ctx context.Context,
	message kafka_events_ship_assembly.ShipAssembledEvent,
) error

func NewAssembliesWatcherConsumer(
	assembliesConsumer platform_kafka.Consumer,
	handleAssemblyCompletedMessage handleAssemblyCompletedEvent,
) *AssembliesWatcherConsumer {
	assembliesConsumer.SetProcessMessageErrHandlers([]func(err error){
		processKafkaErrorsHandler,
	})

	return &AssembliesWatcherConsumer{
		assembliesConsumer:           assembliesConsumer,
		handleAssemblyCompletedEvent: handleAssemblyCompletedMessage,
	}
}

func (a *AssembliesWatcherConsumer) ConsumeAssemblyCompletedMessage(ctx context.Context) error {
	return a.assembliesConsumer.Consume(
		ctx,
		[]string{config.Config.Kafka().AssemblyTopic()},
		a.handleAssemblyMessage,
	)
}

func (a *AssembliesWatcherConsumer) handleAssemblyMessage(
	ctx context.Context,
	message platform_kafka.Message,
) error {
	defer func() {
		if r := recover(); r != nil {
			logger.Logger().Error("Assemblies Watcher Consumer: Panic", zap.Any("panic", r))
		}
	}()

	payload := message.Value

	var assemblyEventEnvelope assembly_events_v1.ShipAssemblyEventEnvelope
	if err := proto.Unmarshal(payload, &assemblyEventEnvelope); err != nil {
		return fmt.Errorf("%w: %s", platform_kafka.ErrEncodeDecode, err.Error())
	}

	eventType := assemblyEventEnvelope.EventType

	switch eventType {
	case assembly_events_v1.ShipAssemblyEventType_SHIP_ASSEMBLY_EVENT_TYPE_ASSEMBLED:
		payload, ok := assemblyEventEnvelope.Event.(*assembly_events_v1.ShipAssemblyEventEnvelope_ShipAssembled)
		if !ok {
			return fmt.Errorf(
				"failed to process assembly completed event, eventType: %v,\n payload\n :%v",
				assemblyEventEnvelope.EventType,
				assemblyEventEnvelope.Event,
			)
		}
		return a.handleAssemblyCompletedEvent(
			ctx,
			kafka_events_ship_assembly.ConvertShipAssembledProtoToModel(payload.ShipAssembled),
		)
	default:
		return nil
	}
}

func processKafkaErrorsHandler(err error) {
	logger.Logger().Error(
		"Assemblies Watcher Consumer: Kafka errors handler",
		zap.Error(err),
	)
}
