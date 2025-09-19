package app

import (
	"context"
	"path/filepath"

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

type App struct {
	closer closer.Closer
}

func NewApp(ctx context.Context, closer closer.Closer) (*App, error) {
	app := App{
		closer: closer,
	}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return &app, nil
}

func (a *App) Run(ctx context.Context) error {
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	type InitFunc func(context.Context) error

	initFuncs := []InitFunc{
		a.initConfig,
		a.initLogger,
	}

	for _, initFunc := range initFuncs {
		if err := initFunc(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	return config.LoadConfig(
		config.WithEnvPath(filepath.Join("assembly", ".env")),
	)
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		logger.WithLevel(config.Config.Logger().Level()),
		logger.WithEncoder(config.Config.Logger().Encoder()),
	)
}
