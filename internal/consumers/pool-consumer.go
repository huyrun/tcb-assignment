package consumers

import (
	"context"
	"encoding/json"
	"tcb-assignment/internal/pkg/message-queue/consumer"
	"tcb-assignment/internal/services/pools"
)

type PoolConsumer interface {
	Start()
	Consumers() []consumer.Consumer
}

type poolConsumer struct {
	poolService pools.Service
	consumers   []consumer.Consumer
}

func NewPoolConsumer(poolService pools.Service, numOfConsumer int) *poolConsumer {
	if numOfConsumer <= 0 {
		panic("Number of consumer must be positive")
	}

	p := &poolConsumer{
		poolService: poolService,
		consumers:   make([]consumer.Consumer, numOfConsumer),
	}

	for i := 0; i < numOfConsumer; i++ {
		p.consumers[i] = consumer.NewConsumer(
			consumer.Name(p.name()),
			consumer.HandlerFunc(p.consume))
	}

	return p
}

func (p *poolConsumer) name() string {
	return "pool-consumer"
}

func (p *poolConsumer) Start() {
	for _, consumer := range p.consumers {
		consumer.Start()
	}
}

func (p *poolConsumer) Consumers() []consumer.Consumer {
	return p.consumers
}

func (p *poolConsumer) consume(msg []byte) error {
	poolMsg, err := decode(msg)
	if err != nil {
		return err
	}

	return p.poolService.AddPoolValuesWithRetry(context.Background(), poolMsg.PoolID, poolMsg.PoolValues, pools.NumOfRetries)
}

func decode(msg []byte) (*pools.PoolMsg, error) {
	poolMsg := new(pools.PoolMsg)
	err := json.Unmarshal(msg, poolMsg)
	if err != nil {
		return nil, err
	}

	return poolMsg, nil
}
