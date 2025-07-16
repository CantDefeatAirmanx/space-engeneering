package model_payment_method_converter

import (
	model_payment_method "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/model/payment_method"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

func ToModel(proto payment_v1.PaymentMethod) model_payment_method.PaymentMethod {
	return model_payment_method.PaymentMethod(proto)
}

func ToProto(model model_payment_method.PaymentMethod) payment_v1.PaymentMethod {
	return payment_v1.PaymentMethod(model)
}
