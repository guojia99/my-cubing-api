/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/17 下午6:55.
 *  * Author: guojia(https://github.com/guojia99)
 */

package result

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	ContestRequest struct {
		ContestID uint `uri:"contest_id"`
	}
)

func GetContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ContestRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		contest, err := svc.Core.GetContest(req.ContestID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, contest)
	}
}

type (
	GetContestsResponse struct {
		Size     int64           `json:"Size"`
		Count    int64           `json:"Count"`
		Contests []model.Contest `json:"Contests"`
	}
)

func GetContests(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Try to get the cache.
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "20"))
		typ, _ := ctx.GetQuery("type")

		count, contests, err := svc.Core.GetContests(page, size, typ)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert to interface contents.
		var resp = GetContestsResponse{
			Count:    count,
			Contests: contests,
			Size:     int64(len(contests)),
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func CreateContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req core.AddContestRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := svc.Core.AddContest(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

type (
	DeleteContestRequest struct {
		Id uint `uri:"contest_id"`
	}
)

func DeleteContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req DeleteContestRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := svc.Core.RemoveContest(req.Id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
