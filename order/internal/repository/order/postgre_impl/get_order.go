package repository_order_postgre

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (o *OrderRepositoryPostgre) GetOrder(
	ctx context.Context,
	orderUUID string,
) (*model_order.Order, error) {
	query, _, err := squirrel.
		Select(
			columnOrderUUID,
			columnUserUUID,
			columnPartsUUIDs,
			columnTotalPrice,
			columnTransactionUUID,
			columnPaymentMethod,
			columnStatus,
			columnCreatedAt,
			columnUpdatedAt,
		).
		From(tableOrders).
		Where(squirrel.Eq{columnOrderUUID: orderUUID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var repoOrder Order
	err = o.db.QueryRow(ctx, query, orderUUID).Scan(
		&repoOrder.OrderUUID,
		&repoOrder.UserUUID,
		&repoOrder.PartUuids,
		&repoOrder.TotalPrice,
		&repoOrder.TransactionUUID,
		&repoOrder.PaymentMethod,
		&repoOrder.Status,
		&repoOrder.CreatedAt,
		&repoOrder.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model_order.ErrOrderNotFound
		}
		return nil, err
	}

	modelOrder, err := ToModel(repoOrder)
	if err != nil {
		return nil, err
	}

	return modelOrder, nil
}
