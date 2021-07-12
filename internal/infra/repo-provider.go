package infra

import (
	"tcb-assignment/internal/services/pools"
	"tcb-assignment/internal/storages"
)

func ProvidePoolRepo() pools.PoolRepo {
	return storages.NewPoolRepo()
}
