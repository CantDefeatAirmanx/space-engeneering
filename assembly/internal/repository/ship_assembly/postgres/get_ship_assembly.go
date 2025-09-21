package repository_ship_assembly_postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyRepositoryPostgres) GetShipAssembly(
	ctx context.Context,
	params model_ship_assembly.SelectShipAssemblyParams,
) (*model_ship_assembly.ShipAssembly, error) {
	query, args, err := squirrel.
		Select(
			columnAssemblyUUID,
			columnOrderUUID,
			columnStatus,
			columnCreatedAt,
			columnUpdatedAt,
		).
		From(tableShipAssembly).
		Where(getSquirelShipAssemblySelectParams(params)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var result ShipAssembly
	err = s.db.QueryRow(ctx, query, args...).Scan(
		&result.AssemblyUUID,
		&result.OrderUUID,
		&result.Status,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model_ship_assembly.ErrAssemblyNotFound
		}
		return nil, err
	}

	return AssemblyToModel(result), nil
}
