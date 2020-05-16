package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/k-ueki/AGPlus/server/adaptor/api/input"

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
	classID := ctx.Param("id")
	q := input.ReviewClassRequest{}
	if err := ctx.BindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind query"))
		return
	}

	fmt.Println(q)
}
