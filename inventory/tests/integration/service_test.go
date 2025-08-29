//go:build integration

package integration

import (
	"context"
	"net"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

var (
	ctx    context.Context
	cancel context.CancelFunc

	grpcConn        *grpc.ClientConn
	inventoryClient inventory_v1.InventoryServiceClient
)

var _ = Describe("Inventory Service Integration Tests", Ordered, func() {
	BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)

		url := env.App.Address()
		logger.Logger().Info("Подключение к GRPC серверу", zap.String("url", url))

		// Даем дополнительное время серверу для полного запуска
		logger.Logger().Info("Ожидание полного запуска GRPC сервера...")
		time.Sleep(3 * time.Second)

		// Проверяем доступность порта
		logger.Logger().Info("Проверка доступности TCP порта...")
		tcpConn, err := net.DialTimeout("tcp", url, 5*time.Second)
		if err != nil {
			logger.Logger().Error("TCP порт недоступен", zap.String("url", url), zap.Error(err))
			Fail("TCP порт недоступен: " + err.Error())
		}
		tcpConn.Close()
		logger.Logger().Info("TCP порт доступен", zap.String("url", url))

		// Создаем GRPC клиент
		conn, err := grpc.NewClient(
			url,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		Expect(err).To(Not(HaveOccurred()))

		logger.Logger().Info("GRPC клиент создан, проверяем health check...")

		// Проверяем health check если доступен
		healthClient := grpc_health_v1.NewHealthClient(conn)
		healthCtx, healthCancel := context.WithTimeout(ctx, 5*time.Second)
		defer healthCancel()

		_, err = healthClient.Check(healthCtx, &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			logger.Logger().Warn("Health check неуспешен (может не поддерживаться)", zap.Error(err))
		} else {
			logger.Logger().Info("Health check прошел успешно")
		}

		logger.Logger().Info("Создание inventory клиента...")

		grpcClient := inventory_v1.NewInventoryServiceClient(conn)

		grpcConn = conn
		inventoryClient = grpcClient
	})

	AfterEach(func() {
		grpcConn.Close()
		cancel()
	})

	It("should successfully start", func() {
		Expect(true).To(BeTrue())
	})

	It("should list all parts", func() {
		_, err := inventoryClient.ListParts(
			ctx,
			&inventory_v1.ListPartsRequest{
				Filter: &inventory_v1.PartsFilter{},
			},
		)
		Expect(err).To(Not(HaveOccurred()))
	})
})
