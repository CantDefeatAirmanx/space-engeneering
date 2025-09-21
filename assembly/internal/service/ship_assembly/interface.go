package service_ship_assembly

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
)

type ShipAssemblyService interface {
	CreateShipAssembly(
		ctx context.Context,
		params CreateShipAssemblyParams,
	) (*model_ship_assembly.ShipAssembly, error)

	AssemblyStarted(
		ctx context.Context,
		params AssemblyStartedParams,
	) (*AssemblyStartedReturn, error)

	AssemblyCompleted(
		ctx context.Context,
		params AssemblyCompletedParams,
	) (*AssemblyCompletedReturn, error)

	GetAssemblyInfo(
		ctx context.Context,
		params GetAssemblyInfoParams,
	) (*model_ship_assembly.ShipAssembly, error)

	interfaces.WithClose

	watchOrderPaidEvent(ctx context.Context)
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
