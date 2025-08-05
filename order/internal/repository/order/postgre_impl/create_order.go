package repository_order_postgre

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (o *OrderRepositoryPostgre) CreateOrder(
	ctx context.Context,
	order model_order.Order,
) error {
	repoOrder, err := ToRepository(&order)
	if err != nil {
		return err
	}

	columns := []string{
		columnUserUUID,
		columnPartsUUIDs,
		columnTotalPrice,
		columnTransactionUUID,
		columnPaymentMethod,
		columnStatus,
	}
	args := []interface{}{
		repoOrder.UserUUID,
		repoOrder.PartUuids,
		repoOrder.TotalPrice,
		repoOrder.TransactionUUID,
		repoOrder.PaymentMethod,
		repoOrder.Status,
	}

	query, _, err := squirrel.Insert("orders").
		Columns(columns...).
		Values(args...).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	res, err := o.db.Exec(ctx, query, args...)
	fmt.Println(res)

	return err
}
