package client_payment_v1

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
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
			PaymentMethod: common_v1.PaymentMethod(params.PaymentMethod),
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

		return nil, getStatusError(statusErr, err, params)
	}

	return &PayOrderResult{
		TransactionUUID: response.TransactionUuid,
	}, nil
}

func getStatusError(
	statusErr *status.Status,
	err error,
	params PayOrderParams,
) error {
	switch statusErr.Code() {
	case codes.Internal:
		return ErrInternalServerError{
			Err:            err,
			PayOrderParams: params,
		}
	case codes.InvalidArgument:
		return fmt.Errorf("%w: %s", ErrInvalidArguments, statusErr.Message())
	default:
		return ErrInternalServerError{
			Err:            err,
			PayOrderParams: params,
		}
	}
}
