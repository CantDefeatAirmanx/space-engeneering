package service_order_producer

import (
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
)

var _ service_order.OrderProducer = (*OrderProducerImpl)(nil)

type OrderProducerImpl struct {
	producer platform_kafka.Producer
}

func NewOrderProducer(producer platform_kafka.Producer) *OrderProducerImpl {
	return &OrderProducerImpl{
		producer: producer,
	}
}
