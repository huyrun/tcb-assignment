package infra

import (
	"tcb-assignment/internal/consumers"
	"tcb-assignment/internal/pkg/message-queue/producer"
	"tcb-assignment/internal/services/pools"
)

func ProvidePoolProducer(cfg *AppConfig) producer.Producer {
	return producer.NewProducer(producer.Topic(cfg.PoolTopic))
}

func ProvidePoolConsumer(poolService pools.Service, cfg *AppConfig) consumers.PoolConsumer {
	return consumers.NewPoolConsumer(poolService, cfg.NumOfPoolConsumer)
}
