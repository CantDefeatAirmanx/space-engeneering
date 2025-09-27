package kafka_events_order

type OrderPaidEvent struct {
	EventUUID     string
	OrderUUID     string
	UserUUID      string
	PaymentMethod PaymentMethod
}
