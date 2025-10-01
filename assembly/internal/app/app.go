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

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	"github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/app/di"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/interceptor"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

type App struct {
	closer      closer.Closer
	diContainer *di.DiContainer
	listener    net.Listener
	grpcServer  *grpc.Server
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
	go a.runOrdersWatcherConsumer(ctx)

	if err := a.runGRPCServer(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPCServer(_ context.Context) error {
	logger.Logger().Info(fmt.Sprintf("running Assembly GRPC server on %s", a.listener.Addr().String()))

	if err := a.grpcServer.Serve(a.listener); err != nil {
		return err
	}

	return nil
}

func (a *App) runOrdersWatcherConsumer(ctx context.Context) {
	a.diContainer.GetShipAssemblyConsumer(ctx).WatchOrderPaidEvent(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	type InitFunc func(context.Context) error

	initFuncs := []InitFunc{
		a.initConfig,
		a.initLogger,
		a.initDiContainer,
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
		config.WithEnvPath(filepath.Join("assembly", ".env")),
	)
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		logger.WithLevel(config.Config.Logger().Level()),
		logger.WithEncoder(config.Config.Logger().Encoder()),
	)
}

func (a *App) initDiContainer(ctx context.Context) error {
	a.diContainer = di.NewDiContainer(a.closer)

	return nil
}

func (a *App) initListener(ctx context.Context) error {
	listener, err := net.Listen("tcp4", net.JoinHostPort(
		config.Config.GRPC().Host(),
		strconv.Itoa(config.Config.GRPC().Port()),
	))
	if err != nil {
		return err
	}

	a.listener = listener
	a.closer.AddNamed("Assembly GRPC Listener", func(ctx context.Context) error {
		return listener.Close()
	})

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.WithLogParamsInterceptor(),
			interceptor.ValidateInterceptor(),
			interceptor.UnaryErrorInterceptor(
				interceptor.WithLogger(logger.Logger()),
			),
		),
	)
	a.closer.AddNamed("Assembly GRPC Server", func(ctx context.Context) error {
		grpcServer.GracefulStop()

		return nil
	})

	reflection.Register(grpcServer)
	assembly_v1.RegisterAssemblyServiceServer(grpcServer, a.diContainer.GetAssemblyApi(ctx))

	a.grpcServer = grpcServer

	return nil
}
