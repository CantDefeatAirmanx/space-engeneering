package platform_kafka_client

func (kc *KafkaClient) GetTopics() (map[string]TopicConfig, error) {
	topics, err := kc.admin.ListTopics()
	if err != nil {
		return nil, err
	}

	topicsConfig := make(map[string]TopicConfig)
	for topicName, config := range topics {
		topicsConfig[topicName] = TopicConfig{
			Topic:             topicName,
			NumPartitions:     int(config.NumPartitions),
			ReplicationFactor: int(config.ReplicationFactor),
		}
	}

	return topicsConfig, nil
}
