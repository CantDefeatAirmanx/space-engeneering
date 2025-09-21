package repository_ship_assembly_postgres

import (
	"github.com/Masterminds/squirrel"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func getSquirelShipAssemblySelectParams(
	params model_ship_assembly.SelectShipAssemblyParams,
) squirrel.Eq {
	res := squirrel.Eq{}

	if params.AssemblyUUID != "" {
		res[columnAssemblyUUID] = params.AssemblyUUID
	}

	if params.OrderUUID != "" {
		res[columnOrderUUID] = params.OrderUUID
	}

	return res
}
