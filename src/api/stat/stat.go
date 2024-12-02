package stat

import (
	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-api/src/api/stat/gaoxiao"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func AddStatRoute(g *gin.RouterGroup, svc *svc.Context) {
	rp := g.Group("/stat")

	rp.GET("/gaoxiao_results", gaoxiao.GetGaoXRank(svc))
}
