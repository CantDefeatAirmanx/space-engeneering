package repository_ship_assembly_postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyRepositoryPostgres) CreateShipAssembly(
	ctx context.Context,
	shipAssembly *model_ship_assembly.ShipAssembly,
) (*model_ship_assembly.ShipAssembly, error) {
	query, args, err := squirrel.Insert(tableShipAssembly).
		Columns(columnAssemblyUUID, columnOrderUUID, columnStatus).
		Values(shipAssembly.AssemblyUUID, shipAssembly.OrderUUID, shipAssembly.Status).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}

	res := model_ship_assembly.ShipAssembly{
		AssemblyUUID: shipAssembly.AssemblyUUID,
		OrderUUID:    shipAssembly.OrderUUID,
		Status:       shipAssembly.Status,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &res, nil
}
