package service_pay_order

var _ PayOrderService = (*PayOrderServiceImpl)(nil)

type PayOrderServiceImpl struct{}

func NewPayOrderServiceImpl() *PayOrderServiceImpl {
	return &PayOrderServiceImpl{}
}

const (
	maxDelaySeconds = 3
)
