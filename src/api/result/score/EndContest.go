package score

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type EndContestRequest struct {
	ContestID uint `json:"ContestID"`
}

func EndContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req EndContestRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}
		if err := svc.Core.EndContestScore(req.ContestID); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
