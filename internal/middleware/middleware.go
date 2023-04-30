package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	err "strike_layout_template/internal/error"
	"time"
)

func HTTPResponseTimeoutMiddleware(duration time.Duration) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctx, cancelFunc := context.WithTimeout(ginCtx, duration)
		defer cancelFunc()

		select {
		case <-ctx.Done():
			ginCtx.AbortWithStatusJSON(http.StatusRequestTimeout, err.TimeOutError)
			return
		default:
			ginCtx.Next()
		}
	}

}
