package model_payment_method

type PaymentMethod int

const (
	PaymentMethodUnknown       PaymentMethod = 0
	PaymentMethodCard          PaymentMethod = 1
	PaymentMethodSPB           PaymentMethod = 2
	PaymentMethodCreditCard    PaymentMethod = 3
	PaymentMethodInvestorMoney PaymentMethod = 4
)
