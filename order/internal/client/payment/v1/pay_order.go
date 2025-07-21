package client_payment_v1

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

func (client *paymentV1GrpcClient) PayOrder(
	ctx context.Context,
	params PayOrderParams,
) (*PayOrderResult, error) {
	response, err := client.grpcClient.PayOrder(
		ctx,
		&payment_v1.PayOrderRequest{
			OrderUuid:     params.OrderUUID,
			UserUuid:      params.UserUUID,
			PaymentMethod: payment_v1.PaymentMethod(params.PaymentMethod),
		},
	)
	if err != nil {
		statusErr, ok := status.FromError(err)

		if !ok {
			return nil, ErrInternalServerError{
				Err:            err,
				PayOrderParams: params,
			}
		}

		return handleStatusError(statusErr, err, params)
	}

	return &PayOrderResult{
		TransactionUUID: response.TransactionUuid,
	}, nil
}

func handleStatusError(
	statusErr *status.Status,
	err error,
	params PayOrderParams,
) (*PayOrderResult, error) {
	switch statusErr.Code() {
	case codes.Internal:
		return nil, ErrInternalServerError{
			Err:            err,
			PayOrderParams: params,
		}
	case codes.InvalidArgument:
		return nil, fmt.Errorf("%w: %s", ErrInvalidArguments, statusErr.Message())
	default:
		return nil, ErrInternalServerError{
			Err:            err,
			PayOrderParams: params,
		}
	}
}
