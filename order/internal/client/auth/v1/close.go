package client_auth_v1

func (c *authV1GrpcClient) Close() error {
	return c.conn.Close()
}
