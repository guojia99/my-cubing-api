package xlog

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type XLog struct {
	gorm.Model

	Title       string `json:"Title,omitempty"`
	CreatedTime string `json:"CreatedTime,omitempty"`
	Area        string `json:"Area,omitempty"`
	Messages    string `json:"Messages,omitempty"`
}

func AddXLogRoute(g *gin.RouterGroup, authMiddleware gin.HandlerFunc, svc *svc.Context) {
	xLog := g.Group("x-log")
	xLog.GET("/", GetXLogs(svc))
	xLog.PUT("/", authMiddleware, AddXLog(svc))
	xLog.DELETE("/:x_id", authMiddleware, DeleteXLog(svc))
}
