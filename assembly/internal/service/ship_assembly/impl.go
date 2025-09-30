package service_ship_assembly

import (
	"context"

	model_consumer_order "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/consumer/order"
	model_producer_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/producer/assembly"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

var _ ShipAssemblyService = (*ShipAssemblyServiceImpl)(nil)

type ShipAssemblyServiceImpl struct {
	repository ShipAssemblyRepository
	cancel     context.CancelFunc
	txManager  platform_transaction.TxManager[platform_transaction.Executor]

	consumer model_consumer_order.OrderConsumer
	producer model_producer_assembly.ShipAssemblyProducer
}

func NewShipAssemblyService(
	ctx context.Context,
	repository ShipAssemblyRepository,
	txManager platform_transaction.TxManager[platform_transaction.Executor],
	consumer model_consumer_order.OrderConsumer,
	producer model_producer_assembly.ShipAssemblyProducer,
) *ShipAssemblyServiceImpl {
	service := &ShipAssemblyServiceImpl{
		repository: repository,
		txManager:  txManager,
		consumer:   consumer,
		producer:   producer,
	}

	withCancel, cancel := context.WithCancel(ctx)
	service.cancel = cancel
	service.consumer.SetOrderPaidProcessor(service)

	go service.consumer.WatchOrderPaidEvent(withCancel)

	return service
}
