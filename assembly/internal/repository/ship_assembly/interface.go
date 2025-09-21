package repository_ship_assembly

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

type ShipAssemblyRepository interface {
	CreateShipAssembly(ctx context.Context, shipAssembly *model_ship_assembly.ShipAssembly) error

	GetShipAssembly(
		ctx context.Context,
		selectParams model_ship_assembly.SelectShipAssemblyParams,
	) (*model_ship_assembly.ShipAssembly, error)

	UpdateShipAssembly(
		ctx context.Context,
		selectParams model_ship_assembly.SelectShipAssemblyParams,
		shipAssembly model_ship_assembly.UpdateShipAssemblyFields,
	) error
}
