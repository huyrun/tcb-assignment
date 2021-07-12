package pools

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrInternalError     = errors.New("Internal error")
	ErrBadRequest        = errors.New("Request is invalid")
	ErrNoPool            = errors.New("Pool do not exist")
	ErrFailedCalculation = errors.New("Failed calculate quantile")
)

func handleError(ctx *gin.Context, err error) {
	switch err {
	case ErrNoPool:
		makeErrorResponse(ctx, http.StatusOK, "NOT_EXIST", err.Error())
	case ErrFailedCalculation:
		makeErrorResponse(ctx, http.StatusOK, "FAILED_CALCULATION", err.Error())
	case ErrBadRequest:
		makeErrorResponse(ctx, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
	default:
		makeErrorResponse(ctx, http.StatusInternalServerError, "INTERNAL_ERROR", ErrInternalError.Error())
	}
}

func makeErrorResponse(ctx *gin.Context, statusCode int, code, message string) {
	ctx.JSON(statusCode, gin.H{
		"status":  "ERROR",
		"code":    code,
		"message": message,
	})
}
