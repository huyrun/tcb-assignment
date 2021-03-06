package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type header struct {
	Authorization string `header:"authorization"`
}

func Middleware(authSrv Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := &header{}
		err := ctx.BindHeader(header)
		if err != nil {
			logrus.WithError(err).Error("Middleware, bind header")
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}

		tokenString := strings.Replace(header.Authorization, "Bearer ", "", -1)

		token, err := validateToken(tokenString, authSrv.getSecretJWT())
		if err != nil {
			logrus.WithError(err).Errorf("Middleware, token: %s", token)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if token == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logrus.Infof("Middleware, cast claims failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if expiredTime, ok := claims["expired_time"]; ok {
			if t, ok := expiredTime.(float64); ok {
				if t-float64(time.Now().Unix()) < 0 {
					ctx.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			} else {
				logrus.Infof("Middleware, cast expired_time failed")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		ctx.Next()
	}
}

func validateToken(encodedToken string, secretJWT string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretJWT), nil
		}

		return nil, errors.New("Invalid token")
	})
}
