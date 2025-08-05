package repository_order_postgre

import (
	"context"

	"github.com/Masterminds/squirrel"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (o *OrderRepositoryPostgre) GetOrder(
	ctx context.Context,
	orderUUID string,
) (*model_order.Order, error) {
	query, _, err := squirrel.Select(
		columnOrderUUID,
		columnUserUUID,
		columnPartsUUIDs,
		columnTotalPrice,
		columnTransactionUUID,
		columnPaymentMethod,
		columnStatus,
	).From("orders").Where(
		squirrel.Eq{columnOrderUUID: orderUUID},
	).ToSql()
	if err != nil {
		return nil, err
	}

	var order model_order.Order
	row := o.db.QueryRow(ctx, query, orderUUID)

	err = row.Scan(
		&order.OrderUUID,
		&order.UserUUID,
		&order.PartUuids,
		&order.TotalPrice,
		&order.TransactionUUID,
		&order.TransactionUUID,
		&order.PaymentMethod,
		&order.Status,
	)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
