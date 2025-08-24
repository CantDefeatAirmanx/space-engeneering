package test_containers_app

import (
	"time"
)

const (
	defaultContainerName  = "app"
	defaultDockerFileDir  = "."
	defaultDockerFileName = "Dockerfile"
	defaultStartupTimeout = 1 * time.Minute
	defaultAppPort        = "50051"
)
