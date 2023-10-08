package contest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	core "github.com/guojia99/my-cubing-core"

	"github.com/guojia99/my-cubing-api/src/api/common"
	"github.com/guojia99/my-cubing-api/src/svc"
)

type CreateContestRequest = core.AddContestRequest

func CreateContest(svc *svc.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateContestRequest
		if err := ctx.Bind(&req); err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		err := svc.Core.AddContest(req)
		if err != nil {
			common.Error(ctx, http.StatusBadRequest, 0, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
