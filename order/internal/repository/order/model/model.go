package repository_order_model

type Order struct {
	// Уникальный идентификатор заказа.
	OrderUUID string
	// UUID пользователя.
	UserUUID string
	// Список UUID деталей.
	PartUuids []string
	// Итоговая стоимость.
	TotalPrice float64
	// UUID транзакции (если оплачен).
	TransactionUUID *string
	// Способ оплаты (если оплачен).
	PaymentMethod *PaymentMethod
	// Статус (PENDING_PAYMENT, PAID, CANCELLED).
	Status OrderStatus
}

type OrderStatus string

const (
	OrderStatusPendingPayment OrderStatus = "PENDING_PAYMENT"
	OrderStatusPaid           OrderStatus = "PAID"
	OrderStatusCancelled      OrderStatus = "CANCELLED"
)

type PaymentMethod string

const (
	PaymentMethodUnknown       PaymentMethod = "UNKNOWN"
	PaymentMethodCard          PaymentMethod = "CARD"
	PaymentMethodSBP           PaymentMethod = "SBP"
	PaymentMethodCreditCard    PaymentMethod = "CREDIT_CARD"
	PaymentMethodInvestorMoney PaymentMethod = "INVESTOR_MONEY"
)
