package service_ship_assembly

import (
	"context"

	service_ship_assembly_consumer "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly/consumer"
	service_ship_assembly_producer "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly/producer"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

var _ ShipAssemblyService = (*ShipAssemblyServiceImpl)(nil)

type ShipAssemblyServiceImpl struct {
	repository ShipAssemblyRepository
	cancel     context.CancelFunc

	consumer service_ship_assembly_consumer.ShipAssemblyConsumer
	producer service_ship_assembly_producer.ShipAssemblyProducer
}

func NewShipAssemblyService(
	ctx context.Context,
	repository ShipAssemblyRepository,
	consumer platform_kafka.Consumer,
	producer platform_kafka.Producer,
) *ShipAssemblyServiceImpl {
	service := &ShipAssemblyServiceImpl{repository: repository}

	withCancel, cancel := context.WithCancel(ctx)
	service.cancel = cancel

	shipAssemblyConsumer := service_ship_assembly_consumer.NewShipAssemblyConsumer(
		consumer,
		service.processOrderPaidEvent,
	)
	service.consumer = shipAssemblyConsumer
	go shipAssemblyConsumer.WatchOrderPaidEvent(withCancel)

	shipAssemblyProducer := service_ship_assembly_producer.NewShipAssemblyProducer(producer)
	service.producer = shipAssemblyProducer

	return service
}
