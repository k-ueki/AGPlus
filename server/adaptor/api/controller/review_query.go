package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	ReviewQueryController struct {
		service.ReviewGetService
	}
)

func NewReviewQueryController(db *gorm.DB) *ReviewQueryController {
	return &ReviewQueryController{
		ReviewGetService: service.ReviewGetService{
			ReviewGetRepository: repository.ReviewGetRepository{
				DB: db,
			},
		},
	}
}

func (c *ReviewQueryController) List(ctx *gin.Context) {
	reviews, err := c.ReviewGetService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, reviews)
	return
}
