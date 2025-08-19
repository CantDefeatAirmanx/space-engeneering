package repository_order_postgre

import (
	"context"

	"github.com/Masterminds/squirrel"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (o *OrderRepositoryPostgre) DeleteOrder(
	ctx context.Context,
	orderUUID string,
) error {
	query, args, err := squirrel.
		Delete(tableOrders).
		Where(squirrel.Eq{columnOrderUUID: orderUUID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	result, err := o.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return model_order.ErrOrderNotFound
	}

	return nil
}
