package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	ReviewPostController struct {
		service.ReviewPostService
	}
)

func NewReviewPostController(db *gorm.DB) *ReviewPostController {
	return &ReviewPostController{
		ReviewPostService: service.ReviewPostService{
			ReviewPostRepository: repository.ReviewPostRepository{
				DB: db,
			},
		},
	}
}

func (c *ReviewPostController) Store(ctx *gin.Context) {

}
