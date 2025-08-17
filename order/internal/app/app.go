package app

import (
	"context"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/CantDefeatAirmanx/space-engeneering/order/config"
	"github.com/CantDefeatAirmanx/space-engeneering/order/internal/app/di"
)

type App struct {
	diContainer *di.DiContainer
	httpServer  *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	if err := app.init(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.runHttpServer(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) runHttpServer(_ context.Context) error {
	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (a *App) init(ctx context.Context) error {
	type initFunc func(ctx context.Context) error

	initFuncs := []initFunc{
		a.initConfig,
		a.initDiContainer,
		a.initHttpServer,
	}

	for _, initFunc := range initFuncs {
		if err := initFunc(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	envPath := filepath.Join("order", ".env")

	return config.LoadConfig(ctx, config.WithEnvPath(envPath))
}

func (a *App) initDiContainer(_ context.Context) error {
	diContainer := &di.DiContainer{}

	a.diContainer = diContainer

	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(
		time.Duration(config.Config.HttpServer().Timeout()) * time.Millisecond,
	))

	router.Mount("/", a.diContainer.GetOrderServer(ctx))

	a.httpServer = &http.Server{
		Addr: net.JoinHostPort(
			config.Config.HttpServer().Host(),
			strconv.Itoa(config.Config.HttpServer().Port()),
		),
		Handler: router,
		ReadHeaderTimeout: time.Duration(
			config.Config.HttpServer().ReadHeaderTimeout(),
		) * time.Millisecond,
	}

	return nil
}
