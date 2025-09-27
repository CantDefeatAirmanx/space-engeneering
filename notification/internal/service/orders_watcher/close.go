package service_orders_watcher

func (o *OrdersWatcherServiceImpl) Close() error {
	if err := o.ordersConsumer.Close(); err != nil {
		return err
	}
	return nil
}
