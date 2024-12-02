package best

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/svc"
)

func AvgRelativeSor(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, svc.Core.GetAvgRelativeSor())
	}
}
