package service_topics_configurator

import "context"

func (tc *TopicsConfigurator) Close(_ context.Context) error {
	return tc.client.Close()
}
