package app

import (
	"context"
	"net"
	"path/filepath"
	"strconv"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/config"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/app/di"
	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interceptor"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	if err := app.init(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.grpcServer.Serve(a.listener); err != nil {
		return err
	}

	return nil
}

func (a *App) init(ctx context.Context) error {
	type initFunc func(ctx context.Context) error

	initFuncs := []initFunc{
		a.initConfig,
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

func (a *App) initListener(ctx context.Context) error {
	listener, err := net.Listen(
		"tcp",
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
			interceptor.UnaryErrorInterceptor(),
			interceptor.ValidateInterceptor(),
		),
		grpc.Creds(insecure.NewCredentials()),
	)

	inventory_v1.RegisterInventoryServiceServer(
		a.grpcServer,
		di.NewDiContainer().GetInventoryAPI(),
	)
	reflection.Register(a.grpcServer)

	return nil
}
