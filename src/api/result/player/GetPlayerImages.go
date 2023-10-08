package player

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/svc"
)

func GetPlayerImages(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusOK, gin.H{})
			return
		}

		var out Images
		if err := svc.DB.Where("player_id = ?", req.Id).First(&out).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{})
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}
