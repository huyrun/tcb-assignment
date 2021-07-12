package pools

func queryPoolRequestValidator(req queryPoolRequest) error {
	if req.Percentile <= 0 || req.Percentile >= 100 {
		return ErrBadRequest
	}

	return nil
}
