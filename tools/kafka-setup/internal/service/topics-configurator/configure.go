package service_topics_configurator

import (
	"context"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"go.uber.org/zap"

	platform_kafka_client "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/client"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

func (tc *TopicsConfigurator) Configure(
	ctx context.Context,
) error {
	yamlFile, err := os.ReadFile(tc.yamlConfigPath)
	if err != nil {
		logger.Logger().Error("Error while reading topics file", zap.Error(err))
		return err
	}

	var kafkaConfig KafkaConfigYaml
	err = yaml.Unmarshal(yamlFile, &kafkaConfig)
	if err != nil {
		logger.Logger().Error("Error while parsing configuration", zap.Error(err))
		return err
	}

	logger.Logger().Info("Creating topics...")

	for _, topicConf := range kafkaConfig.Topics {
		topicConfig := platform_kafka_client.TopicConfig{
			Topic:             topicConf.Name,
			NumPartitions:     getPartitions(topicConf.Partitions, kafkaConfig.DefaultTopicConfig.Partitions),
			ReplicationFactor: getReplicationFactor(topicConf.ReplicationFactor, kafkaConfig.DefaultTopicConfig.ReplicationFactor),
			ConfigEntries: []platform_kafka_client.ConfigEntry{
				{
					ConfigName:  "min.insync.replicas",
					ConfigValue: fmt.Sprintf("%d", kafkaConfig.DefaultTopicConfig.MinInSyncReplicas),
				},
				{
					ConfigName:  "retention.ms",
					ConfigValue: fmt.Sprintf("%d", kafkaConfig.DefaultTopicConfig.RetentionMs),
				},
				{
					ConfigName:  "cleanup.policy",
					ConfigValue: kafkaConfig.DefaultTopicConfig.CleanupPolicy,
				},
			},
		}

		err = tc.client.CreateTopic(ctx, topicConfig)
		if err != nil {
			logger.Logger().Error("Error while creating topic", zap.String("topic", topicConf.Name), zap.Error(err))
			continue
		}

		logger.Logger().Info("Topic successfully created", zap.String("topic", topicConf.Name))
	}

	logger.Logger().Info("Topics configuration completed")

	return nil
}

func getPartitions(topicPartitions, defaultPartitions int) int {
	if topicPartitions > 0 {
		return topicPartitions
	}
	return defaultPartitions
}

func getReplicationFactor(topicReplication, defaultReplication int) int {
	if topicReplication > 0 {
		return topicReplication
	}
	return defaultReplication
}
