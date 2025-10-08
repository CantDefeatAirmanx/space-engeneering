package service_ship_assembly

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
	kafka_events_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/ship-assembly/v1"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

func (s *ShipAssemblyServiceImpl) createAutomaticAssembly(
	ctx context.Context,
	orderPaidEvent *order_events_v1.OrderPaidEvent,
) error {
	var startedTime time.Time
	var assembly *model_ship_assembly.ShipAssembly

	err := s.txManager.BeginTx(
		ctx,
		func(ctx context.Context, tx platform_transaction.Transaction) error {
			repoWithTx, err := getRepoWithTx(ctx, s)
			if err != nil {
				return err
			}
			defer func() {
				if r := recover(); r != nil {
					panic(r)
				}
			}()

			repoCreateRes, err := repoWithTx.CreateShipAssembly(
				ctx,
				&model_ship_assembly.ShipAssembly{
					AssemblyUUID: uuid.Must(uuid.NewV7()).String(),
					OrderUUID:    orderPaidEvent.OrderUuid,
					Status:       model_ship_assembly.ShipAssemblyStatusNotStarted,
				},
			)
			if err != nil {
				return err
			}

			assembly = repoCreateRes
			assemblySelectParams := model_ship_assembly.SelectShipAssemblyParams{
				OrderUUID: orderPaidEvent.OrderUuid,
			}

			err = repoWithTx.SetShipAssemblyStatusPending(ctx, assemblySelectParams)
			if err != nil {
				return err
			}

			startedTime = time.Now()
			ticker := time.NewTicker(10 * time.Second)
			defer ticker.Stop()

			select {
			case <-ticker.C:
				err = repoWithTx.SetShipAssemblyStatusCompleted(ctx, assemblySelectParams)
				if err != nil {
					return err
				}

				if err != nil {
					return err
				}
			case <-ctx.Done():
				return ctx.Err()
			}

			return nil
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

func getRepoWithTx(
	ctx context.Context,
	s *ShipAssemblyServiceImpl,
) (ShipAssemblyRepository, error) {
	executor, err := s.txManager.ExtractExecutorFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return s.repository.WithExecutor(executor), nil
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
