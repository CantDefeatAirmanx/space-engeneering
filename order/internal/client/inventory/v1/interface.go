package client_inventory_v1

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

type InventoryV1Client interface {
	// GetPart checks if the part is in the inventory and returns it or an error
	//
	// Errors:
	//
	// - ([client_inventory_v1.ErrPartNotFound]): if the part is not found
	//
	// - ([client_inventory_v1.ErrInternalServerError]): if the server is not available
	//
	// - ([client_inventory_v1.ErrInvalidArguments]): if the arguments are invalid
	GetPart(
		ctx context.Context,
		params *inventory_v1.GetPartRequest,
	) (*model_part.Part, error)

	// ListParts returns a list of parts that match the filter
	//
	// Errors:
	//
	// - ([client_inventory_v1.ErrInternalServerError]): if the server is not available
	//
	// - ([client_inventory_v1.ErrInvalidArguments]): if the arguments are invalid
	ListParts(
		ctx context.Context,
		params ListPartsParams,
	) ([]*model_part.Part, error)
}

type GetPartParams struct {
	UUID string
}

type ListPartsParams struct {
	Uuids                 []string
	Categories            []model_part.Category
	ManufacturerCountries []string
	Tags                  []string
	Names                 []string
}
