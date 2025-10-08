package client_auth_v1

import auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"

func (c *authV1GrpcClient) GetRawClient() auth_v1.AuthServiceClient {
	return c.grpcClient
}
