package platform_kafka_client

func (kc *KafkaClient) Close() error {
	if kc.admin != nil {
		err := kc.admin.Close()
		if err != nil {
			return err
		}
	}
	if kc.client != nil {
		err := kc.client.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
