package infra

import (
	"net/http"
	"tcb-assignment/internal/services/auth"
	"tcb-assignment/internal/services/pools"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func ProvideRestAPIHandler(
	authSrv auth.Service,
	poolService pools.Service,
) RestAPIHandler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(healthcheck.Default())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := router.Group("/v1")
	{
		//authorizedV1.Use(auth.Middleware(authSrv))

		poolGroup := v1.Group("/pool")
		{
			pools.Handler(poolGroup, poolService)
		}
	}

	return router
}
