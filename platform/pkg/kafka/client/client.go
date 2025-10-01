package platform_kafka_client

import (
	"errors"

	"github.com/IBM/sarama"

	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

type KafkaClient struct {
	client sarama.Client
	admin  sarama.ClusterAdmin
}

func NewKafkaClient(brokers []string) (KafkaClientInterface, error) {
	cfg := sarama.NewConfig()

	cfg.Version = platform_kafka.KafkaVersion

	client, err := sarama.NewClient(brokers, cfg)
	if err != nil {
		return nil, err
	}

	admin, err := sarama.NewClusterAdmin(brokers, cfg)
	if err != nil {
		closeErr := client.Close()
		if closeErr != nil {
			return nil, errors.Join(err, closeErr)
		}
		return nil, err
	}

	return &KafkaClient{
		client: client,
		admin:  admin,
	}, nil
}

type TopicConfig struct {
	Topic             string
	NumPartitions     int
	ReplicationFactor int
	ConfigEntries     []ConfigEntry
}

type ConfigEntry struct {
	ConfigName  string
	ConfigValue string
}
