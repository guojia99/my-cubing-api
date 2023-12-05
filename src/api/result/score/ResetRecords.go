package score

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func ResetRecords(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := svc.Core.ReSetRecords()
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
