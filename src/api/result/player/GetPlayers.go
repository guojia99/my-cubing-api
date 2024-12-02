package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type PlayersResponse struct {
	Size    int64             `json:"Size"`
	Count   int64             `json:"Count"`
	Players []GetPlayerDetail `json:"Players"`
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

		var out []GetPlayerDetail
		for _, val := range players {
			pu := svc.Core.GetPlayerUser(val)
			out = append(out, GetPlayerDetail{
				PlayerDetail: core.PlayerDetail{
					Player: val,
				},
				QQ: pu.QQ,
			})
		}

		var resp = PlayersResponse{
			Size:    int64(len(players)),
			Count:   count,
			Players: out,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
