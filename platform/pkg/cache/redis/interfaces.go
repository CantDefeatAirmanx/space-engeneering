package platform_redis

import "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"

type RedisCache interface {
	String() StringCache
	Set() SetCache
	Hash() HashCache
	Key() KeyCache
	interfaces.WithClose
}
