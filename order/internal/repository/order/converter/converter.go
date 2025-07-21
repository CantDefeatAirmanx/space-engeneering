package repository_order_converter

import (
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

func ToModel(repo *repository_order_model.Order) model_order.Order {
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

func ToRepository(model *model_order.Order) repository_order_model.Order {
	result := repository_order_model.Order{
		OrderUUID:       model.OrderUUID,
		UserUUID:        model.UserUUID,
		PartUuids:       model.PartUuids,
		TotalPrice:      model.TotalPrice,
		TransactionUUID: model.TransactionUUID,
		Status:          repository_order_model.OrderStatus(model.Status),
	}

	if model.PaymentMethod != nil {
		paymentMethod := repository_order_model.PaymentMethod(*model.PaymentMethod)
		result.PaymentMethod = &paymentMethod
	}

	return result
}
