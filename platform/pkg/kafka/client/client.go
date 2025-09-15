package platform_kafka_client

import (
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	"github.com/IBM/sarama"
)

type KafkaClient struct {
	client sarama.Client
	admin  sarama.ClusterAdmin
}

func NewKafkaClient(brokers []string) (*KafkaClient, error) {
	cfg := sarama.NewConfig()

	cfg.Version = platform_kafka.KafkaVersion

	client, err := sarama.NewClient(brokers, cfg)
	if err != nil {
		return nil, err
	}

	admin, err := sarama.NewClusterAdmin(brokers, cfg)
	if err != nil {
		client.Close()
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
