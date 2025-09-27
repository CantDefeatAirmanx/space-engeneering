package service_assemblies_watcher

import "context"

func (a *AssembliesWatcherServiceImpl) WatchAssemblies(ctx context.Context) error {
	return a.serviceConsumer.ConsumeAssemblyCompletedMessage(ctx)
}
