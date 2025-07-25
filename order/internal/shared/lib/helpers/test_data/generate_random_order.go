package helpers_test_data

import (
	"github.com/brianvoe/gofakeit/v7"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func GenerateRandomOrder() *model_order.Order {
	transactionUUID := gofakeit.UUID()

	paymentMethod := model_order.PaymentMethod(gofakeit.RandomString([]string{
		string(model_order.PaymentMethodUnknown),
		string(model_order.PaymentMethodCard),
		string(model_order.PaymentMethodSBP),
		string(model_order.PaymentMethodCreditCard),
		string(model_order.PaymentMethodInvestorMoney),
	}))

	result := &model_order.Order{
		OrderUUID: gofakeit.UUID(),
		UserUUID:  gofakeit.UUID(),
		Status: model_order.OrderStatus(gofakeit.RandomString([]string{
			string(model_order.OrderStatusPendingPayment),
			string(model_order.OrderStatusPaid),
			string(model_order.OrderStatusCancelled),
		})),
		TotalPrice:      gofakeit.Float64Range(100, 1000),
		TransactionUUID: &transactionUUID,
		PaymentMethod:   &paymentMethod,
	}

	randPartCount := gofakeit.IntRange(1, 10)
	partUuids := make([]string, randPartCount)
	for idx := range partUuids {
		partUuids[idx] = gofakeit.UUID()
	}
	result.PartUuids = partUuids[:randPartCount]

	return result
}
