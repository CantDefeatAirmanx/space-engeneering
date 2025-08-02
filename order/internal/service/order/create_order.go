package service_order

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	client_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/inventory/v1"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	model_part "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part"
)

const (
	partsTimeout = 3 * time.Second
	orderTimeout = 3 * time.Second
)

func (s *OrderServiceImpl) CreateOrder(
	ctx context.Context,
	params CreateOrderParams,
) (*CreateOrderResult, error) {
	orderUUID, err := uuid.NewV7()
	if err != nil {
		return nil, model_order.ErrOrderInternal
	}

	partsCtx, partsCancel := context.WithTimeout(ctx, partsTimeout)
	defer partsCancel()

	parts, err := s.inventoryClient.ListParts(
		partsCtx,
		client_inventory_v1.ListPartsParams{
			Uuids: params.PartUuids,
		},
	)
	if err != nil {
		if errors.Is(err, model_part.ErrPartInvalidArguments) {
			return nil, fmt.Errorf("%w: %s", model_order.ErrOrderInvalidArguments, "some parts not found")
		}

		return nil, model_order.ErrOrderInternal
	}

	if len(parts) != len(params.PartUuids) {
		return nil, fmt.Errorf("%w: %s", model_order.ErrOrderInvalidArguments, "some parts not found")
	}

	totalPrice := 0.0
	for _, part := range parts {
		totalPrice += part.Price
	}

	order := model_order.Order{
		OrderUUID:  orderUUID.String(),
		UserUUID:   params.UserUUID,
		PartUuids:  params.PartUuids,
		Status:     model_order.OrderStatusPendingPayment,
		TotalPrice: totalPrice,
	}

	orderCtx, orderCancel := context.WithTimeout(ctx, orderTimeout)
	defer orderCancel()

	err = s.orderRepository.CreateOrder(orderCtx, order)
	if err != nil {
		return nil, err
	}

	return &CreateOrderResult{
		OrderUUID:  order.OrderUUID,
		TotalPrice: order.TotalPrice,
	}, nil
}
