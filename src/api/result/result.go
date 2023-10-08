package result

import (
	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/result/contest"
	"github.com/guojia99/my-cubing-api/src/api/result/player"
	"github.com/guojia99/my-cubing-api/src/api/result/score"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func AddResultRoute(api *gin.RouterGroup, authMiddleware gin.HandlerFunc, svc *svc.Context) {
	{ // 玩家
		api.GET("/players", player.GetPlayers(svc))                                // 获取玩家列表
		api.GET("/player/:player_id", player.GetPlayer(svc))                       // 获取单个玩家信息
		api.POST("/player", authMiddleware, player.CreatePlayer(svc))              //  添加玩家
		api.PUT("/player", authMiddleware, player.UpdatePlayer(svc))               // 修改玩家
		api.DELETE("/player/:player_id", authMiddleware, player.DeletePlayer(svc)) // 删除玩家

		api.GET("/player/:player_id/images", player.GetPlayerImages(svc))                     // 获取玩家头像、背景等信息
		api.POST("/player/:player_id/images", authMiddleware, player.CreatePlayerImages(svc)) // 创建玩家头像、背景等信息
	}
	{ // 比赛
		api.GET("/contests", contest.GetContests(svc))                                 // 获取比赛列表
		api.GET("/contest/:contest_id", contest.GetContest(svc))                       // 获取单场比赛信息
		api.POST("/contest", authMiddleware, contest.CreateContest(svc))               // 添加比赛
		api.DELETE("/contest/:contest_id", authMiddleware, contest.DeleteContest(svc)) // 删除比赛
	}
	{ // 成绩
		api.GET("/score/player/:player_id/contest/:contest_id", authMiddleware, score.GetScores(svc)) // 获取某场比赛玩家的所有成绩
		api.POST("/score", authMiddleware, score.CreateScore(svc))                                    // 上传成绩
		api.PUT("/score/end_contest", authMiddleware, score.EndContest(svc))                          // 结束比赛并统计
		api.DELETE("/score/:score_id", authMiddleware, score.DeleteScore(svc))                        // 删除成绩
	}
}
