package api_order_v1

import (
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func handleServiceError[T any](err error) (T, error) {
	var result T

	switch {
	case errors.Is(err, model_order.ErrOrderNotFound):
		if errResp, ok := any(&order_v1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}).(T); ok {
			return errResp, nil
		}

	case errors.Is(err, model_order.ErrOrderInvalidArguments):
		if errResp, ok := any(&order_v1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}).(T); ok {
			return errResp, nil
		}

	case errors.Is(err, model_order.ErrOrderConflict):
		if errResp, ok := any(&order_v1.ConflictError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		}).(T); ok {
			return errResp, nil
		}
	}

	if errResp, ok := any(&order_v1.InternalServerError{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf("%s: %s", internalServerErrorMessage, err.Error()),
	}).(T); ok {
		logger.Logger().Error(
			internalServerErrorMessage,
			zap.String("errDesc", err.Error()),
		)
		return errResp, nil
	}

	return result, fmt.Errorf("unable to handle service error for type %T", result)
}
