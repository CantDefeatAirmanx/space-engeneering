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
) (*model_order.Order, error) {
	repoOrder, err := ToRepository(&order)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", model_order.ErrOrderInvalidArguments, err)
	}

	columns := []string{
		columnUserUUID,
		columnPartsUUIDs,
		columnTotalPrice,
		columnTransactionUUID,
		columnPaymentMethod,
		columnStatus,
	}
	args := []any{
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
		Suffix("RETURNING " + columnOrderUUID + ", " + columnCreatedAt + ", " + columnUpdatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", model_order.ErrOrderInternal, err)
	}

	createdOrder := Order{
		UserUUID:        repoOrder.UserUUID,
		PartUuids:       repoOrder.PartUuids,
		TotalPrice:      repoOrder.TotalPrice,
		TransactionUUID: repoOrder.TransactionUUID,
		PaymentMethod:   repoOrder.PaymentMethod,
		Status:          repoOrder.Status,
	}
	err = o.db.QueryRow(ctx, query, args...).Scan(
		&createdOrder.OrderUUID,
		&createdOrder.CreatedAt,
		&createdOrder.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", model_order.ErrOrderInternal, err)
	}

	modelOrder, err := ToModel(createdOrder)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", model_order.ErrOrderInternal, err)
	}

	return modelOrder, nil
}
