package service_ship_assembly_consumer

func (s *ShipAssemblyConsumerImpl) Close() error {
	if err := s.consumer.Close(); err != nil {
		return err
	}
	return nil
}
