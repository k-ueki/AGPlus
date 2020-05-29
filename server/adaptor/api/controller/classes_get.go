package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/k-ueki/AGPlus/server/adaptor/api/converter"

	"github.com/k-ueki/AGPlus/server/adaptor/api/input"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	ClassGetController struct {
		service.ClassGetService
	}
)

func NewClassGetController(db *gorm.DB) *ClassGetController {
	return &ClassGetController{service.NewClassGetService(db)}
}

func (c *ClassGetController) List(ctx *gin.Context) {
	param := input.ListClass{}
	if err := ctx.BindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind query"))
		return
	}

	classes, err := c.ClassGetService.List(converter.ConvertListClassParamToQuery(&param))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, classes)
	return
}

func (c *ClassGetController) Show(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind parameters"))
		return
	}

	class, err := c.ClassGetService.Show(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, class)
	return
}
