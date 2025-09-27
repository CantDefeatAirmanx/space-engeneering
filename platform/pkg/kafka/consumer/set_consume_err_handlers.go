package platform_kafka_consumer

func (k *KafkaConsumerImpl) SetKafkaErrorsHandlers(errHandlers []func(err error)) {
	k.kafkaErrorsHandlers = errHandlers
}

func (k *KafkaConsumerImpl) SetProcessMessageErrHandlers(errHandlers []func(err error)) {
	k.processMessageErrHandlers = errHandlers
}
