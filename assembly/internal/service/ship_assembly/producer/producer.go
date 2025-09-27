package service_ship_assembly_producer

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	kafka_events_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/ship-assembly/v1"
	assembly_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/assembly/v1"
)

var _ ShipAssemblyProducer = (*ShipAssemblyProducerImpl)(nil)

type ShipAssemblyProducerImpl struct {
	producer platform_kafka.Producer
}

func NewShipAssemblyProducer(producer platform_kafka.Producer) *ShipAssemblyProducerImpl {
	return &ShipAssemblyProducerImpl{producer: producer}
}

func (s *ShipAssemblyProducerImpl) ProduceAssemblyCompleted(
	ctx context.Context,
	assemblyCompletedEvent kafka_events_ship_assembly.ShipAssembledEvent,
) error {
	defer func() {
		if r := recover(); r != nil {
			logger.Logger().Error("Ship Assembly Producer: Panic", zap.Any("panic", r))
		}
	}()

	eventUUID := uuid.Must(uuid.NewV7()).String()
	protoPayload := kafka_events_ship_assembly.ConvertShipAssembledModelToProto(
		&assemblyCompletedEvent,
	)
	protoPayload.EventUuid = eventUUID
	protoPayload.AssemblyUuid = eventUUID
	message := assembly_events_v1.ShipAssemblyEventEnvelope{
		EventType: assembly_events_v1.ShipAssemblyEventType_SHIP_ASSEMBLY_EVENT_TYPE_ASSEMBLED,
		EventUuid: eventUUID,
		Event: &assembly_events_v1.ShipAssemblyEventEnvelope_ShipAssembled{
			ShipAssembled: &protoPayload,
		},
	}

	protoBytes, err := proto.Marshal(&message)
	if err != nil {
		return err
	}

	_, _, err = s.producer.ProduceSync(ctx, platform_kafka.NewProducerMessage(
		config.Config.Kafka().AssemblyTopic(),
		[]byte(eventUUID),
		protoBytes,
	))
	if err != nil {
		return err
	}

	return nil
}
