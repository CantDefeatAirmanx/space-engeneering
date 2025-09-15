package app

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/config"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/app/di"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/interceptor"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

type App struct {
	grpcServer *grpc.Server
	listener   net.Listener
	closer     closer.Closer
	di         *di.DiContainer
}

func NewApp(ctx context.Context, closer closer.Closer) (*App, error) {
	app := &App{
		closer: closer,
		di:     di.NewDiContainer(closer),
	}

	if err := app.init(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.runGRPCServer(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) Close(ctx context.Context) error {
	return a.closer.CloseAll(ctx)
}

func (a *App) runGRPCServer(_ context.Context) error {
	logger.Logger().Info(
		fmt.Sprintf("running Inventory GRPC server on %s", a.listener.Addr().String()),
	)

	if err := a.grpcServer.Serve(a.listener); err != nil {
		return err
	}

	return nil
}

func (a *App) init(ctx context.Context) error {
	type initFunc func(ctx context.Context) error

	initFuncs := []initFunc{
		a.initConfig,
		a.initLogger,
		a.initListener,
		a.initGRPCServer,
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
		config.WithEnvPath(filepath.Join("inventory", ".env")),
	)
}

func (a *App) initLogger(_ context.Context) error {
	err := logger.Init(
		logger.WithLevel(config.Config.Logger().Level()),
		logger.WithEncoder(config.Config.Logger().Encoder()),
	)
	if err != nil {
		return err
	}

	a.closer.SetLogger(logger.Logger())

	return nil
}

func (a *App) initListener(ctx context.Context) error {
	listener, err := net.Listen(
		"tcp4",
		net.JoinHostPort(
			config.Config.GRPC().Host(),
			strconv.Itoa(config.Config.GRPC().Port()),
		),
	)
	if err != nil {
		return err
	}

	a.listener = listener

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.ValidateInterceptor(),
			interceptor.WithLogParamsInterceptor(),
			interceptor.UnaryErrorInterceptor(
				interceptor.WithLogger(logger.Logger()),
			),
		),
		grpc.Creds(insecure.NewCredentials()),
	)
	a.closer.AddNamed("Inventory GRPC server", func(ctx context.Context) error {
		a.grpcServer.GracefulStop()

		return nil
	})

	inventory_v1.RegisterInventoryServiceServer(
		a.grpcServer,
		a.di.GetInventoryAPI(ctx),
	)
	reflection.Register(a.grpcServer)

	return nil
}
