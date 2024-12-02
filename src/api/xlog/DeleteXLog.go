package xlog

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type DeleteXLogReq struct {
	ID uint `uri:"x_id"`
}

func DeleteXLog(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req DeleteXLogReq
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 1, "错误"+err.Error())
			return
		}
		svc.DB.Delete(&XLog{}, "id = ?", req.ID)
	}
}
