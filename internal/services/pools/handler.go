package pools

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Status string

const (
	appended Status = "appended"
	inserted Status = "inserted"
)

func Handler(
	poolGroup *gin.RouterGroup,
	poolService Service,
) {
	poolGroup.POST("/add", addPoolValuesAsync(poolService))
	poolGroup.POST("/query", queryPool(poolService))

	// For compare
	poolGroup.POST("/add-sync", addPoolValuesSync(poolService))
}

func addPoolValuesAsync(poolService Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		r := addPoolValuesRequest{}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			logrus.WithError(err).Error("Bind request body")
			handleError(ctx, ErrBadRequest)
			return
		}

		var status Status
		isExist := poolService.IsPoolIdExist(r.PoolID)
		if isExist {
			status = appended
		} else {
			status = inserted
		}

		_ = poolService.PublishAddPoolValues(r.PoolID, r.PoolValues)

		ctx.JSON(http.StatusOK, gin.H{
			"status": status,
		})
	}
}

func addPoolValuesSync(poolService Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		r := addPoolValuesRequest{}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			logrus.WithError(err).Error("Bind request body")
			handleError(ctx, ErrBadRequest)
			return
		}

		isInsert, err := poolService.AddPoolValuesSync(ctx, r.PoolID, r.PoolValues)
		if err != nil {
			handleError(ctx, ErrBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": func(isInsert bool) Status {
				if isInsert {
					return inserted
				}
				return appended
			}(isInsert),
		})
	}
}

func queryPool(poolService Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		r := queryPoolRequest{}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			logrus.WithError(err).Error("Bind request body")
			handleError(ctx, ErrBadRequest)
			return
		}

		err := queryPoolRequestValidator(r)
		if err != nil {
			handleError(ctx, ErrBadRequest)
			return
		}

		calculatedQuantile, total, err := poolService.QueryPool(ctx, r.PoolID, r.Percentile)
		if err != nil {
			logrus.WithError(err).Errorf("QueryPool pool_id (%d), percentile (%d)", r.PoolID, r.Percentile)
			handleError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data": gin.H{
				"calculated_quantile": calculatedQuantile,
				"total":               total,
			},
		})
	}
}
