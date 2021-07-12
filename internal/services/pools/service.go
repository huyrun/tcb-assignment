package pools

import (
	"context"
	"fmt"
	"tcb-assignment/internal/pkg/message-queue/producer"
	"tcb-assignment/internal/services/cache"
	"time"

	"github.com/cenkalti/backoff/v4"
)

const (
	NumOfRetries = 5
)

type Service interface {
	IsPoolIdExist(pooID int) bool
	AddPoolValues(ctx context.Context, pooID int, values []int) error
	AddPoolValuesWithRetry(ctx context.Context, pooID int, values []int, numsOfRetries int) error
	QueryPool(ctx context.Context, poolID int, percentile float64) (float64, int, error)
	PublishAddPoolValues(pooID int, values []int) error

	// For compare
	AddPoolValuesSync(ctx context.Context, pooID int, values []int) (bool, error)
}

type service struct {
	cache        cache.Cache
	poolRepo     PoolRepo
	poolProducer producer.Producer
}

func NewPoolService(
	cache cache.Cache,
	poolRepo PoolRepo,
	poolProducer producer.Producer,
) *service {
	return &service{
		cache:        cache,
		poolRepo:     poolRepo,
		poolProducer: poolProducer,
	}
}

func (s *service) IsPoolIdExist(pooID int) bool {
	return s.cache.IsIntegerKeyExist(pooID)
}

func (s *service) AddPoolValues(ctx context.Context, pooID int, values []int) error {
	return s.poolRepo.Save(ctx, pooID, values)
}

func (s *service) AddPoolValuesSync(ctx context.Context, pooID int, values []int) (bool, error) {
	return s.poolRepo.SaveV2(ctx, pooID, values)
}

func (s *service) AddPoolValuesWithRetry(ctx context.Context, pooID int, values []int, numsOfRetries int) error {
	b := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(numsOfRetries))

	return backoff.RetryNotify(func() (err error) {
		err = s.AddPoolValues(ctx, pooID, values)
		return err
	}, b, func(err error, t time.Duration) {
		fmt.Printf("AddPoolValues fail err = %v, retry after %v\n", err, t)
	})
}

func (s *service) QueryPool(ctx context.Context, poolID int, percentile float64) (float64, int, error) {
	return s.poolRepo.QueryByPercentile(ctx, poolID, percentile)
}

func (s *service) PublishAddPoolValues(pooID int, values []int) error {
	msg := &PoolMsg{
		PoolID:     pooID,
		PoolValues: values,
	}

	return backoff.Retry(func() error {
		err := s.cache.AddIntegerKey(pooID)
		if err != nil {
			return err
		}

		return s.poolProducer.Publish(msg)
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), NumOfRetries))
}
