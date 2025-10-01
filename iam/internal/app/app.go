package app

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"strconv"

	"github.com/CantDefeatAirmanx/space-engeneering/iam/config"
	"github.com/CantDefeatAirmanx/space-engeneering/iam/internal/app/di"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/interceptor"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	closer      closer.Closer
	diContainer *di.DiContainer
	listener    net.Listener
	grpcServer  *grpc.Server
}

func NewApp(ctx context.Context, closer closer.Closer) (*App, error) {
	app := &App{
		closer: closer,
	}

	if err := app.initDeps(ctx); err != nil {
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

func (a *App) runGRPCServer(_ context.Context) error {
	logger.Logger().Info(
		fmt.Sprintf("running IAM GRPC server on %s", a.listener.Addr().String()),
	)

	return a.grpcServer.Serve(a.listener)
}

func (a *App) initDeps(ctx context.Context) error {
	type InitFunc func(ctx context.Context) error

	initFuncs := []InitFunc{
		a.initConfig,
		a.initLogger,
		a.initDiContainer,
		a.initListener,
		a.initGrpcServer,
	}

	for _, initFunc := range initFuncs {
		if err := initFunc(ctx); err != nil {
			return err
		}
	}

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

	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.ValidateInterceptor(),
			interceptor.WithLogParamsInterceptor(),
			interceptor.UnaryErrorInterceptor(
				interceptor.WithLogger(logger.Logger()),
			),
		),
	)
	a.closer.AddNamed("IAM GRPC Server", func(ctx context.Context) error {
		grpcServer.GracefulStop()

		return nil
	})
	reflection.Register(grpcServer)

	// auth_v1.RegisterAuthServiceServer(grpcServer, a.diContainer.GetAuthV1API(ctx))
	// user_v1.RegisterUserServiceServer(grpcServer, a.diContainer.GetUserV1API(ctx))

	a.grpcServer = grpcServer

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	return config.LoadConfig(
		config.WithEnvPath(filepath.Join("iam", ".env")),
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
