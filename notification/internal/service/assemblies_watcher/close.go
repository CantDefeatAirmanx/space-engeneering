package service_assemblies_watcher

func (a *AssembliesWatcherServiceImpl) Close() error {
	return a.assembliesConsumer.Close()
}
