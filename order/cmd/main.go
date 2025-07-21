package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	api_order_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/api/order/v1"
	client_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/inventory/v1"
	client_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/payment/v1"
	repository_order_map "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/map_impl"
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	configs_order "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/order"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

const (
	internalServerErrorMessage = "Internal server error"
)

func main() {
	ctx := context.Background()

	inventoryClient, inventoryErr := client_inventory_v1.NewInventoryClient(
		ctx,
	)
	if inventoryErr != nil {
		log.Fatalf("Error creating inventory client: %v", inventoryErr)
	}

	paymentClient, paymentErr := client_payment_v1.NewPaymentClient(
		ctx,
	)
	if paymentErr != nil {
		log.Fatalf("Error creating payment client: %v", paymentErr)
	}

	orderRepo := repository_order_map.NewOrderRepositoryMap()
	orderService := service_order.NewOrderService(
		service_order.NewOrderServiceParams{
			OrderRepository: orderRepo,
			InventoryClient: inventoryClient,
			PaymentClient:   paymentClient,
		},
	)

	orderHandler := api_order_v1.NewApi(
		api_order_v1.NewApiParams{
			OrderService: orderService,
		},
	)

	orderServer, pErr := order_v1.NewServer(orderHandler)
	if pErr != nil {
		log.Fatalf("Ошибка при создании сервера заказов: %v", pErr)
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(configs_order.Timeout))

	router.Mount("/", orderServer)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", configs_order.Port),
		Handler:           router,
		ReadHeaderTimeout: configs_order.ReadHeaderTimeout,
	}

	go func() {
		log.Printf("HTTP сервер запущен на порту %d", configs_order.Port)
		cerr := httpServer.ListenAndServe()
		if cerr != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %v", cerr)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Завершение работы сервера")

	ctx, cancel := context.WithTimeout(
		ctx,
		configs_order.ShutdownTimeout,
	)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при завершении работы сервера: %v", err)
	}

	// Закрываем gRPC соединения
	if err := paymentClient.Conn.Close(); err != nil {
		log.Printf("Ошибка при закрытии payment gRPC клиента: %v", err)
	}

	if err := inventoryClient.Conn.Close(); err != nil {
		log.Printf("Ошибка при закрытии part gRPC клиента: %v", err)
	}

	log.Println("Сервер остановлен")
}
