package player

import (
	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/svc"
)

func CreatePlayerImages(svc *svc.Context) gin.HandlerFunc {
	_ = svc.DB.AutoMigrate(&Images{})
	return func(ctx *gin.Context) {
		// todo 接入图片创建
	}
}
