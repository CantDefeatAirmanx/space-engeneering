package service_orders_watcher

import (
	"context"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
)

type OrdersWatcherService interface {
	WatchOrders(ctx context.Context) error
	interfaces.WithClose
}
