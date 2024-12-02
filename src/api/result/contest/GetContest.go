package contest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	GetContestRequest struct {
		ContestID uint `uri:"contest_id"`
	}
)

func GetContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetContestRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "(获取比赛)创建比赛需要正确参数")
			return
		}
		contest, err := svc.Core.GetContest(req.ContestID)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		ctx.JSON(http.StatusOK, contest)
	}
}
