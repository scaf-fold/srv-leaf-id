package id

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/srv-leaf-id/global"
	"github.com/scaf-fold/tools/pkg/courier"
	"net/http"
)

type IdParser struct {
	courier.MethodGet
	Id string `uri:"id"`
}

func (p *IdParser) Path() string {
	return "/parser/:id"
}

func (p *IdParser) Context(ctx *gin.Context) {
	err := ctx.ShouldBindUri(&p)
	fmt.Println(">>>>", p.Id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	snow, err := global.ID.IdInverse(p.Id, global.StartTime)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, snow)
}
