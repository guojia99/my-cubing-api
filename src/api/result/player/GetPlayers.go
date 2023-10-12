package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type PlayersResponse struct {
	Size    int64          `json:"Size"`
	Count   int64          `json:"Count"`
	Players []model.Player `json:"Players"`
}

func GetPlayers(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "100"))
		count, players, err := svc.Core.GetPlayers(page, size)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}
		var resp = PlayersResponse{
			Size:    int64(len(players)),
			Count:   count,
			Players: players,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
