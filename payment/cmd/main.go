package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	configs_payment "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/payment"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

func main() {
	paymentService := NewPaymentServiceServer()

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

	payment_v1.RegisterPaymentServiceServer(grpcServer, paymentService)

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

type PaymentServiceServer struct {
	payment_v1.UnimplementedPaymentServiceServer
}

func NewPaymentServiceServer() payment_v1.PaymentServiceServer {
	return &PaymentServiceServer{}
}

const (
	maxDelaySeconds = 3
)

func (s *PaymentServiceServer) PayOrder(
	ctx context.Context,
	req *payment_v1.PayOrderRequest,
) (*payment_v1.PayOrderResponse, error) {
	fmt.Printf("PayOrder: %+v\n", req.OrderUuid)

	randomDelay := math.Ceil(rand.Float64() * maxDelaySeconds)
	time.Sleep(time.Duration(randomDelay) * time.Second)

	transactionUUID := uuid.New().String()
	fmt.Printf("PayOrder: %s, User: %s, Method: %s, Transaction: %s\n",
		req.OrderUuid,
		req.UserUuid,
		req.PaymentMethod,
		transactionUUID,
	)

	return &payment_v1.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil
}
