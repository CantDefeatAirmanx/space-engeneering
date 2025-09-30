package platform_redis

type RedisCache interface {
	String() StringCache
	Set() SetCache
	Hash() HashCache
}
