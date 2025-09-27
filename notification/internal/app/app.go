package app

import (
	"context"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/notification/config"
	"github.com/CantDefeatAirmanx/space-engeneering/notification/internal/app/di"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

type App struct {
	closer      closer.Closer
	diContainer *di.DiContainer
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
	go func() {
		if err := a.runAssembliesWatcherService(ctx); err != nil {
			logger.Logger().Error("failed to run assemblies watcher service", zap.Error(err))
			panic(err)
		}
	}()

	if err := a.runOrdersWatcherService(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) runAssembliesWatcherService(ctx context.Context) error {
	return a.diContainer.GetAssembliesWatcherService(ctx).WatchAssemblies(ctx)
}

func (a *App) runOrdersWatcherService(ctx context.Context) error {
	return a.diContainer.GetOrdersWatcherService(ctx).WatchOrders(ctx)
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

func (a *App) initConfig(ctx context.Context) error {
	return config.LoadConfig(
		config.WithEnvPath(filepath.Join("notification", ".env")),
	)
}

func (a *App) initLogger(ctx context.Context) error {
	return logger.Init(
		logger.WithLevel(config.Config.Logger().Level()),
		logger.WithEncoder(config.Config.Logger().Encoder()),
	)
}

func (a *App) initDiContainer(ctx context.Context) error {
	a.diContainer = di.NewDiContainer(a.closer)

	return nil
}
