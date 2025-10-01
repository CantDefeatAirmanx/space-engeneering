package service_ship_assembly

func (s *ShipAssemblyServiceImpl) Close() error {
	s.cancel()
	return nil
}
