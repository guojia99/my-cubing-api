package score

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type GetScoresRequest struct {
	PlayerID  uint `uri:"player_id"`
	ContestID uint `uri:"contest_id"`
}

func GetScores(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetScoresRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}
		scores, err := svc.Core.GetScoreByPlayerContest(req.PlayerID, req.ContestID)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}
		ctx.JSON(http.StatusOK, scores)
	}
}
