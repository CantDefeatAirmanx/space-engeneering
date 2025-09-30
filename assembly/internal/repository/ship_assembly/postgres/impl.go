package repository_ship_assembly_postgres

import (
	repository_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/repository/ship_assembly"
	platform_postgres "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/db/postgres"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

var _ repository_ship_assembly.ShipAssemblyRepository = (*ShipAssemblyRepositoryPostgres)(nil)

type ShipAssemblyRepositoryPostgres struct {
	db platform_postgres.Executor
}

func NewShipAssemblyRepositoryPostgres(
	db platform_postgres.Executor,
) *ShipAssemblyRepositoryPostgres {
	return &ShipAssemblyRepositoryPostgres{db: db}
}

func (s *ShipAssemblyRepositoryPostgres) WithExecutor(
	executor platform_transaction.Executor,
) repository_ship_assembly.ShipAssemblyRepository {
	return NewShipAssemblyRepositoryPostgres(executor)
}
