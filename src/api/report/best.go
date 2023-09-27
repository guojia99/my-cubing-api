/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/17 下午5:01.
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

type BestReportResponse struct {
	BestSingle map[model.Project]model.Score `json:"BestSingle"`
	BestAvg    map[model.Project]model.Score `json:"BestAvg"`
}

func BestReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bestSingle, bestAvg := svc.Core.GetAllProjectBestScores()
		ctx.JSON(http.StatusOK, BestReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
		})
	}
}

type (
	BestAllScoreReportRequest struct {
		Project string `query:"project"`
	}

	BestAllScoreReportResponse struct {
		BestSingle map[model.Project][]model.Score `json:"BestSingle"`
		BestAvg    map[model.Project][]model.Score `json:"BestAvg"`
	}
	BestAllScoreReportByProjectResponse struct {
		BestSingle []model.Score `json:"BestSingle"`
		BestAvg    []model.Score `json:"BestAvg"`
	}
)

func BestAllScoreReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bestSingle, bestAvg := svc.Core.GetBestScore()
		ctx.JSON(http.StatusOK, BestAllScoreReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
		})
	}
}

type BestSorReportResponse struct {
	BestSingle map[model.SorStatisticsKey][]core.SorScore `json:"BestSingle"`
	BestAvg    map[model.SorStatisticsKey][]core.SorScore `json:"BestAvg"`
}

func BestSorReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bestSingle, bestAvg := svc.Core.GetSor()
		ctx.JSON(http.StatusOK, BestSorReportResponse{
			BestSingle: bestSingle,
			BestAvg:    bestAvg,
		})
	}
}

func BestPodiumReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, svc.Core.GetPodiums())
	}
}
