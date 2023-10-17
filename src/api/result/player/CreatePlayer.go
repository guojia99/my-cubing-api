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
		var req CommonPlayerRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "参数错误")
			return
		}

		// 创建玩家基本信息
		p := model.Player{
			Name:       req.Name,
			WcaID:      req.WcaID,
			ActualName: req.ActualName,
			TitlesVal:  req.TitlesVal,
		}
		if err := svc.Core.AddPlayer(p); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "写入错误"+err.Error())
			return
		}

		// 回查玩家信息
		var player model.Player
		if err := svc.DB.First(&player, "name = ?", p.Name).Error; err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "读取错误"+err.Error())
			return
		}

		pu := model.PlayerUser{
			PlayerID: player.ID,
			LoginID:  req.LoginID,
			QQ:       req.QQ,
			WeChat:   req.WeChat,
			Phone:    req.Phone,
		}
		if err := svc.Core.AddPlayerUser(player, pu); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "写入玩家用户信息"+err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
