package repository_ship_assembly_postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyRepositoryPostgres) UpdateShipAssembly(
	ctx context.Context,
	selectParams model_ship_assembly.SelectShipAssemblyParams,
	shipAssembly model_ship_assembly.UpdateShipAssemblyFields,
) error {
	updatedAt := time.Now()

	repoUpdate, err := AssemblyUpdateFieldsToRepo(&shipAssembly)
	if err != nil {
		return err
	}

	builder := squirrel.Update(tableShipAssembly)
	if repoUpdate.StartTime != nil {
		builder = builder.Set(columnStartTime, repoUpdate.StartTime)
	}

	if repoUpdate.Status != "" {
		builder = builder.Set(columnStatus, repoUpdate.Status)
	}

	query, args, err := builder.
		Set(columnUpdatedAt, updatedAt).
		Where(getSquirelShipAssemblySelectParams(selectParams)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	res, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return model_ship_assembly.ErrAssemblyNotFound
	}

	return nil
}
