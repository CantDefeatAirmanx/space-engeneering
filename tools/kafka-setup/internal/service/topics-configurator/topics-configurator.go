package service_topics_configurator

import platform_kafka_client "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/client"

type TopicsConfigurator struct {
	client         platform_kafka_client.KafkaClientInterface
	yamlConfigPath string
}

func NewTopicsConfigurator(
	client platform_kafka_client.KafkaClientInterface,
	yamlConfigPath string,
) *TopicsConfigurator {
	return &TopicsConfigurator{
		client:         client,
		yamlConfigPath: yamlConfigPath,
	}
}
