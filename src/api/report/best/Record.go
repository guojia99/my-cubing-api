package best

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type RecordsResponse struct {
	Size    int64          `json:"Size"`
	Count   int64          `json:"Count"`
	Records []model.Record `json:"Records"`
}

func Records(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "1000"))

		count, records, err := svc.Core.GetRecords(page, size)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "获取不到记录列表数据")
			return
		}
		ctx.JSON(http.StatusOK, RecordsResponse{
			Size:    int64(len(records)),
			Count:   count,
			Records: records,
		})
	}
}
