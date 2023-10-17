package pre_score

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type GetPreScoresResp struct {
	Size   int64            `json:"Size"`
	Count  int64            `json:"Count"`
	Scores []model.PreScore `json:"Scores"`
}

func GetPreScores(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "100"))
		final, _ := strconv.Atoi(ctx.DefaultQuery("final", "1"))

		count, scores, err := svc.Core.GetPreScores(page, size, core.Bool(final))
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "获取待处理成绩错误"+err.Error())
			return
		}
		var resp = GetPreScoresResp{
			Size:   int64(len(scores)),
			Count:  count,
			Scores: scores,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
