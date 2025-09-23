package service_topics_configurator

import (
	"context"
	"fmt"
	"os"

	"github.com/IBM/sarama"
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

	logger.Logger().Info("Read YAML file",
		zap.String("path", tc.yamlConfigPath),
		zap.String("content", string(yamlFile)))

	var kafkaConfig KafkaConfigYaml
	err = yaml.Unmarshal(yamlFile, &kafkaConfig)
	if err != nil {
		logger.Logger().Error("Error while parsing configuration", zap.Error(err))
		return err
	}

	logger.Logger().Info("Parsed configuration",
		zap.Int("default_partitions", kafkaConfig.DefaultTopicConfig.Partitions),
		zap.Int("default_replication_factor", kafkaConfig.DefaultTopicConfig.ReplicationFactor),
		zap.Int("default_min_insync_replicas", kafkaConfig.DefaultTopicConfig.MinInSyncReplicas),
		zap.Int64("default_retention_ms", kafkaConfig.DefaultTopicConfig.RetentionMs),
		zap.String("default_cleanup_policy", kafkaConfig.DefaultTopicConfig.CleanupPolicy),
		zap.Int("topics_count", len(kafkaConfig.Topics)))

	logger.Logger().Info("Creating topics...")
	defaultConfig := kafkaConfig.DefaultTopicConfig

	for _, topicConf := range kafkaConfig.Topics {
		topicConfig := platform_kafka_client.TopicConfig{
			Topic:             topicConf.Name,
			NumPartitions:     getPartitions(topicConf.Partitions, defaultConfig.Partitions),
			ReplicationFactor: getReplicationFactor(topicConf.ReplicationFactor, defaultConfig.ReplicationFactor),
			ConfigEntries: []platform_kafka_client.ConfigEntry{
				{
					ConfigName:  "min.insync.replicas",
					ConfigValue: fmt.Sprintf("%d", defaultConfig.MinInSyncReplicas),
				},
				{
					ConfigName:  "retention.ms",
					ConfigValue: fmt.Sprintf("%d", defaultConfig.RetentionMs),
				},
				{
					ConfigName:  "cleanup.policy",
					ConfigValue: defaultConfig.CleanupPolicy,
				},
			},
		}

		err = tc.client.CreateTopic(ctx, topicConfig)
		if err != nil && err != sarama.ErrTopicAlreadyExists {
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
