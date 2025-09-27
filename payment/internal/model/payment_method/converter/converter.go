package model_payment_method_converter

import (
	model_payment_method "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/model/payment_method"
	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
)

func ToModel(proto common_v1.PaymentMethod) model_payment_method.PaymentMethod {
	return model_payment_method.PaymentMethod(proto)
}

func ToProto(model model_payment_method.PaymentMethod) common_v1.PaymentMethod {
	return common_v1.PaymentMethod(model)
}
