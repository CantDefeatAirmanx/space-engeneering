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

	"github.com/CantDefeatAirmanx/space-engeneering/payment/config"
	"github.com/CantDefeatAirmanx/space-engeneering/payment/internal/app/di"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/interceptor"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

type App struct {
	diContainer *di.DiContainer
	listener    net.Listener
	grpcServer  *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := App{}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return &app, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.runGrpcServer(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) runGrpcServer(_ context.Context) error {
	logger.Logger().Info(fmt.Sprintf("running Payment GRPC server on %s", a.listener.Addr().String()))

	if err := a.grpcServer.Serve(a.listener); err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	type InitFunc func(context.Context) error

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

func (a *App) initConfig(_ context.Context) error {
	opts := []config.LoadConfigOption{}
	if config.IS_DEV {
		opts = append(opts, config.WithEnvPath(
			filepath.Join("payment", ".env"),
		))
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

func (a *App) initListener(_ context.Context) error {
	listener, err := net.Listen("tcp", net.JoinHostPort(
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
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.WithLogParamsInterceptor(),
			interceptor.ValidateInterceptor(),
			interceptor.UnaryErrorInterceptor(
				interceptor.WithLogger(logger.Logger()),
			),
		),
	)

	reflection.Register(grpcServer)
	payment_v1.RegisterPaymentServiceServer(grpcServer, a.diContainer.GetPaymentV1Api(ctx))

	a.grpcServer = grpcServer

	return nil
}
