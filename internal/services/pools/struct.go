package pools

type PoolMsg struct {
	PoolID     int   `json:"pool_id"`
	PoolValues []int `json:"pool_values"`
}

type addPoolValuesRequest struct {
	PoolID     int   `json:"pool_id" binding:"required"`
	PoolValues []int `json:"pool_values" binding:"required"`
}

type queryPoolRequest struct {
	PoolID     int     `json:"pool_id" binding:"required"`
	Percentile float64 `json:"percentile" binding:"required"`
}
