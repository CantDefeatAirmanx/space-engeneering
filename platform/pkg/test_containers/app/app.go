package test_containers_app

import (
	"context"
	"io"
	"os"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/go-faster/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
)

type Container struct {
	container    testcontainers.Container
	externalHost string
	externalPort string
	cfg          *Opts
}

func NewContainer(ctx context.Context, opts ...optsFunc) (*Container, error) {
	cfg := &Opts{
		Port:           defaultAppPort,
		ContainerName:  defaultContainerName,
		Logger:         logger.DefaultInfoLogger(),
		DockerFileDir:  defaultDockerFileDir,
		DockerFileName: defaultDockerFileName,
		Env:            map[string]string{},
		Networks:       []string{},
		LogOutput:      os.Stdout,
	}
	cfg.StartupWait = wait.ForListeningPort(nat.Port(cfg.Port)).WithStartupTimeout(defaultStartupTimeout)

	for _, opt := range opts {
		opt(cfg)
	}

	req := testcontainers.ContainerRequest{
		Name: cfg.ContainerName,
		FromDockerfile: testcontainers.FromDockerfile{
			Context:        cfg.DockerFileDir,
			Dockerfile:     cfg.DockerFileName,
			BuildLogWriter: cfg.LogOutput,
		},
		Networks:     cfg.Networks,
		Env:          cfg.Env,
		WaitingFor:   cfg.StartupWait,
		ExposedPorts: []string{cfg.Port + "/tcp"},
		HostConfigModifier: func(hc *container.HostConfig) {
			hc.AutoRemove = true
		},
	}

	genericContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		return nil, errors.Errorf("failed to start app genericContainer: %v", err)
	}

	mappedPort, err := genericContainer.MappedPort(ctx, nat.Port(cfg.Port+"/tcp"))
	if err != nil {
		return nil, errors.Errorf("failed to get mapped externalPort: %v", err)
	}

	host, err := genericContainer.Host(ctx)
	if err != nil {
		return nil, errors.Errorf("failed to get genericContainer externalHost: %v", err)
	}
	go streamContainerLogs(ctx, genericContainer, cfg.LogOutput, cfg.Logger)

	return &Container{
		container:    genericContainer,
		externalHost: host,
		externalPort: mappedPort.Port(),
		cfg:          cfg,
	}, nil
}

func streamContainerLogs(
	ctx context.Context,
	container testcontainers.Container,
	out io.Writer,
	logger Logger,
) {
	logs, err := container.Logs(ctx)

	if err != nil {
		logger.Error("failed to get container logs", zap.Error(err))
		return
	}
	defer func() {
		err = logs.Close()
		if err != nil {
			logger.Error("failed to close container logs", zap.Error(err))
		}
	}()

	go func() {
		_, err = io.Copy(out, logs)
		if err != nil && !errors.Is(err, io.EOF) {
			logger.Error("error copying container logs", zap.Error(err))
		}
	}()
}
