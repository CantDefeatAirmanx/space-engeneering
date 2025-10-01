package di

import (
	"context"

	api_auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/api/auth/v1"
	api_user_v1 "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/api/user/v1"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
)

type DiContainer struct {
	closer closer.Closer

	userV1API *api_user_v1.Api

	authV1API *api_auth_v1.Api
}

func NewDiContainer(closer closer.Closer) *DiContainer {
	return &DiContainer{
		closer: closer,
	}
}

func (d *DiContainer) GetUserV1API(ctx context.Context) *api_user_v1.Api {
	if d.userV1API != nil {
		return d.userV1API
	}

	d.userV1API = api_user_v1.NewApi()

	return d.userV1API
}

func (d *DiContainer) GetAuthV1API(ctx context.Context) *api_auth_v1.Api {
	if d.authV1API != nil {
		return d.authV1API
	}

	d.authV1API = api_auth_v1.NewApi()

	return d.authV1API
}
