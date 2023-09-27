package result

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type GetRecordResponse struct {
	Size    int64          `json:"Size"`
	Count   int64          `json:"Count"`
	Records []model.Record `json:"Records"`
}

func GetRecords(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "1000"))

		count, records, err := svc.Core.GetRecords(page, size)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, GetRecordResponse{
			Size:    int64(len(records)),
			Count:   count,
			Records: records,
		})
	}
}
