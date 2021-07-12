package pools

import "context"

type PoolRepo interface {
	Save(ctx context.Context, poolID int, poolValues []int) error
	QueryByPercentile(ctx context.Context, poolID int, percentile float64) (float64, int, error)

	// For compare
	SaveV2(ctx context.Context, poolID int, poolValues []int) (bool, error)
}
