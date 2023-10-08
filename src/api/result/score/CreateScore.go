package score

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	CreateScoreRequest struct {
		PlayerName string `json:"PlayerName"` // 为兼容迁移v1数据的接口

		PlayerID  uint               `json:"PlayerID"`
		ContestID uint               `json:"ContestID"`
		Project   model.Project      `json:"Project"`
		RouteID   uint               `json:"RouteNum"`
		Penalty   model.ScorePenalty `json:"Penalty"`
		Results   []float64          `json:"Results"`
	}
)

func CreateScore(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateScoreRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}

		if len(req.PlayerName) > 0 {
			var p model.Player
			if err := svc.DB.First(&p, "name = ?", req.PlayerName).Error; err != nil {
				common.Error(ctx, http.StatusBadRequest, 0, err)
				return
			}
			req.PlayerID = p.ID
		}

		if err := svc.Core.AddScore(core.AddScoreRequest{
			PlayerID:  req.PlayerID,
			ContestID: req.ContestID,
			Project:   req.Project,
			RoundId:   req.RouteID,
			Result:    req.Results,
			Penalty:   req.Penalty,
		}); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
