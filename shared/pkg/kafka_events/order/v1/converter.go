package kafka_events_order

import (
	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
	"github.com/gogo/protobuf/proto"
)

var paymentMethodsToModel = map[common_v1.PaymentMethod]PaymentMethod{
	common_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED: PaymentMethodUnknown,
	common_v1.PaymentMethod_PAYMENT_METHOD_CARD:                PaymentMethodCard,
	common_v1.PaymentMethod_PAYMENT_METHOD_SPB:                 PaymentMethodSBP,
	common_v1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD:         PaymentMethodCreditCard,
	common_v1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY:      PaymentMethodInvestorMoney,
}

var paymentMethodsToProto = map[PaymentMethod]common_v1.PaymentMethod{
	PaymentMethodUnknown:       common_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED,
	PaymentMethodCard:          common_v1.PaymentMethod_PAYMENT_METHOD_CARD,
	PaymentMethodSBP:           common_v1.PaymentMethod_PAYMENT_METHOD_SPB,
	PaymentMethodCreditCard:    common_v1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD,
	PaymentMethodInvestorMoney: common_v1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY,
}

func DecodeToOrderEnvelope(message []byte) (*order_events_v1.OrderEventEnvelope, error) {
	var protoMessage order_events_v1.OrderEventEnvelope
	if err := proto.Unmarshal(message, &protoMessage); err != nil {
		return nil, err
	}
	return &protoMessage, nil
}

func ConvertOrderPaidProtoToModel(eventProto *order_events_v1.OrderPaidEvent) OrderPaidEvent {
	return OrderPaidEvent{
		EventUUID:     eventProto.EventUuid,
		OrderUUID:     eventProto.OrderUuid,
		UserUUID:      eventProto.UserUuid,
		PaymentMethod: paymentMethodsToModel[eventProto.PaymentMethod],
	}
}

func ConvertOrderPaidModelToProto(eventModel *OrderPaidEvent) order_events_v1.OrderPaidEvent {
	return order_events_v1.OrderPaidEvent{
		EventUuid:     eventModel.EventUUID,
		OrderUuid:     eventModel.OrderUUID,
		UserUuid:      eventModel.UserUUID,
		PaymentMethod: paymentMethodsToProto[eventModel.PaymentMethod],
	}
}
