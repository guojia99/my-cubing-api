package contest

import (
	"github.com/gin-gonic/gin"
	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
	"github.com/guojia99/my-cubing-core/model"
	"net/http"
)

type (
	UpdateContestGroupRequest struct {
		ContestID uint   `json:"ContestID"`
		Groups    string `json:"groups"`
	}
)

func UpdateContestGroup(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req UpdateContestGroupRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误"+err.Error())
			return
		}
		var contest model.Contest
		if err := svc.DB.First(&contest, "id = ?", req.ContestID).Error; err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, "错误", err.Error())
			return
		}

		contest.GroupID = req.Groups
		svc.DB.Save(&contest)
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
