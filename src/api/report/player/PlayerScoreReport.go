package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type ScoreReportResponse struct {
	BestSingle []model.Score          `json:"BestSingle"`
	BestAvg    []model.Score          `json:"BestAvg"`
	Scores     []core.ScoresByContest `json:"Scores"`
}

func ScoreReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "(成绩)查询需要正确的选手ID")
			return
		}
		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "(成绩)查询不到该选手")
			return
		}
		bestSingle, bestAvg, scores := svc.Core.GetPlayerScore(player.ID)
		ctx.JSON(http.StatusOK, ScoreReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
			Scores:     scores,
		})
	}
}
