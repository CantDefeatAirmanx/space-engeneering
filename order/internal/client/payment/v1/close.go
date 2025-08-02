package client_payment_v1

func (c *paymentV1GrpcClient) Close() error {
	return c.conn.Close()
}
