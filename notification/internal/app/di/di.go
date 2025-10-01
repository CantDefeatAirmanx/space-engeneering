package di

import (
	"context"

	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/notification/config"
	model_notification_sender "github.com/CantDefeatAirmanx/space-engeneering/notification/internal/model/notification_sender"
	service_assemblies_watcher "github.com/CantDefeatAirmanx/space-engeneering/notification/internal/service/assemblies_watcher"
	service_orders_watcher "github.com/CantDefeatAirmanx/space-engeneering/notification/internal/service/orders_watcher"
	"github.com/CantDefeatAirmanx/space-engeneering/notification/internal/service/telegram_notification_sender"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	platform_kafka_consumer "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/consumer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	platform_telegram "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/telegram"
)

type DiContainer struct {
	closer      closer.Closer
	telegramBot platform_telegram.TelegramClient

	assembliesWatcherService     service_assemblies_watcher.AssembliesWatcherService
	assembliesNotificationSender model_notification_sender.NotificationSender
	assembliesConsumer           platform_kafka.Consumer
	ordersWatcherService         service_orders_watcher.OrdersWatcherService
	ordersNotificationSender     model_notification_sender.NotificationSender
	ordersConsumer               platform_kafka.Consumer
}

func NewDiContainer(closer closer.Closer) *DiContainer {
	return &DiContainer{
		closer: closer,
	}
}

func (d *DiContainer) GetAssembliesWatcherService(ctx context.Context) service_assemblies_watcher.AssembliesWatcherService {
	if d.assembliesWatcherService != nil {
		return d.assembliesWatcherService
	}

	assembliesWatcherService := service_assemblies_watcher.NewAssembliesWatcherServiceImpl(
		d.GetAssembliesConsumer(ctx),
		d.GetAssembliesNotificationSender(ctx),
	)
	d.assembliesWatcherService = assembliesWatcherService

	return assembliesWatcherService
}

func (d *DiContainer) GetOrdersWatcherService(ctx context.Context) service_orders_watcher.OrdersWatcherService {
	if d.ordersWatcherService != nil {
		return d.ordersWatcherService
	}

	ordersWatcherService := service_orders_watcher.NewOrdersWatcherServiceImpl(
		d.GetOrdersConsumer(ctx),
		d.GetOrdersNotificationSender(ctx),
	)
	d.ordersWatcherService = ordersWatcherService

	return ordersWatcherService
}

func (d *DiContainer) GetAssembliesNotificationSender(ctx context.Context) model_notification_sender.NotificationSender {
	if d.assembliesNotificationSender != nil {
		return d.assembliesNotificationSender
	}

	assembliesNotificationSender := telegram_notification_sender.NewTelegramNotificationSender(
		d.GetTelegramBot(ctx),
		config.Config.Telegram().AssembliesNotificationsChatId(),
		platform_telegram.WithThreadId(config.Config.Telegram().AssembliesNotificationsThreadId()),
	)
	d.assembliesNotificationSender = assembliesNotificationSender

	return d.assembliesNotificationSender
}

func (d *DiContainer) GetOrdersNotificationSender(ctx context.Context) model_notification_sender.NotificationSender {
	if d.ordersNotificationSender != nil {
		return d.ordersNotificationSender
	}

	ordersNotificationSender := telegram_notification_sender.NewTelegramNotificationSender(
		d.GetTelegramBot(ctx),
		config.Config.Telegram().OrdersNotificationsChatId(),
		platform_telegram.WithThreadId(config.Config.Telegram().OrdersNotificationsThreadId()),
	)
	d.ordersNotificationSender = ordersNotificationSender

	return d.ordersNotificationSender
}

func (d *DiContainer) GetAssembliesConsumer(ctx context.Context) platform_kafka.Consumer {
	if d.assembliesConsumer != nil {
		return d.assembliesConsumer
	}

	cfg := config.Config
	_ = cfg

	assembliesConsumer, err := platform_kafka_consumer.NewKafkaConsumer(
		config.Config.Kafka().Brokers(),
		"Assemblies Consumer",
	)
	d.closer.AddNamed("Assemblies Consumer", func(ctx context.Context) error {
		return assembliesConsumer.Close()
	})
	if err != nil {
		logger.Logger().Error("Failed to create kafka consumer", zap.Error(err))
		panic(err)
	}

	d.assembliesConsumer = assembliesConsumer

	return d.assembliesConsumer
}

func (d *DiContainer) GetOrdersConsumer(ctx context.Context) platform_kafka.Consumer {
	if d.ordersConsumer != nil {
		return d.ordersConsumer
	}

	ordersConsumer, err := platform_kafka_consumer.NewKafkaConsumer(
		config.Config.Kafka().Brokers(),
		"Orders Consumer",
	)
	if err != nil {
		logger.Logger().Error("Failed to create kafka consumer", zap.Error(err))
		panic(err)
	}
	d.closer.AddNamed("Orders Consumer", func(ctx context.Context) error {
		return ordersConsumer.Close()
	})

	d.ordersConsumer = ordersConsumer

	return d.ordersConsumer
}

func (d *DiContainer) GetTelegramBot(ctx context.Context) platform_telegram.TelegramClient {
	if d.telegramBot != nil {
		return d.telegramBot
	}

	telegramBot, err := platform_telegram.NewTelegramClient(
		config.Config.Telegram().BotToken(),
	)
	if err != nil {
		logger.Logger().Error("Failed to create telegram bot", zap.Error(err))
		panic(err)
	}

	d.telegramBot = telegramBot

	return d.telegramBot
}
