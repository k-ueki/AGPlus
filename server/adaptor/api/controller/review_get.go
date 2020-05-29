package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	ReviewGetController struct {
		service.ReviewGetService
	}
)

func NewReviewGetController(db *gorm.DB) *ReviewGetController {
	return &ReviewGetController{service.NewReviewGetService(db)}
}

func (c *ReviewGetController) List(ctx *gin.Context) {
	reviews, err := c.ReviewGetService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, reviews)
	return
}

func (c *ReviewGetController) Show(ctx *gin.Context) {
	classID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messages": "failed to get parameters"})
		return
	}

	reviews, err := c.ReviewGetService.Show(classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, reviews)
	return
}
