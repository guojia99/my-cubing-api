package contest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type (
	DeleteContestRequest struct {
		Id uint `uri:"contest_id"`
	}
)

func DeleteContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req DeleteContestRequest
		if err := ctx.BindUri(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		err := svc.Core.RemoveContest(req.Id)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
