package api_inventory_v1

import (
	"context"

	model_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part/converter"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
	"go.uber.org/zap"
)

func (api *api) GetPart(
	ctx context.Context,
	req *inventory_v1.GetPartRequest,
) (*inventory_v1.GetPartResponse, error) {
	contexts.GetLogParamsSetterFunc(ctx)(
		[]zap.Field{
			zap.String(partUUIDLogKey, req.Uuid),
		},
	)

	part, err := api.partService.GetPart(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	protoPart := model_converter_part.ToProto(part)

	return &inventory_v1.GetPartResponse{
		Part: protoPart,
	}, nil
}
