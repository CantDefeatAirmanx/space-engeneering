package repository_order_map

import (
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func ToModel(repo *Order) model_order.Order {
	result := model_order.Order{
		OrderUUID:       repo.OrderUUID,
		UserUUID:        repo.UserUUID,
		PartUuids:       repo.PartUuids,
		TotalPrice:      repo.TotalPrice,
		TransactionUUID: repo.TransactionUUID,
		Status:          model_order.OrderStatus(repo.Status),
	}

	if repo.PaymentMethod != nil {
		paymentMethod := model_order.PaymentMethod(*repo.PaymentMethod)
		result.PaymentMethod = &paymentMethod
	}

	return result
}

func ToRepository(model *model_order.Order) Order {
	result := Order{
		OrderUUID:       model.OrderUUID,
		UserUUID:        model.UserUUID,
		PartUuids:       model.PartUuids,
		TotalPrice:      model.TotalPrice,
		TransactionUUID: model.TransactionUUID,
		Status:          OrderStatus(model.Status),
	}

	if model.PaymentMethod != nil {
		paymentMethod := PaymentMethod(*model.PaymentMethod)
		result.PaymentMethod = &paymentMethod
	}

	return result
}
