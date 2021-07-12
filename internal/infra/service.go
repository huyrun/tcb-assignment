package infra

import (
	cache2 "tcb-assignment/internal/pkg/cache"
	"tcb-assignment/internal/pkg/message-queue/producer"
	"tcb-assignment/internal/services/auth"
	"tcb-assignment/internal/services/cache"
	"tcb-assignment/internal/services/pools"
)

func ProvideAuthService(cfg *AppConfig) auth.Service {
	return auth.NewAuthService(cfg.SecretJWT)
}

func ProvideCacheService() cache.Cache {
	return cache2.NewCache()
}

func ProvidePoolService(cache cache.Cache, poolRepo pools.PoolRepo, poolProducer producer.Producer) pools.Service {
	return pools.NewPoolService(cache, poolRepo, poolProducer)
}
