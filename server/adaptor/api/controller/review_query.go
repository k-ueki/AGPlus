package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type (
	ReviewQueryController struct {
	}
)

func NewReviewQueryController(db *gorm.DB) *ReviewQueryController {
	return &ReviewQueryController{
		//ReviewService: service.ReviewPostService{
		//	ReviewPostRepository: repository.ReviewPostRepository{
		//		DB: db,
		//	},
		//},
	}
}

func (c *ReviewQueryController) List(ctx *gin.Context) {
}
