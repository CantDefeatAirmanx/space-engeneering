package platform_kafka_client

func (kc *KafkaClient) Close() error {
	if kc.admin != nil {
		kc.admin.Close()
	}
	if kc.client != nil {
		return kc.client.Close()
	}
	return nil
}
