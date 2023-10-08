package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"k8s.io/apimachinery/pkg/util/cache"
)

func NewRateMiddleware(secondRete int) gin.HandlerFunc {
	c := *cache.NewLRUExpireCache(10000)
	return func(ctx *gin.Context) {
		clientIp := ctx.ClientIP()

		var r *rate.Limiter
		if val, ok := c.Get(clientIp); ok {
			r = val.(*rate.Limiter)
		} else {
			r = rate.NewLimiter(rate.Limit(1), secondRete)
		}

		if !r.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many req"})
			return
		}
		ctx.Next()
	}
}
