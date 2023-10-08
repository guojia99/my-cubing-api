package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func CreatePlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.Player
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}

		if err := svc.Core.AddPlayer(req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
