package storages

import (
	"context"
	"sync"
	tcb_assignment "tcb-assignment"
	"tcb-assignment/internal/pkg/data-structure/rbtree"
	"tcb-assignment/internal/services/pools"
)

type repo struct {
	mu      sync.RWMutex
	storage map[int]*rbtree.Rbtree
}

func NewPoolRepo() *repo {
	return &repo{
		storage: make(map[int]*rbtree.Rbtree),
	}
}

func (r *repo) Save(ctx context.Context, poolID int, poolValues []int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	rbt, ok := r.storage[poolID]
	if !ok {
		if tcb_assignment.VisualizeRbtreeMode {
			rbt = rbtree.NewRbtree(rbtree.Visualize())
		} else {
			rbt = rbtree.NewRbtree()
		}

		r.storage[poolID] = rbt
	}

	rbt.AddMany(poolValues)

	return nil
}

func (r *repo) SaveV2(ctx context.Context, poolID int, poolValues []int) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var isInsert bool
	rbt, ok := r.storage[poolID]
	if !ok {
		isInsert = true
		if tcb_assignment.VisualizeRbtreeMode {
			rbt = rbtree.NewRbtree(rbtree.Visualize())
		} else {
			rbt = rbtree.NewRbtree()
		}

		r.storage[poolID] = rbt
	}

	rbt.AddMany(poolValues)

	return isInsert, nil
}

func (r *repo) QueryByPercentile(ctx context.Context, poolID int, percentile float64) (float64, int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rbt, ok := r.storage[poolID]
	if !ok {
		return 0, 0, pools.ErrNoPool
	}

	total := rbt.Len()

	p := (percentile / 100) * float64(total)
	x := int(p)

	var lower int
	if x > 0 {
		lower = rbt.Rank(x)
		if lower == -1 {
			return 0, 0, pools.ErrFailedCalculation
		}
	}

	remain := p - float64(x)
	var upper int
	if remain > 0 {
		upper = rbt.Rank(x + 1)
		if upper == -1 {
			return 0, 0, pools.ErrFailedCalculation
		}
	}

	calculatedQuantile := float64(lower) + remain*float64(upper-lower)

	return calculatedQuantile, total, nil
}
