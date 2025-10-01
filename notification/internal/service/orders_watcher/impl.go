package service_orders_watcher

import (
	model_notification_sender "github.com/CantDefeatAirmanx/space-engeneering/notification/internal/model/notification_sender"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

var _ OrdersWatcherService = (*OrdersWatcherServiceImpl)(nil)

type OrdersWatcherServiceImpl struct {
	ordersConsumer     platform_kafka.Consumer
	serviceConsumer    *OrdersWatcherConsumer
	notificationSender model_notification_sender.NotificationSender
}

func NewOrdersWatcherServiceImpl(
	ordersConsumer platform_kafka.Consumer,
	notificationSender model_notification_sender.NotificationSender,
) *OrdersWatcherServiceImpl {
	service := OrdersWatcherServiceImpl{
		ordersConsumer:     ordersConsumer,
		notificationSender: notificationSender,
	}

	consumer := NewOrdersWatcherConsumer(
		ordersConsumer,
		service.handleOrderPaidMessage,
	)
	service.serviceConsumer = consumer

	return &service
}
