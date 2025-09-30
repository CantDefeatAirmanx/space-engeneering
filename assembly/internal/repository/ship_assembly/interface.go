package repository_ship_assembly

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

type ShipAssemblyRepository interface {
	CreateShipAssembly(ctx context.Context, shipAssembly *model_ship_assembly.ShipAssembly) (*model_ship_assembly.ShipAssembly, error)

	GetShipAssembly(
		ctx context.Context,
		selectParams model_ship_assembly.SelectShipAssemblyParams,
	) (*model_ship_assembly.ShipAssembly, error)

	UpdateShipAssembly(
		ctx context.Context,
		selectParams model_ship_assembly.SelectShipAssemblyParams,
		shipAssembly model_ship_assembly.UpdateShipAssemblyFields,
	) error

	SetShipAssemblyStatusPending(ctx context.Context, selectParams model_ship_assembly.SelectShipAssemblyParams) error
	SetShipAssemblyStatusCompleted(ctx context.Context, selectParams model_ship_assembly.SelectShipAssemblyParams) error

	platform_transaction.WithExecutor[ShipAssemblyRepository, platform_transaction.Executor]
}
