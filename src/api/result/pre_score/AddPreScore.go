package pre_score

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func AddPreScore(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req core.AddPreScoreRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 1, "参数错误:"+err.Error())
			return
		}
		if err := svc.Core.AddPreScore(req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 1, "添加错误:"+err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
