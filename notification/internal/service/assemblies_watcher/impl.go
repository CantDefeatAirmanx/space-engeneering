package service_assemblies_watcher

import (
	model_notification_sender "github.com/CantDefeatAirmanx/space-engeneering/notification/internal/model/notification_sender"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

var _ AssembliesWatcherService = (*AssembliesWatcherServiceImpl)(nil)

type AssembliesWatcherServiceImpl struct {
	assembliesConsumer platform_kafka.Consumer
	notificationSender model_notification_sender.NotificationSender
	serviceConsumer    *AssembliesWatcherConsumer
}

func NewAssembliesWatcherServiceImpl(
	assembliesConsumer platform_kafka.Consumer,
	notificationSender model_notification_sender.NotificationSender,
) *AssembliesWatcherServiceImpl {
	service := AssembliesWatcherServiceImpl{
		assembliesConsumer: assembliesConsumer,
		notificationSender: notificationSender,
	}

	serviceConsumer := NewAssembliesWatcherConsumer(
		assembliesConsumer,
		service.handleAssemblyCompleted,
	)
	service.serviceConsumer = serviceConsumer

	return &service
}
