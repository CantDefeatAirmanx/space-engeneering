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
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	configs_order "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/order"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

const (
	internalServerErrorMessage = "Internal server error"
)

func main() {
	ctx := context.Background()
	orderRepo := NewOrderRepositoryMap()

	paymentGrpcClient, pErr := grpc.NewClient(
		configs_order.PaymentServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	inventoryGrpcClient, iErr := grpc.NewClient(
		configs_order.InventoryServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if pErr != nil || iErr != nil {
		log.Fatalf("Ошибка при создании клиента gRPC: %v", pErr)
	}
	paymentGrpcClient.Connect()
	inventoryGrpcClient.Connect()

	for _, con := range []*grpc.ClientConn{paymentGrpcClient, inventoryGrpcClient} {
		con.WaitForStateChange(ctx, connectivity.Ready)
	}

	paymentClient := payment_v1.NewPaymentServiceClient(paymentGrpcClient)
	inventoryClient := inventory_v1.NewInventoryServiceClient(inventoryGrpcClient)

	orderHandler := NewOrderHandler(NewOrderHandlerParams{
		OrderRepo:       orderRepo,
		PaymentClient:   paymentClient,
		InventoryClient: inventoryClient,
	})

	orderServer, pErr := order_v1.NewServer(orderHandler)
	if pErr != nil {
		log.Fatalf("Ошибка при создании сервера заказов: %v", pErr)
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	// router.Use(middleware.Timeout(configs_order.Timeout))

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
	if err := paymentGrpcClient.Close(); err != nil {
		log.Printf("Ошибка при закрытии payment gRPC клиента: %v", err)
	}
	if err := inventoryGrpcClient.Close(); err != nil {
		log.Printf("Ошибка при закрытии inventory gRPC клиента: %v", err)
	}

	log.Println("Сервер остановлен")
}

type OrderHandler struct {
	repo            OrderRepository
	paymentClient   payment_v1.PaymentServiceClient
	inventoryClient inventory_v1.InventoryServiceClient
}

type NewOrderHandlerParams struct {
	OrderRepo       OrderRepository
	PaymentClient   payment_v1.PaymentServiceClient
	InventoryClient inventory_v1.InventoryServiceClient
}

func NewOrderHandler(params NewOrderHandlerParams) *OrderHandler {
	return &OrderHandler{
		repo:            params.OrderRepo,
		paymentClient:   params.PaymentClient,
		inventoryClient: params.InventoryClient,
	}
}

func (handler *OrderHandler) CreateOrder(
	ctx context.Context,
	req *order_v1.CreateOrderRequestBody,
) (order_v1.CreateOrderRes, error) {
	orderUUID := uuid.New().String()

	inventoryCtx, cancel := context.WithTimeout(
		ctx,
		10*time.Second,
	)
	defer cancel()

	orderParts, err := handler.inventoryClient.ListParts(
		inventoryCtx,
		&inventory_v1.ListPartsRequest{
			Filter: &inventory_v1.PartsFilter{
				Uuids: req.PartUuids,
			},
		},
	)
	if err != nil {
		return &order_v1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: internalServerErrorMessage,
		}, nil
	}

	if len(orderParts.Parts) != len(req.PartUuids) {
		notAvailableParts := []string{}
		for _, uuid := range req.PartUuids {
			for _, orderPart := range orderParts.Parts {
				if orderPart.Uuid == uuid {
					break
				}
				notAvailableParts = append(notAvailableParts, uuid)
			}
		}

		return &order_v1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Not all parts are available: %v", notAvailableParts),
		}, nil
	}

	order := order_v1.Order{
		OrderUUID:  orderUUID,
		UserUUID:   req.UserUUID,
		PartUuids:  req.PartUuids,
		TotalPrice: 100,
		Status:     order_v1.OrderStatusPENDINGPAYMENT,
	}

	err = handler.repo.CreateOrder(ctx, order)
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

func (handler *OrderHandler) GetOrder(
	ctx context.Context,
	params order_v1.GetOrderParams,
) (order_v1.GetOrderRes, error) {
	order, err := handler.repo.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		return &order_v1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Order %s not found", params.OrderUUID),
		}, nil
	}

	return &order, nil
}

var paymentMethodMap = map[order_v1.PaymentMethod]payment_v1.PaymentMethod{
	order_v1.PaymentMethodUNKNOWN:       payment_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED,
	order_v1.PaymentMethodCARD:          payment_v1.PaymentMethod_PAYMENT_METHOD_CARD,
	order_v1.PaymentMethodSBP:           payment_v1.PaymentMethod_PAYMENT_METHOD_SPB,
	order_v1.PaymentMethodCREDITCARD:    payment_v1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD,
	order_v1.PaymentMethodINVESTORMONEY: payment_v1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY,
}

func (handler *OrderHandler) PayOrder(
	ctx context.Context,
	req *order_v1.PayOrderRequestBody,
	params order_v1.PayOrderParams,
) (order_v1.PayOrderRes, error) {
	order, err := handler.repo.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		return &order_v1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Order %s not found", params.OrderUUID),
		}, nil
	}

	if order.Status != order_v1.OrderStatusPENDINGPAYMENT {
		return &order_v1.ConflictError{
			Code:    http.StatusConflict,
			Message: fmt.Sprintf("Order %s is not in pending payment status", params.OrderUUID),
		}, nil
	}

	payDeadline, cancel := context.WithTimeout(
		ctx,
		10*time.Second,
	)
	defer cancel()

	paymentResponse, err := handler.paymentClient.PayOrder(
		payDeadline,
		&payment_v1.PayOrderRequest{
			OrderUuid:     params.OrderUUID,
			PaymentMethod: paymentMethodMap[req.PaymentMethod],
			UserUuid:      order.UserUUID,
		},
	)
	if err != nil {
		return &order_v1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: internalServerErrorMessage,
		}, nil
	}

	transactionUUID := paymentResponse.TransactionUuid
	newStatus := order_v1.OrderStatusPAID

	err = handler.repo.UpdateOrderFields(ctx, params.OrderUUID, OrderUpdate{
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

func (handler *OrderHandler) CancelOrder(
	ctx context.Context,
	params order_v1.CancelOrderParams,
) (order_v1.CancelOrderRes, error) {
	order, err := handler.repo.GetOrder(ctx, params.OrderUUID)
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
	err = handler.repo.UpdateOrderFields(ctx, params.OrderUUID, OrderUpdate{
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
