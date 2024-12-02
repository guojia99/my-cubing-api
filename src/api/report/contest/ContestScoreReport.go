/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/19 下午2:17.
 *  * Author: guojia(https://github.com/guojia99)
 */

package contest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type ScoreReportResponse struct {
	Scores map[model.Project][]core.RoutesScores `json:"Scores"`
}

func ScoreReport(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "获取不到该比赛成绩列表")
			return
		}
		ctx.JSON(http.StatusOK, ScoreReportResponse{
			Scores: svc.Core.GetContestScore(req.ContestID),
		})
	}
}
