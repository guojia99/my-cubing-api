/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/19 下午2:18.
 *  * Author: guojia(https://github.com/guojia99)
 */

package report

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type PlayerBestResponse struct {
	Best map[model.Project]core.RankScore `json:"Best"`
	Avg  map[model.Project]core.RankScore `json:"Avg"`
}

func PlayerBest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		best, avg := svc.Core.GetPlayerBestScore(player.ID)
		ctx.JSON(http.StatusOK, PlayerBestResponse{
			Best: best,
			Avg:  avg,
		})
	}
}

type PlayerRequest struct {
	PlayerId uint `uri:"player_id"`
}

func PlayerPodiumReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, svc.Core.GetPlayerPodiums(player.ID))
	}
}

type PlayerScoreReportResponse struct {
	BestSingle []model.Score          `json:"BestSingle"`
	BestAvg    []model.Score          `json:"BestAvg"`
	Scores     []core.ScoresByContest `json:"Scores"`
}

func PlayerScoreReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		bestSingle, bestAvg, scores := svc.Core.GetPlayerScore(player.ID)
		ctx.JSON(http.StatusOK, PlayerScoreReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
			Scores:     scores,
		})
	}
}

func PlayerRecord(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, svc.Core.GetPlayerRecord(req.PlayerId))
	}
}

type PlayerSorResponse struct {
	Single map[model.SorStatisticsKey]core.SorScore `json:"Single"`
	Avg    map[model.SorStatisticsKey]core.SorScore `json:"Avg"`
}

func PlayerSor(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		signal, avg := svc.Core.GetPlayerSor(player.ID)
		ctx.JSON(http.StatusOK, PlayerSorResponse{Single: signal, Avg: avg})
	}
}

func PlayerOldEnemy(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var player model.Player
		if err := svc.DB.Where("id = ?", req.PlayerId).First(&player).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, svc.Core.GetPlayerOldEnemy(player.ID))
	}
}
