package client_auth_v1

import (
	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

type AuthV1Client interface {
	GetRawClient() auth_v1.AuthServiceClient
	interfaces.WithClose
}
