package service_ship_assembly_producer

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	kafka_events_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/ship-assembly/v1"
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
	eventUUID := uuid.Must(uuid.NewV7()).String()

	protoPayload := kafka_events_ship_assembly.ConvertShipAssembledModelToProto(
		&assemblyCompletedEvent,
	)
	protoPayload.EventUuid = eventUUID

	protoBytes, err := proto.Marshal(&protoPayload)
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
