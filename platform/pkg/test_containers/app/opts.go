package test_containers_app

import (
	"io"

	"github.com/testcontainers/testcontainers-go/wait"
)

type Opts struct {
	Port           string
	Logger         Logger
	ContainerName  string
	DockerFileDir  string
	DockerFileName string
	Env            map[string]string
	Networks       []string
	LogOutput      io.Writer
	StartupWait    wait.Strategy
}

type optsFunc func(opts *Opts)

func WithPort(port string) optsFunc {
	return func(opts *Opts) {
		opts.Port = port
	}
}

func WithLogger(logger Logger) optsFunc {
	return func(opts *Opts) {
		opts.Logger = logger
	}
}

func WithContainerName(containerName string) optsFunc {
	return func(opts *Opts) {
		opts.ContainerName = containerName
	}
}

func WithDockerFileDir(dockerFileDir string) optsFunc {
	return func(opts *Opts) {
		opts.DockerFileDir = dockerFileDir
	}
}

func WithDockerFileName(dockerFileName string) optsFunc {
	return func(opts *Opts) {
		opts.DockerFileName = dockerFileName
	}
}

func WithEnv(env map[string]string) optsFunc {
	return func(opts *Opts) {
		opts.Env = env
	}
}

func WithLogOutput(logOutput io.Writer) optsFunc {
	return func(opts *Opts) {
		opts.LogOutput = logOutput
	}
}

func WithStartupWait(startupWait wait.Strategy) optsFunc {
	return func(opts *Opts) {
		opts.StartupWait = startupWait
	}
}

func WithNetworks(networks []string) optsFunc {
	return func(opts *Opts) {
		opts.Networks = networks
	}
}
