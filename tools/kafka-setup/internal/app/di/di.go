package di

import (
	"path/filepath"

	platform_kafka_client "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/client"
	"github.com/CantDefeatAirmanx/space-engeneering/tools/kafka-setup/config"
	service_topics_configurator "github.com/CantDefeatAirmanx/space-engeneering/tools/kafka-setup/internal/service/topics-configurator"
)

type DiContainer struct {
	topicsConfigurator *service_topics_configurator.TopicsConfigurator
	kafkaClient        platform_kafka_client.KafkaClientInterface
}

func NewDiContainer() *DiContainer {
	return &DiContainer{}
}

func (d *DiContainer) GetTopicsConfigurator() *service_topics_configurator.TopicsConfigurator {
	if d.topicsConfigurator != nil {
		return d.topicsConfigurator
	}

	topicsConfigurator := service_topics_configurator.NewTopicsConfigurator(
		d.GetKafkaClient(),
		filepath.Join("configs", "kafka", "kafka.yaml"),
	)
	d.topicsConfigurator = topicsConfigurator

	return topicsConfigurator
}

func (d *DiContainer) GetKafkaClient() platform_kafka_client.KafkaClientInterface {
	if d.kafkaClient != nil {
		return d.kafkaClient
	}

	kafkaClient, err := platform_kafka_client.NewKafkaClient(
		config.Config.Kafka().Brokers(),
	)
	if err != nil {
		panic(err)
	}
	d.kafkaClient = kafkaClient

	return d.kafkaClient
}
