package service_assemblies_watcher

import (
	"context"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
)

type AssembliesWatcherService interface {
	WatchAssemblies(ctx context.Context) error
	interfaces.WithClose
}
