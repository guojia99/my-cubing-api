package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type GetPlayerDetail struct {
	core.PlayerDetail

	QQ string
}

func GetPlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}

		out, err := svc.Core.GetPlayer(req.Id)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "查询不到该选手")
			return
		}

		pu := svc.Core.GetPlayerUser(out.Player)
		resp := GetPlayerDetail{
			PlayerDetail: out,
			QQ:           pu.QQ,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
