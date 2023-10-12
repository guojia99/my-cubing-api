package contest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type SorReportResponse struct {
	Single map[model.SorStatisticsKey][]core.SorScore `json:"Single"`
	Avg    map[model.SorStatisticsKey][]core.SorScore `json:"Avg"`
}

func SorReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "获取不到该比赛Sor数据")
			return
		}

		single, avg := svc.Core.GetContestSor(req.ContestID)
		ctx.JSON(http.StatusOK, SorReportResponse{
			Single: single,
			Avg:    avg,
		})
	}
}
