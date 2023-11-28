package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

// ValidToken 获取合法token
func ValidToken(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetTokenRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "密码格式错误")
			return
		}
		if req.UserName == "" || req.PassWord == "" {
			common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "密码格式错误")
			return
		}

		var admin Admin
		if err := svc.DB.Where("user_name = ?", req.UserName).First(&admin).Error; err != nil {
			common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "查询不到角色")
			return
		}
		if admin.Password != req.PassWord {
			common.Error(ctx, http.StatusNetworkAuthenticationRequired, 20, "密码错误")
			return
		}
		if time.Now().Sub(admin.Timeout) > 0 || admin.Token == "" {
			admin.Token = generateUniqueToken(req.UserName, time.Now().Unix())
			admin.Timeout = time.Now().Add(time.Hour * 48)
			svc.DB.Save(&admin)
		}

		resp := GetTokenResponse{
			Ts:    admin.Timeout.Unix(),
			Token: admin.Token,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
