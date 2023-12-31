/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/7/19 下午2:17.
 *  * Author: guojia(https://github.com/guojia99)
 */

package contest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

func Record(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CommonRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "获取不到该比赛的记录数据")
			return
		}
		ctx.JSON(http.StatusOK, svc.Core.GetContestRecord(req.ContestID))
	}
}
