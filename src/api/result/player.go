/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/17 下午6:13.
 *  * Author: guojia(https://github.com/guojia99)
 */

package result

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-core/model"

	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	PlayerRequest struct {
		Id uint `uri:"player_id"`
	}
)

func GetPlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		out, err := svc.Core.GetPlayer(req.Id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, out)
	}
}

type PlayersResponse struct {
	Size    int64          `json:"Size"`
	Count   int64          `json:"Count"`
	Players []model.Player `json:"Players"`
}

func GetPlayers(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "100"))
		count, players, err := svc.Core.GetPlayers(page, size)
		if err != nil {
			ctx.JSON(http.StatusOK, err.Error())
			return
		}
		var resp = PlayersResponse{
			Size:    int64(len(players)),
			Count:   count,
			Players: players,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func CreatePlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.Player
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := svc.Core.AddPlayer(req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func UpdatePlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.Player
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := svc.Core.UpdatePlayer(req.ID, req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func DeletePlayer(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PlayerRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := svc.Core.RemovePlayer(req.Id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
