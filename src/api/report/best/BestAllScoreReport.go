package best

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type AllScoreReportResponse struct {
	BestSingle map[model.Project][]model.Score `json:"BestSingle"`
	BestAvg    map[model.Project][]model.Score `json:"BestAvg"`
}

func AllScoreReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bestSingle, bestAvg := svc.Core.GetBestScore()
		ctx.JSON(http.StatusOK, AllScoreReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
		})
	}
}
