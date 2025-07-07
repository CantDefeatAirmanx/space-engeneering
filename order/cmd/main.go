package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	configs_order "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

const (
	internalServerErrorMessage = "Internal server error"
)

func main() {
	orderRepo := NewOrderRepositoryMap()
	orderHandler := NewOrderHandler(orderRepo)

	orderServer, err := order_v1.NewServer(orderHandler)
	if err != nil {
		log.Fatalf("Ошибка при создании сервера заказов: %v", err)
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
		context.Background(),
		configs_order.ShutdownTimeout,
	)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		cancel()
		log.Fatalf("Ошибка при завершении работы сервера: %v", err)
	}
	log.Println("Сервер остановлен")
}

type OrderHandler struct {
	repo OrderRepository
}

func NewOrderHandler(orderRepo OrderRepository) *OrderHandler {
	return &OrderHandler{
		repo: orderRepo,
	}
}

func (o *OrderHandler) CreateOrder(
	ctx context.Context,
	req *order_v1.CreateOrderRequestBody,
) (order_v1.CreateOrderRes, error) {
	// ToDo: integration with inventory service
	orderUUID := uuid.New().String()

	order := order_v1.Order{
		OrderUUID:  orderUUID,
		UserUUID:   req.UserUUID,
		PartUuids:  req.PartUuids,
		TotalPrice: 100,
		Status:     order_v1.OrderStatusPENDINGPAYMENT,
	}

	err := o.repo.CreateOrder(ctx, order)
	if err != nil {
		return &order_v1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: internalServerErrorMessage,
		}, nil
	}

	return &order_v1.CreateOrderResponseBody{
		OrderUUID:  order.OrderUUID,
		TotalPrice: order.TotalPrice,
	}, nil
}

func (o *OrderHandler) GetOrder(
	ctx context.Context,
	params order_v1.GetOrderParams,
) (order_v1.GetOrderRes, error) {
	order, err := o.repo.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		return &order_v1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Order %s not found", params.OrderUUID),
		}, nil
	}

	return &order, nil
}

func (o *OrderHandler) PayOrder(
	ctx context.Context,
	req *order_v1.PayOrderRequestBody,
	params order_v1.PayOrderParams,
) (order_v1.PayOrderRes, error) {
	_, err := o.repo.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		return &order_v1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Order %s not found", params.OrderUUID),
		}, nil
	}
	// ToDo: integration with payment service

	transactionUUID := uuid.New().String()

	newStatus := order_v1.OrderStatusPAID

	err = o.repo.UpdateOrderFields(ctx, params.OrderUUID, OrderUpdate{
		Status:          &newStatus,
		TransactionUUID: &transactionUUID,
		PaymentMethod:   &req.PaymentMethod,
	})

	if err != nil {
		return &order_v1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &order_v1.PayOrderResponseBody{
		TransactionUUID: transactionUUID,
	}, nil
}

type conflictStatus struct {
	Value         order_v1.OrderStatus
	GetErrMessage func(orderUUID string) string
}

var conflictStatuses = []conflictStatus{
	{
		Value: order_v1.OrderStatusPAID,
		GetErrMessage: func(orderUUID string) string {
			return fmt.Sprintf("Order %s already paid", orderUUID)
		},
	},
	{
		Value: order_v1.OrderStatusCANCELLED,
		GetErrMessage: func(orderUUID string) string {
			return fmt.Sprintf("Order %s already cancelled", orderUUID)
		},
	},
}

func (o *OrderHandler) CancelOrder(
	ctx context.Context,
	params order_v1.CancelOrderParams,
) (order_v1.CancelOrderRes, error) {
	order, err := o.repo.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		return &order_v1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Order %s not found", params.OrderUUID),
		}, nil
	}

	conflictIdx := slices.IndexFunc(conflictStatuses, func(c conflictStatus) bool {
		return c.Value == order.Status
	})

	if conflictIdx != -1 {
		return &order_v1.ConflictError{
			Code:    http.StatusConflict,
			Message: conflictStatuses[conflictIdx].GetErrMessage(params.OrderUUID),
		}, nil
	}

	newStatus := order_v1.OrderStatusCANCELLED
	err = o.repo.UpdateOrderFields(ctx, params.OrderUUID, OrderUpdate{
		Status: &newStatus,
	})
	if err != nil {
		return &order_v1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: internalServerErrorMessage,
		}, nil
	}

	return &order_v1.CancelOrderOK{}, nil
}

type OrderUpdate struct {
	Status          *order_v1.OrderStatus
	TransactionUUID *string
	PaymentMethod   *order_v1.PaymentMethod
}

type OrderRepository interface {
	GetOrder(ctx context.Context, orderID string) (order_v1.Order, error)
	CreateOrder(ctx context.Context, order order_v1.Order) error
	UpdateOrderFields(ctx context.Context, orderID string, update OrderUpdate) error
	DeleteOrder(ctx context.Context, orderID string) error
}

type OrderRepositoryMap struct {
	mu     sync.RWMutex
	orders map[string]order_v1.Order
}

func NewOrderRepositoryMap() *OrderRepositoryMap {
	return &OrderRepositoryMap{
		orders: make(map[string]order_v1.Order),
		mu:     sync.RWMutex{},
	}
}

func (o *OrderRepositoryMap) GetOrder(ctx context.Context, orderID string) (order_v1.Order, error) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	order, ok := o.orders[orderID]
	if !ok {
		return order_v1.Order{}, errors.New("order not found")
	}

	return order, nil
}

func (o *OrderRepositoryMap) CreateOrder(ctx context.Context, order order_v1.Order) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.orders[order.OrderUUID] = order

	return nil
}

func (o *OrderRepositoryMap) UpdateOrderFields(
	ctx context.Context,
	orderID string,
	update OrderUpdate,
) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	order, ok := o.orders[orderID]
	if !ok {
		return fmt.Errorf("order %s not found", orderID)
	}

	if update.Status != nil {
		order.Status = *update.Status
	}

	if update.TransactionUUID != nil {
		order.TransactionUUID = order_v1.NewOptNilString(*update.TransactionUUID)
	}

	if update.PaymentMethod != nil {
		order.PaymentMethod = order_v1.NewOptPaymentMethod(*update.PaymentMethod)
	}

	o.orders[orderID] = order

	return nil
}

func (o *OrderRepositoryMap) DeleteOrder(ctx context.Context, orderID string) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	delete(o.orders, orderID)

	return nil
}
