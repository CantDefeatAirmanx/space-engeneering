package model_order

type UpdateOrderFields struct {
	Status          *OrderStatus
	TransactionUUID *string
	PaymentMethod   *PaymentMethod
}
