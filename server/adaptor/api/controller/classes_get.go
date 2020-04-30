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

func (c *ClassGetController) List(ctx *gin.Context) error {
	classes, err := c.ClassGetService.List()
	if err != nil {
		// TODO: ErrorHandlingしっかり
		return err
	}
	ctx.JSON(http.StatusOK, classes)
	return nil
}
