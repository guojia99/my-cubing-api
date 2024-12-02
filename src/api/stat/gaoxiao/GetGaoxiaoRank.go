package gaoxiao

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func GetGaoXRank(svc *svc.Context) gin.HandlerFunc {
	var data map[string][]WcaResult
	go func() {
		data = GetSorAllResults()
		ticker := time.NewTicker(time.Minute * 60)
		for {
			select {
			case <-ticker.C:
				data = GetSorAllResults()
			}
		}
	}()

	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, data)
	}
}
