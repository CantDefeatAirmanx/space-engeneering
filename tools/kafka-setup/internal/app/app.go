package app

import (
	"context"
	"path/filepath"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"github.com/CantDefeatAirmanx/space-engeneering/tools/kafka-setup/config"
	"github.com/CantDefeatAirmanx/space-engeneering/tools/kafka-setup/internal/app/di"
)

type App struct {
	diContainer *di.DiContainer
	closer      closer.Closer
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
	if err := a.runKafkaSetup(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) runKafkaSetup(ctx context.Context) error {
	topicsConfigurator := a.diContainer.GetTopicsConfigurator()
	a.closer.AddNamed("Topics Configurator", topicsConfigurator.Close)

	return topicsConfigurator.Configure(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	type InitFunc func(context.Context) error

	initFuncs := []InitFunc{
		a.initConfig,
		a.initLogger,
		a.initDiContainer,
	}

	for _, initFunc := range initFuncs {
		if err := initFunc(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	opts := []config.LoadConfigOption{
		config.WithEnvPath(
			filepath.Join("tools", "kafka-setup", ".env"),
		),
	}

	if err := config.LoadConfig(opts...); err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		logger.WithLevel(config.Config.Logger().Level()),
		logger.WithEncoder(config.Config.Logger().Encoder()),
	)
}

func (a *App) initDiContainer(_ context.Context) error {
	a.diContainer = di.NewDiContainer()

	return nil
}
