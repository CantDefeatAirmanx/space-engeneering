package client_inventory_v1

func (c *inventoryV1GrpcClient) Close() error {
	return c.conn.Close()
}
