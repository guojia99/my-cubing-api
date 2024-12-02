package pre_score

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func NeglectPreScore(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 1, "参数错误"+err.Error())
			return
		}

		err := svc.Core.ProcessPreScore(core.ProcessPreScoreRequest{
			Id:           req.ID,
			Processor:    req.Processor,
			FinishDetail: model.FinishDetailNeglect,
		})
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 1, "执行操作错误"+err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
