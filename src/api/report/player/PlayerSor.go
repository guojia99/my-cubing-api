package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type SorResponse struct {
	Single map[model.SorStatisticsKey]core.SorScore `json:"Single"`
	Avg    map[model.SorStatisticsKey]core.SorScore `json:"Avg"`
}

func Sor(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}

		signal, avg := svc.Core.GetPlayerSor(player.ID)
		ctx.JSON(http.StatusOK, SorResponse{
			Single: signal,
			Avg:    avg,
		})
	}
}
