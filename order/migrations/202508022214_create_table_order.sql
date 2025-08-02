-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    order_uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_uuid UUID NOT NULL,
    part_uuids UUID[] NOT NULL DEFAULT '{}',
    total_price DECIMAL(12,2) NOT NULL DEFAULT 0.00,
    transaction_uuid UUID NULL,

    payment_method VARCHAR(50) NULL CHECK (
        payment_method IN ('UNKNOWN', 'CARD', 'SBP', 'CREDIT_CARD', 'INVESTOR_MONEY')
    ),

    status VARCHAR(50) NOT NULL CHECK (
        status IN ('PENDING_PAYMENT', 'PAID', 'CANCELLED')
    ) DEFAULT 'PENDING_PAYMENT',

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_orders_user_uuid ON orders (user_uuid);
CREATE INDEX idx_orders_status ON orders (status);
CREATE INDEX idx_orders_created_at ON orders (created_at);
CREATE INDEX idx_orders_transaction_uuid ON orders (transaction_uuid) WHERE transaction_uuid IS NOT NULL;

-- Функция для автоматического обновления updated_at
-- +goose statementbegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose statementend

-- Триггер для автоматического обновления updated_at при изменении записи
CREATE TRIGGER update_orders_updated_at 
    BEFORE UPDATE ON orders 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_orders_updated_at ON orders;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP INDEX IF EXISTS idx_orders_transaction_uuid;
DROP INDEX IF EXISTS idx_orders_created_at; 
DROP INDEX IF EXISTS idx_orders_status;
DROP INDEX IF EXISTS idx_orders_user_uuid;
DROP TABLE IF EXISTS orders;