package service_pay_order

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

func (s *PayOrderServiceImpl) PayOrder(
	ctx context.Context,
	params PayOrderMethodParams,
) (*PayOrderMethodReturn, error) {
	randomDelay := math.Ceil(rand.Float64() * maxDelaySeconds)
	time.Sleep(time.Duration(randomDelay) * time.Second)

	fmt.Printf("Оплата прошла успешно, transaction_uuid: %s\n", params.OrderUUID)
	transactionUUID := uuid.New().String()

	return &PayOrderMethodReturn{
		TransactionUUID: transactionUUID,
	}, nil
}
