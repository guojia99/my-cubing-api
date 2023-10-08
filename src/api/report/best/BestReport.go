package best

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type ReportResponse struct {
	BestSingle map[model.Project]model.Score `json:"BestSingle"`
	BestAvg    map[model.Project]model.Score `json:"BestAvg"`
}

func Report(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bestSingle, bestAvg := svc.Core.GetAllProjectBestScores()
		ctx.JSON(http.StatusOK, ReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
		})
	}
}
