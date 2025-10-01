package service_orders_watcher

import "context"

func (o *OrdersWatcherServiceImpl) WatchOrders(ctx context.Context) error {
	return o.serviceConsumer.ConsumeOrders(ctx)
}
