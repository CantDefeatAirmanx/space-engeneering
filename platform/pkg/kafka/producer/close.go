package platform_kafka_producer

func (p *KafkaProducerImpl) Close() error {
	p.cancel()
	p.asyncProducer.AsyncClose()

	err := p.syncProducer.Close()
	if err != nil {
		return err
	}

	return nil
}
