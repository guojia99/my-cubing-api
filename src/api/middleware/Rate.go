package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"k8s.io/apimachinery/pkg/util/cache"

	"github.com/guojia99/my-cubing-api/src/api/common"
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
			common.Error(ctx, http.StatusTooManyRequests, 1, "请求过于频繁")
			return
		}
		ctx.Next()
	}
}
