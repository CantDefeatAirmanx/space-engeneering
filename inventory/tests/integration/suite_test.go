//go:build integration

package integration

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/config"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	test_containers_app "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/test_containers/app"
	test_containers_mongo "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/test_containers/mongo"
	test_containers_network "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/test_containers/network"
	test_containers_path "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/test_containers/path"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

const (
	appContainerName   = "inventory_test_app"
	mongoContainerName = "inventory_test_mongodb"
	networkName        = "inventory_test_network"
)

type TestEnvironment struct {
	Network      *test_containers_network.Network
	Mongo        *test_containers_mongo.Container
	App          *test_containers_app.Container
	Closer       closer.Closer
	CloserStatus chan closer.Status
}

var (
	testEnv *TestEnvironment

	suiteCtx    context.Context
	suiteCancel context.CancelFunc
)

func TestIntegration(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "Integration tests")
}

var _ = BeforeSuite(func() {
	fmt.Println("Integration test _ 1")

	if err := logger.Init(
		logger.WithLevel(logger.LevelDebug),
	); err != nil {
		panic(fmt.Errorf("failed to init logger: %w", err))
	}

	suiteCtx, suiteCancel = context.WithCancel(context.Background())

	cfgPath := filepath.Join(
		test_containers_path.GetProjectRoot(),
		"inventory",
		".env",
	)

	if err := config.LoadConfig(
		config.WithEnvPath(cfgPath),
		config.WithIsTest(true),
	); err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
	fmt.Println("Integration test _ 2")

	testEnv = setUpTestEnvironment(suiteCtx)
})

var _ = AfterSuite(func() {
	<-testEnv.CloserStatus
	testEnv.Closer.CloseAll(suiteCtx)
	suiteCancel()
})

func setUpTestEnvironment(ctx context.Context) *TestEnvironment {
	logger.DefaultInfoLogger().Info("Setting up Inventory test environment")
	closer, statusCh := closer.NewCloser(
		ctx,
		closer.WithLogger(logger.DefaultInfoLogger()),
	)

	go func() {
		if r := recover(); r != nil {
			go closer.CloseAll(ctx)
			<-statusCh
		}
	}()

	network, err := test_containers_network.NewNetwork(
		ctx,
		networkName,
	)
	closer.AddNamed("network", func(ctx context.Context) error {
		return network.Remove(ctx)
	})
	if err != nil {
		panic(fmt.Errorf("failed to create network: %w", err))
	}

	mongo, err := test_containers_mongo.NewContainer(
		ctx,
		test_containers_mongo.WithContainerName(mongoContainerName),
		test_containers_mongo.WithNetworkName(networkName),
		test_containers_mongo.WithUsername(config.Config.Mongo().Username()),
		test_containers_mongo.WithPassword(config.Config.Mongo().Password()),
		test_containers_mongo.WithDatabase(config.Config.Mongo().DBName()),
		test_containers_mongo.WithLogger(logger.DefaultInfoLogger()),
		test_containers_mongo.WithImageName(config.Config.Mongo().ImageName()),
	)
	closer.AddNamed("mongo", func(ctx context.Context) error {
		return mongo.Terminate(ctx)
	})
	if err != nil {
		panic(fmt.Errorf("failed to create mongo container: %w", err))
	}

	app, err := test_containers_app.NewContainer(
		ctx,
		test_containers_app.WithContainerName(appContainerName),
		test_containers_app.WithNetworks([]string{networkName}),
		test_containers_app.WithLogger(logger.DefaultInfoLogger()),
		test_containers_app.WithDockerFileDir(test_containers_path.GetProjectRoot()),
		test_containers_app.WithDockerFileName(filepath.Join("deploy", "inventory", "DockerFile")),
		test_containers_app.WithEnv(map[string]string{
			"GO_ENV": "production",
			"mongo__uri": fmt.Sprintf("mongodb://%s:%s@%s:27017/%s?authSource=%s",
				config.Config.Mongo().Username(),
				config.Config.Mongo().Password(),
				mongoContainerName,
				config.Config.Mongo().DBName(),
				config.Config.Mongo().AuthSource(),
			),
			"mongo__username":   config.Config.Mongo().Username(),
			"mongo__password":   config.Config.Mongo().Password(),
			"mongo__dbName":     config.Config.Mongo().DBName(),
			"mongo__port":       strconv.Itoa(config.Config.Mongo().Port()),
			"mongo__authSource": config.Config.Mongo().AuthSource(),
			"mongo__imageName":  config.Config.Mongo().ImageName(),
			"grpc__host":        config.Config.GRPC().Host(),
			"grpc__port":        strconv.Itoa(config.Config.GRPC().Port()),
			"logger__level":     "debug",
			"logger__encoder":   "json",
		}),
	)
	closer.AddNamed("app", func(ctx context.Context) error {
		return app.Terminate(ctx)
	})
	if err != nil {
		panic(fmt.Errorf("failed to create app container: %w", err))
	}

	return &TestEnvironment{
		Network:      network,
		Mongo:        mongo,
		App:          app,
		Closer:       closer,
		CloserStatus: statusCh,
	}
}
