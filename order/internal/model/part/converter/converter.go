package model_part_converter

import (
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func ToModel(proto *inventory_v1.Part) model_part.Part {
	result := model_part.Part{
		UUID:          proto.Uuid,
		Name:          proto.Name,
		Description:   proto.Description,
		Price:         proto.Price,
		StockQuantity: proto.StockQuantity,
		Category:      model_part.Category(proto.Category),
		Tags:          proto.Tags,
		Metadata:      make(map[string]*model_part.MetaDataValue),
	}

	if proto.Dimensions != nil {
		result.Dimensions = &model_part.Dimensions{
			Length: proto.Dimensions.Length,
			Width:  proto.Dimensions.Width,
			Height: proto.Dimensions.Height,
			Weight: proto.Dimensions.Weight,
		}
	}

	if proto.Manufacturer != nil {
		result.Manufacturer = &model_part.Manufacturer{
			Name:    proto.Manufacturer.Name,
			Country: proto.Manufacturer.Country,
			Website: proto.Manufacturer.Website,
		}
	}

	if proto.CreatedAt != nil {
		result.CreatedAt = lo.ToPtr(proto.CreatedAt.AsTime())
	}

	if proto.UpdatedAt != nil {
		result.UpdatedAt = lo.ToPtr(proto.UpdatedAt.AsTime())
	}

	for key, value := range proto.Metadata {
		tmpString := value.GetStringValue()
		tmpDouble := value.GetDoubleValue()
		tmpInt64 := value.GetInt64Value()
		tmpBool := value.GetBoolValue()

		switch {
		case tmpString != "":
			result.Metadata[key] = model_part.NewStringValue(tmpString)
		case tmpDouble != 0:
			result.Metadata[key] = model_part.NewDoubleValue(tmpDouble)
		case tmpInt64 != 0:
			result.Metadata[key] = model_part.NewInt64Value(tmpInt64)
		case tmpBool:
			result.Metadata[key] = model_part.NewBoolValue(tmpBool)
		}
	}

	return result
}

func ToProto(model *model_part.Part) inventory_v1.Part {
	result := inventory_v1.Part{
		Uuid:          model.UUID,
		Name:          model.Name,
		Description:   model.Description,
		Price:         model.Price,
		StockQuantity: model.StockQuantity,
		Category:      inventory_v1.Category(model.Category),
		Tags:          model.Tags,
		Metadata:      make(map[string]*inventory_v1.Value),
	}

	if model.Dimensions != nil {
		result.Dimensions = &inventory_v1.Dimensions{
			Length: model.Dimensions.Length,
			Width:  model.Dimensions.Width,
			Height: model.Dimensions.Height,
			Weight: model.Dimensions.Weight,
		}
	}

	if model.Manufacturer != nil {
		result.Manufacturer = &inventory_v1.Manufacturer{
			Name:    model.Manufacturer.Name,
			Country: model.Manufacturer.Country,
			Website: model.Manufacturer.Website,
		}
	}

	if model.CreatedAt != nil {
		result.CreatedAt = timestamppb.New(*model.CreatedAt)
	}

	if model.UpdatedAt != nil {
		result.UpdatedAt = timestamppb.New(*model.UpdatedAt)
	}

	for key, value := range model.Metadata {
		result.Metadata[key] = &inventory_v1.Value{
			Value: &inventory_v1.Value_StringValue{
				StringValue: value.GetStringValue(),
			},
		}
	}

	return result
}
