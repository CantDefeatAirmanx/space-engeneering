package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	platform_kafka_client "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/client"
	platform_logger "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"github.com/CantDefeatAirmanx/space-engeneering/tools/kafka-setup/config"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var (
	logger = platform_logger.DefaultInfoLogger()
)

func main() {
	ctx := context.Background()

	err := config.LoadConfig(
		config.WithEnvPath(filepath.Join("tools", "kafka-setup", ".env")),
	)
	if err != nil {
		logger.Error("Ошибка загрузки конфигурации", zap.Error(err))

		return
	}

	client, err := platform_kafka_client.NewKafkaClient(
		config.Config.Kafka.Brokers,
	)
	if err != nil {
		log.Fatalf("Ошибка создания клиента: %v", err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("Ошибка закрытия клиента: %v", err)
		}
	}()

	topicsPath := filepath.Join("configs", "kafka", "kafka.yaml")
	yamlFile, err := os.ReadFile(topicsPath)

	if err != nil {
		log.Fatalf("Ошибка чтения файла топиков: %v", err)
	}

	var kafkaConfig config.KafkaConfigYaml
	err = yaml.Unmarshal(yamlFile, &kafkaConfig)
	if err != nil {
		log.Fatalf("Ошибка парсинга конфигурации: %v", err)
	}

	log.Printf("Создание %d топиков...", len(kafkaConfig.Topics))

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

		err = client.CreateTopic(ctx, topicConfig)
		if err != nil {
			log.Printf("Ошибка создания топика %s: %v", topicConf.Name, err)
			continue
		}

		log.Printf("Топик %s успешно создан", topicConf.Name)
	}

	log.Println("Настройка топиков завершена")
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
