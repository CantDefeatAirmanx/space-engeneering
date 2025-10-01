package platform_kafka_consumer

func (k *KafkaConsumerImpl) Close() error {
	return k.consumerGroup.Close()
}
