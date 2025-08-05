package repository_order_postgre

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (o *OrderRepositoryPostgre) UpdateOrderFields(
	ctx context.Context,
	orderUUID string,
	update model_order.UpdateOrderFields,
) error {
	panic("unimplemented")
}
