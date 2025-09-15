//go:build integration

package integration

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/config"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/testcontainers"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/testcontainers/app"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/testcontainers/mongo"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/testcontainers/network"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/testcontainers/path"
)

const (
	projectName = "inventory"

	inventoryAppName    = "inventory-app"
	inventoryDockerfile = "deploy/inventory/Dockerfile"

	startupTimeout = 3 * time.Minute
)

type TestEnvironment struct {
	Network *network.Network
	Mongo   *mongo.Container
	App     *app.Container
}

const testsTimeout = 5 * time.Minute

var (
	testEnvironment *TestEnvironment

	suiteCtx    context.Context
	suiteCancel context.CancelFunc
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UFO Service Integration Test Suite")
}

var _ = BeforeSuite(func() {
	err := logger.Init(
		logger.WithLevel(logger.LevelDebug),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}

	suiteCtx, suiteCancel = context.WithTimeout(context.Background(), testsTimeout)

	cfgPath := filepath.Join(
		path.GetProjectRoot(),
		"inventory",
		".env",
	)
	if err := config.LoadConfig(
		config.WithEnvPath(cfgPath),
		config.WithIsTest(true),
	); err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	logger.Logger().Info("Starting test environment...")
	testEnvironment = setupTestEnvironment(suiteCtx)
})

var _ = AfterSuite(func() {
	logger.Logger().Info("Test suite finished")
	if testEnvironment != nil {
		teardownTestEnvironment(suiteCtx, testEnvironment)
	}
	suiteCancel()
})

func setupTestEnvironment(ctx context.Context) *TestEnvironment {
	logger.Logger().Info("🚀 Preparing test environment...")

	generatedNetwork, err := network.NewNetwork(ctx, projectName)
	if err != nil {
		logger.Logger().Fatal("failed to create network", zap.Error(err))
	}
	logger.Logger().Info("✅ Network created")

	mongoUsername := config.Config.Mongo().Username()
	mongoPassword := config.Config.Mongo().Password()
	mongoImageName := config.Config.Mongo().ImageName()
	mongoDatabase := config.Config.Mongo().DBName()

	grpcPort := config.Config.GRPC().Port()

	generatedMongo, err := mongo.NewContainer(ctx,
		mongo.WithNetworkName(generatedNetwork.Name()),
		mongo.WithContainerName(testcontainers.GenerateMongoContainerName()),
		mongo.WithImageName(mongoImageName),
		mongo.WithDatabase(mongoDatabase),
		mongo.WithAuth(mongoUsername, mongoPassword),
		mongo.WithLogger(logger.DefaultInfoLogger()),
	)
	if err != nil {
		cleanupTestEnvironment(ctx, &TestEnvironment{Network: generatedNetwork})
		logger.Logger().Fatal("failed to start MongoDB container", zap.Error(err))
	}
	logger.Logger().Info("✅ MongoDB container started")

	projectRoot := path.GetProjectRoot()
	waitStrategy := wait.ForListeningPort(
		nat.Port(strconv.Itoa(grpcPort) + "/tcp"),
	).WithStartupTimeout(startupTimeout)

	uniqueAppName := fmt.Sprintf("%s-%d", inventoryAppName, time.Now().Unix())
	appContainer, err := app.NewContainer(ctx,
		app.WithName(uniqueAppName),
		app.WithPort(strconv.Itoa(grpcPort)),
		app.WithDockerfile(projectRoot, inventoryDockerfile),
		app.WithNetwork(generatedNetwork.Name()),
		app.WithEnv(map[string]string{
			"GO_ENV": "prod",

			"grpc__host": "0.0.0.0",
			"grpc__port": strconv.Itoa(grpcPort),

			"mongo__host":       generatedMongo.Config().ContainerName,
			"mongo__port":       "27017",
			"mongo__username":   generatedMongo.Config().Username,
			"mongo__password":   generatedMongo.Config().Password,
			"mongo__dbName":     generatedMongo.Config().Database,
			"mongo__authSource": generatedMongo.Config().AuthDB,
			"mongo__imageName":  generatedMongo.Config().ImageName,
			"mongo__network":    generatedNetwork.Name(),

			"logger__level":   "debug",
			"logger__encoder": "json",
		}),
		app.WithLogOutput(os.Stdout),
		app.WithStartupWait(waitStrategy),
		app.WithLogger(logger.Logger()),
	)
	if err != nil {
		cleanupTestEnvironment(ctx, &TestEnvironment{Network: generatedNetwork, Mongo: generatedMongo})
		logger.Logger().Fatal("failed to start application container", zap.Error(err))
	}

	err = appContainer.Status(ctx)
	if err != nil {
		logger.Logger().Warn("failed to get application container status", zap.Error(err))
	}

	time.Sleep(5 * time.Second)

	address := appContainer.Address()
	logger.Logger().Info("Checking GRPC port availability", zap.String("address", address))

	logger.Logger().Info("✅ Application container started")

	logger.Logger().Info("🎉 Test environment ready")
	return &TestEnvironment{
		Network: generatedNetwork,
		Mongo:   generatedMongo,
		App:     appContainer,
	}
}

func teardownTestEnvironment(ctx context.Context, env *TestEnvironment) {
	logger.Logger().Info("🧹 Cleaning up test environment...")

	cleanupTestEnvironment(ctx, env)

	logger.Logger().Info("✅ Test environment cleaned up")
}

func cleanupTestEnvironment(ctx context.Context, env *TestEnvironment) {
	if env.App != nil {
		if err := env.App.Terminate(ctx); err != nil {
			logger.Logger().Error("failed to stop application container", zap.Error(err))
		} else {
			logger.Logger().Info("🛑 Application container stopped")
		}
	}

	if env.Mongo != nil {
		if err := env.Mongo.Terminate(ctx); err != nil {
			logger.Logger().Error("failed to stop MongoDB container", zap.Error(err))
		} else {
			logger.Logger().Info("🛑 MongoDB container stopped")
		}
	}

	if env.Network != nil {
		if err := env.Network.Remove(ctx); err != nil {
			logger.Logger().Error("failed to remove network", zap.Error(err))
		} else {
			logger.Logger().Info("🛑 Network removed")
		}
	}
}
