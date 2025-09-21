package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/migrator"
)

func main() {
	ctx := context.Background()
	closer, done := closer.NewCloser(ctx)

	defer func() {
		go func() {
			if err := closer.CloseAll(ctx); err != nil {
				logger.Logger().Error(fmt.Sprintf("Failed to close app: %v", err))
			}
		}()
		<-done
	}()

	if err := initConfig(ctx); err != nil {
		logger.DefaultInfoLogger().Error(fmt.Sprintf("Failed to init config: %v", err))
		return
	}

	if err := initLogger(ctx); err != nil {
		logger.DefaultInfoLogger().Error(fmt.Sprintf("Failed to init logger: %v", err))
		return
	}

	conn, err := pgx.Connect(ctx, config.Config.Postgres().Uri())
	if err != nil {
		logger.DefaultInfoLogger().Error(fmt.Sprintf("Failed to create db: %v", err))
		return
	}
	closer.AddNamed("Orders Postgres Db", func(ctx context.Context) error {
		if err := conn.Close(ctx); err != nil {
			return err
		}
		return nil
	})

	sqlDb := stdlib.OpenDB(*conn.Config().Copy())
	closer.AddNamed("Orders Postgres Db", func(ctx context.Context) error {
		if err := sqlDb.Close(); err != nil {
			return err
		}
		return nil
	})

	migrator := migrator.NewMigrator(
		sqlDb,
		filepath.Join("assembly", "migrations"),
	)

	if err := migrator.Up(); err != nil {
		logger.DefaultInfoLogger().Error(fmt.Sprintf("Failed to run migrator: %v", err))
		return
	}
}

func initLogger(_ context.Context) error {
	return logger.Init(
		logger.WithLevel(config.Config.Logger().Level()),
		logger.WithEncoder(config.Config.Logger().Encoder()),
	)
}

func initConfig(_ context.Context) error {
	return config.LoadConfig(
		config.WithEnvPath(filepath.Join("assembly", ".env")),
	)
}
