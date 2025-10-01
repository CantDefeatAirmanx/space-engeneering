-- +goose Up
CREATE UNIQUE INDEX idx_ship_assembly_order_uuid 
ON assemblies (order_uuid);

-- +goose Down
DROP INDEX idx_ship_assembly_order_uuid;