package service_pay_order

import (
	"context"
	"math"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

func (s *PayOrderServiceImpl) PayOrder(
	ctx context.Context,
	params PayOrderMethodParams,
) (*PayOrderMethodReturn, error) {
	randomDelay := math.Ceil(rand.Float64() * maxDelaySeconds)
	time.Sleep(time.Duration(randomDelay) * time.Second)

	logger.Logger().Info("Successfully paid order", zap.String("transaction_uuid", params.OrderUUID))
	transactionUUID := uuid.New().String()

	return &PayOrderMethodReturn{
		TransactionUUID: transactionUUID,
	}, nil
}
