/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/19 下午5:59.
 *  * Author: guojia(https://github.com/guojia99)
 */

package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	CreateScoreRequest struct {
		PlayerName string `json:"PlayerName"` // 为兼容迁移v1数据的接口

		PlayerID  uint               `json:"PlayerID"`
		ContestID uint               `json:"ContestID"`
		Project   model.Project      `json:"Project"`
		RouteID   uint               `json:"RouteNum"`
		Penalty   model.ScorePenalty `json:"Penalty"`
		Results   []float64          `json:"Results"`
	}
)

func CreateScore(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateScoreRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(req.PlayerName) > 0 {
			var p model.Player
			if err := svc.DB.First(&p, "name = ?", req.PlayerName).Error; err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			req.PlayerID = p.ID
		}

		if err := svc.Core.AddScore(core.AddScoreRequest{
			PlayerID:  req.PlayerID,
			ContestID: req.ContestID,
			Project:   req.Project,
			RoundId:   req.RouteID,
			Result:    req.Results,
			Penalty:   req.Penalty,
		}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

type (
	DeleteScoreRequest struct {
		ScoreID uint `uri:"score_id"`
	}
)

func DeleteScore(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req DeleteScoreRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := svc.Core.RemoveScore(req.ScoreID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

type EndContestRequest struct {
	ContestID uint `json:"ContestID"`
}

func EndContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req EndContestRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := svc.Core.EndContestScore(req.ContestID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

type GetScoresRequest struct {
	PlayerID  uint `uri:"player_id"`
	ContestID uint `uri:"contest_id"`
}

func GetScores(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetScoresRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		scores, err := svc.Core.GetScoreByPlayerContest(req.PlayerID, req.ContestID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, scores)
	}
}
