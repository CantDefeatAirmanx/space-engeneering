package model_ship_assembly

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

var assemblyModelStatusToProto = map[ShipAssemblyStatus]assembly_v1.AssemblyStatus{
	ShipAssemblyStatusUnspecified: assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_UNSPECIFIED,
	ShipAssemblyStatusNotStarted:  assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_NOT_STARTED,
	ShipAssemblyStatusPending:     assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_PENDING,
	ShipAssemblyStatusCompleted:   assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_COMPLETED,
}

var assemblyProtoStatusToModel = map[assembly_v1.AssemblyStatus]ShipAssemblyStatus{
	assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_UNSPECIFIED: ShipAssemblyStatusUnspecified,
	assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_NOT_STARTED: ShipAssemblyStatusNotStarted,
	assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_PENDING:     ShipAssemblyStatusPending,
	assembly_v1.AssemblyStatus_ASSEMBLY_STATUS_COMPLETED:   ShipAssemblyStatusCompleted,
}

func ConvertToProto(model ShipAssembly) *assembly_v1.AssemblyInfo {
	return &assembly_v1.AssemblyInfo{
		AssemblyUuid: model.AssemblyUUID,
		OrderUuid:    model.OrderUUID,
		Status:       assemblyModelStatusToProto[model.Status],

		CreatedAt: timestamppb.New(model.CreatedAt),
		UpdatedAt: timestamppb.New(model.UpdatedAt),
	}
}

func ConvertToModel(proto *assembly_v1.AssemblyInfo) ShipAssembly {
	return ShipAssembly{
		AssemblyUUID: proto.AssemblyUuid,
		OrderUUID:    proto.OrderUuid,
		Status:       assemblyProtoStatusToModel[proto.Status],

		CreatedAt: proto.CreatedAt.AsTime(),
		UpdatedAt: proto.UpdatedAt.AsTime(),
	}
}
