package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

// ValidMiddleware 授权中间件
func ValidMiddleware(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("auth in %s", ctx.ClientIP())
		token := ctx.GetHeader("Authorization")
		if token == "" {
			common.Error(ctx, http.StatusUnauthorized, 21, "权限不足")
			return
		}

		// token 可用，直接通过
		if _, ok := svc.Cache.Get(token); ok {
			ctx.Next()
			return
		}

		// 查token 是否合法
		var admin Admin
		if err := svc.DB.Where("token = ?", token).First(&admin).Error; err != nil {
			common.Error(ctx, http.StatusUnauthorized, 22, "权限不足")
			return
		}
		if time.Now().Sub(admin.Timeout) > 0 {
			common.Error(ctx, http.StatusUnauthorized, 23, "权限过期")
			return
		}

		_ = svc.Cache.Add(admin.Token, admin.UserName, time.Minute)
		ctx.Next()
	}
}
