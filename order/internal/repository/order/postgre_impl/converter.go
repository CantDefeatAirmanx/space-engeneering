package repository_order_postgre

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func ToModel(order Order) (*model_order.Order, error) {
	partUuids := lo.Map(order.PartUuids, func(partUuid pgtype.UUID, _ int) string {
		return partUuid.String()
	})
	totalPrice, err := order.TotalPrice.Float64Value()
	if err != nil {
		return nil, err
	}
	transactionUUID := order.TransactionUUID.String()
	paymentMethod := model_order.PaymentMethod(*order.PaymentMethod)

	return &model_order.Order{
		OrderUUID:       order.OrderUUID.String(),
		UserUUID:        order.UserUUID.String(),
		PartUuids:       partUuids,
		TotalPrice:      totalPrice.Float64,
		TransactionUUID: &transactionUUID,
		PaymentMethod:   &paymentMethod,
		Status:          model_order.OrderStatus(order.Status),
	}, nil
}

func ToRepository(order *model_order.Order) (*Order, error) {
	result := Order{
		Status: OrderStatus(order.Status),
	}

	var orderUUID pgtype.UUID
	if err := orderUUID.Scan(order.OrderUUID); err != nil {
		return nil, err
	}
	result.OrderUUID = orderUUID

	var userUUID pgtype.UUID
	if err := userUUID.Scan(order.UserUUID); err != nil {
		return nil, err
	}
	result.UserUUID = userUUID

	partUuids := make([]pgtype.UUID, len(order.PartUuids))
	for _, partUUID := range order.PartUuids {
		var pgPartUUID pgtype.UUID
		if err := pgPartUUID.Scan(partUUID); err != nil {
			return nil, err
		}
		partUuids = append(partUuids, pgPartUUID)
	}
	result.PartUuids = partUuids

	totalPrice := pgtype.Numeric{}
	totalPriceStr := strconv.FormatFloat(order.TotalPrice, 'f', 2, 64)
	if err := totalPrice.Scan(totalPriceStr); err != nil {
		return nil, err
	}
	result.TotalPrice = totalPrice

	var transactionUUID *pgtype.UUID
	if order.TransactionUUID != nil {
		transactionUUID = &pgtype.UUID{}

		if err := transactionUUID.Scan(*order.TransactionUUID); err != nil {
			return nil, err
		}
	}
	result.TransactionUUID = transactionUUID

	if order.PaymentMethod != nil {
		converted := PaymentMethod(*order.PaymentMethod)
		result.PaymentMethod = &converted
	}

	return &result, nil
}
