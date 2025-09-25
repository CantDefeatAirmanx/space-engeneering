package service_assemblies_watcher

import (
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	platform_telegram "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/telegram"
)

var _ AssembliesWatcherService = (*AssembliesWatcherServiceImpl)(nil)

type AssembliesWatcherServiceImpl struct {
	assembliesConsumer platform_kafka.Consumer
	telegramClient     platform_telegram.TelegramClient
	serviceConsumer    *AssembliesWatcherConsumer
}

func NewAssembliesWatcherServiceImpl(
	assembliesConsumer platform_kafka.Consumer,
	telegramClient platform_telegram.TelegramClient,
) *AssembliesWatcherServiceImpl {
	service := AssembliesWatcherServiceImpl{
		assembliesConsumer: assembliesConsumer,
		telegramClient:     telegramClient,
	}

	serviceConsumer := NewAssembliesWatcherConsumer(
		assembliesConsumer,
		service.handleAssemblyCompleted,
	)
	service.serviceConsumer = serviceConsumer

	return &service
}
