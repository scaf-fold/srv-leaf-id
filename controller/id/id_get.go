package id

import (
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/srv-leaf-id/global"
	"github.com/scaf-fold/tools/pkg/courier"
	"net/http"
)

type IdGenerator struct {
	courier.MethodGet
	Size int64 `form:"size"`
}

func (receiver *IdGenerator) Path() string {
	return "/id"
}

func (receiver *IdGenerator) Context(ctx *gin.Context) {
	err := ctx.ShouldBindQuery(receiver)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ids, err := global.ID.Ids(receiver.Size)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, ids)
}
