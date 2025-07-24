package service_pay_order_tests

import (
	"github.com/google/uuid"

	service_pay_order "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/service/pay_order"
)

func (s *TestingSuite) TestPayOrder() {
	tcases := []struct {
		name string
		err  error
	}{
		{
			name: "success",
			err:  nil,
		},
	}

	for _, tc := range tcases {
		s.Run(tc.name, func() {
			res, err := s.service.PayOrder(
				s.ctx,
				service_pay_order.PayOrderMethodParams{
					OrderUUID:     uuid.New().String(),
					UserUUID:      uuid.New().String(),
					PaymentMethod: 0,
				},
			)

			s.NotEmpty(res.TransactionUUID)
			s.ErrorIs(err, tc.err)
		})
	}
}
