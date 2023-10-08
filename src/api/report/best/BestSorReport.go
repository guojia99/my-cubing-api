package best

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type SorReportResponse struct {
	BestSingle map[model.SorStatisticsKey][]core.SorScore `json:"BestSingle"`
	BestAvg    map[model.SorStatisticsKey][]core.SorScore `json:"BestAvg"`
}

func SorReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bestSingle, bestAvg := svc.Core.GetSor()
		ctx.JSON(http.StatusOK, SorReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
		})
	}
}
