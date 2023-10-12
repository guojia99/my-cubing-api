package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func PodiumReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "(领奖台)查询需要正确的选手ID")
			return
		}
		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "(领奖台)查询不到该选手")
			return
		}

		ctx.JSON(http.StatusOK, svc.Core.GetPlayerPodiums(player.ID))
	}
}
