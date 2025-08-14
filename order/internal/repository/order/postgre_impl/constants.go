package repository_order_postgre

const (
	tableOrders = "orders"

	columnOrderUUID       = "order_uuid"
	columnUserUUID        = "user_uuid"
	columnPartsUUIDs      = "part_uuids"
	columnTotalPrice      = "total_price"
	columnTransactionUUID = "transaction_uuid"
	columnPaymentMethod   = "payment_method"
	columnStatus          = "status"
	columnCreatedAt       = "created_at"
	columnUpdatedAt       = "updated_at"
)
