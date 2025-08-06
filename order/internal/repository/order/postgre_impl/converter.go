package repository_order_postgre

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func ToModel(order Order) (*model_order.Order, error) {
	result := model_order.Order{
		OrderUUID: order.OrderUUID.String(),
		UserUUID:  order.UserUUID.String(),
		Status:    model_order.OrderStatus(order.Status),
	}

	partUuids := lo.Map(order.PartUuids, func(partUuid pgtype.UUID, _ int) string {
		return partUuid.String()
	})
	result.PartUuids = partUuids

	totalPrice, err := order.TotalPrice.Float64Value()
	if err != nil {
		return nil, err
	}
	result.TotalPrice = totalPrice.Float64

	if order.TransactionUUID != nil {
		transactionUUID := order.TransactionUUID.String()
		result.TransactionUUID = &transactionUUID
	}
	if order.PaymentMethod != nil {
		paymentMethod := model_order.PaymentMethod(*order.PaymentMethod)
		result.PaymentMethod = &paymentMethod
	}

	return &result, nil
}

func ToRepository(order *model_order.Order) (*Order, error) {
	result := Order{
		Status: OrderStatus(order.Status),
	}

	var userUUID pgtype.UUID
	if err := userUUID.Scan(order.UserUUID); err != nil {
		return nil, err
	}
	result.UserUUID = userUUID

	partUuids := make([]pgtype.UUID, 0, len(order.PartUuids))
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

func ToRepositoryUpdate(update model_order.UpdateOrderFields) (*UpdateOrderFields, error) {
	result := UpdateOrderFields{}

	if update.Status != nil {
		status := OrderStatus(*update.Status)
		result.Status = &status
	}

	if update.TransactionUUID != nil {
		transactionUUID := pgtype.UUID{}
		if err := transactionUUID.Scan(*update.TransactionUUID); err != nil {
			return nil, err
		}
		result.TransactionUUID = &transactionUUID
	}

	if update.PaymentMethod != nil {
		paymentMethod := PaymentMethod(*update.PaymentMethod)
		result.PaymentMethod = &paymentMethod
	}

	return &result, nil
}
