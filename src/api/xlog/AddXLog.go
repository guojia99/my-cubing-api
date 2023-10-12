package xlog

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func AddXLog(svc *svc.Context) gin.HandlerFunc {
	_ = svc.DB.AutoMigrate(&XLog{})

	return func(ctx *gin.Context) {
		var req XLog
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 1, "错误"+err.Error())
			return
		}
		svc.DB.Save(&XLog{
			Title:       req.Title,
			CreatedTime: req.CreatedTime,
			Area:        req.Area,
			Messages:    req.Messages,
		})
	}
}
