-- +goose Up
ALTER TABLE assemblies ADD COLUMN assembly_start_time TIMESTAMP WITH TIME ZONE;

-- +goose Down
ALTER TABLE assemblies DROP COLUMN assembly_start_time;