package service_ship_assembly

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	kafka_events_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/ship-assembly/v1"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

func (s *ShipAssemblyServiceImpl) processOrderPaidEvent(
	ctx context.Context,
	orderPaidEvent *order_events_v1.OrderPaidEvent,
) error {
	assembly, err := s.CreateShipAssembly(
		ctx,
		CreateShipAssemblyParams{
			OrderUUID: orderPaidEvent.OrderUuid,
		},
	)
	if err != nil {
		return createErrWithInfo(err, orderPaidEvent)
	}

	_, err = s.AssemblyStarted(
		ctx,
		AssemblyStartedParams{
			OrderUUID:    orderPaidEvent.OrderUuid,
			AssemblyUUID: assembly.AssemblyUUID,
		},
	)
	if err != nil {
		return createErrWithInfo(err, orderPaidEvent)
	}

	startedTime := time.Now()
	time.Sleep(10 * time.Second)

	_, err = s.AssemblyCompleted(
		ctx,
		AssemblyCompletedParams{
			OrderUUID:    orderPaidEvent.OrderUuid,
			AssemblyUUID: assembly.AssemblyUUID,
		},
	)
	if err != nil {
		return createErrWithInfo(err, orderPaidEvent)
	}

	err = s.producer.ProduceAssemblyCompleted(
		ctx,
		kafka_events_ship_assembly.ShipAssembledEvent{
			AssemblyUUID: assembly.AssemblyUUID,
			OrderUUID:    assembly.OrderUUID,
			UserUUID:     orderPaidEvent.UserUuid,
			BuildTimeSec: int(time.Since(startedTime).Seconds()),
		},
	)
	if err != nil {
		return createErrWithInfo(err, orderPaidEvent)
	}
	logger.Logger().Info(
		"Assembly completed",
		zap.String("assembly_uuid", assembly.AssemblyUUID),
		zap.String("order_uuid", assembly.OrderUUID),
		zap.String("user_uuid", orderPaidEvent.UserUuid),
		zap.Int("build_time_sec", int(time.Since(startedTime).Seconds())),
	)

	return nil
}

func createErrWithInfo(
	err error,
	orderPaidEvent *order_events_v1.OrderPaidEvent,
) error {
	return fmt.Errorf(
		"%w: user_uuid: %s, order_uuid: %s",
		err,
		orderPaidEvent.UserUuid,
		orderPaidEvent.OrderUuid,
	)
}
