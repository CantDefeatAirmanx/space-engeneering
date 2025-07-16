package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/api/payment/v1"
	service_pay_order "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/service/pay_order"
	configs_payment "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/payment"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

func main() {
	paymentService := service_pay_order.NewPayOrderServiceImpl()
	api := api_payment_v1.NewApi(api_payment_v1.NewApiParams{
		PayOrderService: paymentService,
	})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configs_payment.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			log.Fatalf("Failed to close listener: %v", err)
		}
	}()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	payment_v1.RegisterPaymentServiceServer(grpcServer, api)

	go func() {
		fmt.Println("Payment service started")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	grpcServer.GracefulStop()
	log.Println("Payment service stopped")
}
