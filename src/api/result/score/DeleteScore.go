package score

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	DeleteScoreRequest struct {
		ScoreID uint `uri:"score_id"`
	}
)

func DeleteScore(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req DeleteScoreRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		if err := svc.Core.RemoveScore(req.ScoreID); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
