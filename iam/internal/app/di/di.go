package di

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"

type DiContainer struct {
	closer closer.Closer
}

func NewDiContainer(closer closer.Closer) *DiContainer {
	return &DiContainer{
		closer: closer,
	}
}
