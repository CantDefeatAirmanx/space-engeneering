package platform_kafka_client

import (
	"context"

	"github.com/IBM/sarama"
)

func (kc *KafkaClient) CreateTopic(ctx context.Context, topicConfig TopicConfig) error {
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     int32(topicConfig.NumPartitions),
		ReplicationFactor: int16(topicConfig.ReplicationFactor),
		ConfigEntries:     make(map[string]*string),
	}

	for _, entry := range topicConfig.ConfigEntries {
		topicDetail.ConfigEntries[entry.ConfigName] = &entry.ConfigValue
	}

	return kc.admin.CreateTopic(topicConfig.Topic, topicDetail, false)
}
