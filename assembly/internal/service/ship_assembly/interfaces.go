package service_ship_assembly

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
	repository_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/repository/ship_assembly"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
)

type ShipAssemblyService interface {
	CreateShipAssembly(
		ctx context.Context,
		params CreateShipAssemblyParams,
	) (*model_ship_assembly.ShipAssembly, error)

	SetAssemblyStarted(
		ctx context.Context,
		params AssemblyStartedParams,
	) (*AssemblyStartedReturn, error)

	SetAssemblyCompleted(
		ctx context.Context,
		params AssemblyCompletedParams,
	) (*AssemblyCompletedReturn, error)

	GetAssemblyInfo(
		ctx context.Context,
		params GetAssemblyInfoParams,
	) (*model_ship_assembly.ShipAssembly, error)

	interfaces.WithClose
}

type CreateShipAssemblyParams struct {
	OrderUUID string
}

type CreateShipAssemblyReturn struct {
	AssemblyInfo *model_ship_assembly.ShipAssembly
}

type AssemblyCompletedParams struct {
	OrderUUID    string
	AssemblyUUID string
}

type AssemblyCompletedReturn struct {
	AssemblyUUID string
}

type AssemblyStartedParams struct {
	OrderUUID    string
	AssemblyUUID string
}

type AssemblyStartedReturn struct {
	AssemblyUUID string
}

type GetAssemblyInfoParams struct {
	AssemblyUUID string
	OrderUUID    string
}

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

	platform_transaction.WithExecutor[
		repository_ship_assembly.ShipAssemblyRepository,
		platform_transaction.Executor,
	]
}
