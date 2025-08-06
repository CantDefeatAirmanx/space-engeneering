package repository_order_postgre

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (o *OrderRepositoryPostgre) UpdateOrderFields(
	ctx context.Context,
	orderUUID string,
	update model_order.UpdateOrderFields,
) error {
	repoUpdate, err := ToRepositoryUpdate(update)
	if err != nil {
		return fmt.Errorf("%w: %v", model_order.ErrOrderInvalidArguments, err)
	}

	query, args, err := squirrel.Update(tableOrders).
		Set(columnStatus, repoUpdate.Status).
		Set(columnTransactionUUID, repoUpdate.TransactionUUID).
		Set(columnPaymentMethod, repoUpdate.PaymentMethod).
		Where(squirrel.Eq{columnOrderUUID: orderUUID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w: %v", model_order.ErrOrderInternal, err)
	}

	result, err := o.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %v", model_order.ErrOrderInternal, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return model_order.ErrOrderNotFound
	}

	return nil
}
