package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func UpdatePlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonPlayerRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}

		// 更新玩家内容
		p := model.Player{
			Model: model.Model{
				ID: req.ID,
			},
			Name:       req.Name,
			WcaID:      req.WcaID,
			ActualName: req.ActualName,
			TitlesVal:  req.TitlesVal,
		}
		if err := svc.Core.UpdatePlayer(req.ID, p); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "更新玩家信息错误"+err.Error())
			return
		}

		// 更新玩家用户信息
		pu := model.PlayerUser{
			PlayerID: req.ID,
			LoginID:  req.LoginID,
			QQ:       req.QQ,
			WeChat:   req.WeChat,
			Phone:    req.Phone,
		}

		var err error
		player := svc.Core.GetPlayerUser(p)
		if player.ID != 0 {
			err = svc.Core.UpdatePlayerUser(p, pu)
		} else {
			err = svc.Core.AddPlayerUser(p, pu)
		}
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "写入用户信息"+err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
