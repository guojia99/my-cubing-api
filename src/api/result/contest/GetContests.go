package contest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	GetContestsResponse struct {
		Size     int64           `json:"Size"`
		Count    int64           `json:"Count"`
		Contests []model.Contest `json:"Contests"`
	}
)

func GetContests(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Try to get the cache.
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "20"))
		typ, _ := ctx.GetQuery("type")

		count, contests, err := svc.Core.GetContests(page, size, typ)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "获取比赛列表错误"+err.Error())
			return
		}

		// Convert to interface contents.
		var resp = GetContestsResponse{
			Count:    count,
			Contests: contests,
			Size:     int64(len(contests)),
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
