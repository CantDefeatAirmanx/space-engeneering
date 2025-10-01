package api_auth_v1

import (
	"context"

	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

func (a *Api) WhoAmI(
	ctx context.Context,
	req *auth_v1.WhoAmIRequest,
) (*auth_v1.WhoAmIResponse, error) {
	panic("unimplemented")
}
