-- +goose Up
CREATE TYPE payment_method_enum AS ENUM ('UNKNOWN', 'CARD', 'SBP', 'CREDIT_CARD', 'INVESTOR_MONEY');
CREATE TYPE order_status_enum AS ENUM ('PENDING_PAYMENT', 'PAID', 'CANCELLED');

ALTER TABLE orders DROP CONSTRAINT IF EXISTS orders_payment_method_check;
ALTER TABLE orders DROP CONSTRAINT IF EXISTS orders_status_check;

-- Удаляем текущий default, чтобы избежать ошибки автокаста при смене типа
ALTER TABLE orders
	ALTER COLUMN status DROP DEFAULT;

-- Меняем типы колонок с явным USING-кастом
ALTER TABLE orders
	ALTER COLUMN payment_method TYPE payment_method_enum USING payment_method::payment_method_enum,
	ALTER COLUMN status TYPE order_status_enum USING status::order_status_enum;

-- Устанавливаем новый default уже для enum-типизированной колонки
ALTER TABLE orders
	ALTER COLUMN status SET DEFAULT 'PENDING_PAYMENT'::order_status_enum;

-- +goose Down
ALTER TABLE orders
	ALTER COLUMN status DROP DEFAULT,
	ALTER COLUMN status TYPE VARCHAR(50),
	ALTER COLUMN payment_method TYPE VARCHAR(50);

ALTER TABLE orders
	ADD CONSTRAINT orders_status_check CHECK (
		status IN ('PENDING_PAYMENT', 'PAID', 'CANCELLED')
	),
	ADD CONSTRAINT orders_payment_method_check CHECK (
		payment_method IN ('UNKNOWN', 'CARD', 'SBP', 'CREDIT_CARD', 'INVESTOR_MONEY')
	);

DROP TYPE order_status_enum;
DROP TYPE payment_method_enum;
