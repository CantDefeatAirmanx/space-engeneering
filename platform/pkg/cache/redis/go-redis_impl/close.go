package platform_redis_redisgo

import "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"

var _ interfaces.WithClose = (*SingleNodeImpl)(nil)

func (s *SingleNodeImpl) Close() error {
	return s.client.Close()
}
