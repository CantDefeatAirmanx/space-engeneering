package service_orders_watcher

import (
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	platform_telegram "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/telegram"
)

var _ OrdersWatcherService = (*OrdersWatcherServiceImpl)(nil)

type OrdersWatcherServiceImpl struct {
	ordersConsumer  platform_kafka.Consumer
	serviceConsumer *OrdersWatcherConsumer
	telegramClient  platform_telegram.TelegramClient
}

func NewOrdersWatcherServiceImpl(
	ordersConsumer platform_kafka.Consumer,
	telegramClient platform_telegram.TelegramClient,
) *OrdersWatcherServiceImpl {
	service := OrdersWatcherServiceImpl{
		ordersConsumer: ordersConsumer,
		telegramClient: telegramClient,
	}

	consumer := NewOrdersWatcherConsumer(
		ordersConsumer,
		service.handleOrderPaidMessage,
	)
	service.serviceConsumer = consumer

	return &service
}
