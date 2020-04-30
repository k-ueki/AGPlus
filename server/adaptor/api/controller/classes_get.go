package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	ClassGetController struct {
		service.ClassGetService
	}
)

func (c *ClassGetController) List(ctx *gin.Context) {
	classes, err := c.ClassGetService.List()
	if err != nil {
		// TODO: ErrorHandlingしっかり
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, classes)
	return
}
