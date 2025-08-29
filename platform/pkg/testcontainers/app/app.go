package app

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

const (
	defaultAppName        = "app"
	defaultAppPort        = "50051"
	defaultStartupTimeout = 1 * time.Minute
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type Config struct {
	Name          string
	DockerfileDir string
	Dockerfile    string
	Port          string
	Env           map[string]string
	Networks      []string
	LogOutput     io.Writer
	StartupWait   wait.Strategy
	Logger        Logger
}

type Container struct {
	container    testcontainers.Container
	externalHost string
	externalPort string
	cfg          *Config
}

func NewContainer(ctx context.Context, opts ...Option) (*Container, error) {
	cfg := &Config{
		Name:          defaultAppName,
		Port:          defaultAppPort,
		Dockerfile:    "Dockerfile",
		DockerfileDir: ".",
		LogOutput:     io.Discard,
		StartupWait:   wait.ForListeningPort(defaultAppPort + "/tcp").WithStartupTimeout(defaultStartupTimeout),
		Env:           make(map[string]string),
		Logger:        logger.DefaultInfoLogger(),
	}
	for _, opt := range opts {
		opt(cfg)
	}

	req := testcontainers.ContainerRequest{
		Name: cfg.Name,
		FromDockerfile: testcontainers.FromDockerfile{
			Context:        cfg.DockerfileDir,
			Dockerfile:     cfg.Dockerfile,
			BuildLogWriter: cfg.LogOutput,
		},
		Networks:           cfg.Networks,
		Env:                cfg.Env,
		WaitingFor:         cfg.StartupWait,
		ExposedPorts:       []string{cfg.Port + "/tcp"},
		HostConfigModifier: DefaultHostConfig(),
	}

	cfg.Logger.Info("Creating app container", zap.String("name", cfg.Name), zap.String("port", cfg.Port))

	genericContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		cfg.Logger.Error("Failed to start app container", zap.Error(err))
		return nil, errors.Errorf("failed to start app genericContainer: %v", err)
	}

	cfg.Logger.Info("App container created successfully")

	mappedPort, err := genericContainer.MappedPort(ctx, nat.Port(cfg.Port+"/tcp"))
	if err != nil {
		cfg.Logger.Error("Failed to get mapped port", zap.Error(err))
		return nil, errors.Errorf("failed to get mapped externalPort: %v", err)
	}

	host, err := genericContainer.Host(ctx)
	if err != nil {
		cfg.Logger.Error("Failed to get container host", zap.Error(err))
		return nil, errors.Errorf("failed to get genericContainer externalHost: %v", err)
	}

	cfg.Logger.Info("App container started", zap.String("uri:", net.JoinHostPort(host, mappedPort.Port())))

	prefixedWriter := NewPrefixedWriter(cfg.LogOutput, fmt.Sprintf("CONTAINER[%s]", cfg.Name))
	streamContainerLogs(ctx, genericContainer, prefixedWriter)

	return &Container{
		container:    genericContainer,
		externalHost: host,
		externalPort: mappedPort.Port(),
		cfg:          cfg,
	}, nil
}

func (a *Container) Address() string {
	return net.JoinHostPort(a.externalHost, a.externalPort)
}

func (a *Container) InternalAddress() string {
	return net.JoinHostPort(a.externalHost, a.cfg.Port)
}

func (a *Container) Terminate(ctx context.Context) error {
	return a.container.Terminate(ctx)
}

func (a *Container) Status(ctx context.Context) error {
	state, err := a.container.State(ctx)
	if err != nil {
		a.cfg.Logger.Error("Failed to get container state", zap.Error(err))
		return err
	}

	a.cfg.Logger.Info("Container state",
		zap.String("status", state.Status),
		zap.Bool("running", state.Running),
		zap.Int("exit_code", state.ExitCode),
		zap.String("error", state.Error),
	)

	return nil
}

func streamContainerLogs(ctx context.Context, container testcontainers.Container, out io.Writer) {
	go func() {
		logs, err := container.Logs(ctx)
		if err != nil {
			logger.Logger().Error("failed to get container logs", zap.Error(err))
			return
		}
		defer func() {
			if closeErr := logs.Close(); closeErr != nil {
				logger.Logger().Error("failed to close container logs", zap.Error(closeErr))
			}
		}()

		_, err = io.Copy(out, logs)
		if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrClosedPipe) {
			logger.Logger().Error("error copying container logs", zap.Error(err))
		}
	}()
}

func DefaultHostConfig() func(hc *container.HostConfig) {
	return func(hc *container.HostConfig) {
		hc.AutoRemove = false
	}
}
