package model_order_converter

import (
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func ToModel(apiOrder *order_v1.Order) model_order.Order {
	result := model_order.Order{
		OrderUUID:  apiOrder.OrderUUID,
		UserUUID:   apiOrder.UserUUID,
		PartUuids:  apiOrder.PartUuids,
		TotalPrice: apiOrder.TotalPrice,
		Status:     model_order.OrderStatus(apiOrder.GetStatus()),
	}

	if apiOrder.GetTransactionUUID().Set {
		transactionUUID := apiOrder.GetTransactionUUID().Value
		result.TransactionUUID = &transactionUUID
	}

	if apiOrder.GetPaymentMethod().Set {
		paymentMethod := model_order.PaymentMethod(apiOrder.GetPaymentMethod().Value)
		result.PaymentMethod = &paymentMethod
	}

	return result
}

func ToApi(model *model_order.Order) order_v1.Order {
	result := order_v1.Order{
		OrderUUID:  model.OrderUUID,
		UserUUID:   model.UserUUID,
		PartUuids:  model.PartUuids,
		TotalPrice: model.TotalPrice,
		Status:     order_v1.OrderStatus(model.Status),
	}

	if model.PaymentMethod != nil {
		paymentMethod := order_v1.PaymentMethod(*model.PaymentMethod)
		result.PaymentMethod.Value = paymentMethod
		result.PaymentMethod.Set = true
	}

	if model.TransactionUUID != nil {
		result.TransactionUUID.Value = *model.TransactionUUID
		result.TransactionUUID.Set = true
	}

	return result
}
