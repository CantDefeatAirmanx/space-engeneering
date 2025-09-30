package consumer_ship_assembly

func (s *ShipAssemblyConsumerImpl) Close() error {
	if err := s.consumer.Close(); err != nil {
		return err
	}
	return nil
}
