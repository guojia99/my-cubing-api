package player

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func GetPlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}

		out, err := svc.Core.GetPlayer(req.Id)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}

		ctx.JSON(http.StatusOK, out)
	}
}
