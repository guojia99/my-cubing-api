package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type BestResponse struct {
	Best map[model.Project]core.RankScore `json:"Best"`
	Avg  map[model.Project]core.RankScore `json:"Avg"`
}

func Best(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "(最佳成绩)查询需要正确的选手ID")
			return
		}
		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "(最佳成绩)查询不到该选手")
			return
		}

		best, avg := svc.Core.GetPlayerBestScore(player.ID)
		ctx.JSON(http.StatusOK, BestResponse{
			Best: best,
			Avg:  avg,
		})
	}
}
