package xlog

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/svc"
)

func GetXLogs(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var out []XLog
		svc.DB.Find(&out)
		ctx.JSON(http.StatusOK, out)
	}
}
