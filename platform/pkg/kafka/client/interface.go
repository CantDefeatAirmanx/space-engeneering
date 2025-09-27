package platform_kafka_client

import (
	"context"

	"github.com/IBM/sarama"
)

type KafkaClientInterface interface {
	Close() error
	CreateTopic(ctx context.Context, topicConfig TopicConfig) error
	GetTopics() (map[string]TopicConfig, error)
	DescribeCluster() (brokers []*sarama.Broker, controllerID int32, err error)
}
