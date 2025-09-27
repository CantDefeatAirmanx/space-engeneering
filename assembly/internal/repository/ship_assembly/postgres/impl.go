package repository_ship_assembly_postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	repository_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/repository/ship_assembly"
)

var _ repository_ship_assembly.ShipAssemblyRepository = (*ShipAssemblyRepositoryPostgres)(nil)

type ShipAssemblyRepositoryPostgres struct {
	db *pgxpool.Pool
}

func NewShipAssemblyRepositoryPostgres(db *pgxpool.Pool) *ShipAssemblyRepositoryPostgres {
	return &ShipAssemblyRepositoryPostgres{db: db}
}
