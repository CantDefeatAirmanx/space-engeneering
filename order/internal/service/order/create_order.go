package service_order

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	client_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/inventory/v1"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
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
		return nil, &model_order.ErrOrderInternal{
			OrderUUID: "",
			Err:       err,
		}
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
		if errors.Is(err, client_inventory_v1.ErrInvalidArguments) {
			return nil, &model_order.ErrOrderInvalidArguments{
				OrderUUID: orderUUID.String(),
				Err:       err,
			}
		}

		return nil, &model_order.ErrOrderInternal{
			OrderUUID: orderUUID.String(),
			Err:       err,
		}
	}

	if len(parts) != len(params.PartUuids) {
		return nil, &model_order.ErrOrderInvalidArguments{
			OrderUUID: orderUUID.String(),
			Err:       errors.New("some parts not found"),
		}
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

	err = s.orderRepository.CreateOrder(orderCtx, repository_order_converter.ToRepository(&order))
	if err != nil {
		return nil, &model_order.ErrOrderInternal{
			OrderUUID: order.OrderUUID,
			Err:       err,
		}
	}

	return &CreateOrderResult{
		OrderUUID:  order.OrderUUID,
		TotalPrice: order.TotalPrice,
	}, nil
}
