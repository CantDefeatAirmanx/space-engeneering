package repository_order_postgre

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Order struct {
	OrderUUID       pgtype.UUID      `db:"order_uuid"`
	UserUUID        pgtype.UUID      `db:"user_uuid"`
	PartUuids       []pgtype.UUID    `db:"part_uuids"`
	TotalPrice      pgtype.Numeric   `db:"total_price"`
	TransactionUUID *pgtype.UUID     `db:"transaction_uuid"`
	PaymentMethod   *PaymentMethod   `db:"payment_method"`
	Status          OrderStatus      `db:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at"`
}

type PaymentMethod string

const (
	PaymentMethodUnknown       PaymentMethod = "UNKNOWN"
	PaymentMethodCard          PaymentMethod = "CARD"
	PaymentMethodSBP           PaymentMethod = "SBP"
	PaymentMethodCreditCard    PaymentMethod = "CREDIT_CARD"
	PaymentMethodInvestorMoney PaymentMethod = "INVESTOR_MONEY"
)

type OrderStatus string

const (
	OrderStatusPendingPayment OrderStatus = "PENDING_PAYMENT"
	OrderStatusPaid           OrderStatus = "PAID"
	OrderStatusCancelled      OrderStatus = "CANCELLED"
)
