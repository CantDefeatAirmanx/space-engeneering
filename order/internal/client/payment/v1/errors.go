package client_payment_v1

import (
	"errors"
	"fmt"
)

var ErrInvalidArguments = errors.New("invalid arguments")

type ErrInternalServerError struct {
	Err error
	PayOrderParams
}

func (e ErrInternalServerError) Error() string {
	return fmt.Sprintf(
		"payment error: %v, order_uuid: %s, user_uuid: %s, payment_method: %d",
		e.Err,
		e.OrderUUID,
		e.UserUUID,
		e.PaymentMethod,
	)
}

func (e ErrInternalServerError) Unwrap() error {
	return e.Err
}

func (e ErrInternalServerError) Is(target error) bool {
	_, ok := target.(*ErrInternalServerError)
	return ok
}
