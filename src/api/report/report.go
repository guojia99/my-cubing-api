package report

import (
	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/report/best"
	"github.com/guojia99/my-cubing-api/src/api/report/contest"
	"github.com/guojia99/my-cubing-api/src/api/report/player"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func AddReportRoute(g *gin.RouterGroup, svc *svc.Context) {
	rp := g.Group("/report")
	// 排行榜
	rp.GET("/record", best.Records(svc))                       // 记录列表
	rp.GET("/best/score", best.Report(svc))                    // 获取最佳成绩榜单，每个项目仅有一个单次和平均
	rp.GET("/best/all_scores", best.AllScoreReport(svc))       // 获取项目每个玩家最佳成绩
	rp.GET("/best/sor", best.SorReport(svc))                   // 获取所有角色的sor汇总榜单
	rp.GET("/best/podium", best.PodiumReport(svc))             // 获取所有玩家领奖台的排行
	rp.GET("/best/relative_sor", best.RelativeSor(svc))        // 排位分数
	rp.GET("/best/avg_relative_sor", best.AvgRelativeSor(svc)) // 平均排位分数

	// 具体到比赛
	rp.GET("/contest/:contest_id/sor", contest.SorReport(svc))       // 某比赛的sor
	rp.GET("/contest/:contest_id/score", contest.ScoreReport(svc))   // 某比赛的成绩统计
	rp.GET("/contest/:contest_id/podium", contest.PodiumReport(svc)) // 某场比赛领奖台
	rp.GET("/contest/:contest_id/record", contest.Record(svc))       // 某场比赛的记录

	// 具体到个人
	rp.GET("/player/:player_id/best", player.Best(svc))                // 某玩家的最佳成绩
	rp.GET("/player/:player_id/score", player.ScoreReport(svc))        // 某个玩家的成绩汇总
	rp.GET("/player/:player_id/podium", player.PodiumReport(svc))      // 某个玩家的领奖台
	rp.GET("/player/:player_id/record", player.Record(svc))            // 某个玩家的记录
	rp.GET("/player/:player_id/sor", player.Sor(svc))                  // 玩家的sor统计
	rp.GET("/player/:player_id/old_enemy", player.OldEnemy(svc))       // 玩家宿敌列表
	rp.GET("/player/:player_id/relative_sor", player.RelativeSor(svc)) // 玩家排位分数
}
