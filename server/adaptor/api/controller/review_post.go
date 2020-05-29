package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/k-ueki/AGPlus/server/domain/entity"

	"github.com/k-ueki/AGPlus/server/adaptor/api/input"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	ReviewPostController struct {
		service.ReviewPostService
	}
)

func NewReviewPostController(db *gorm.DB) *ReviewPostController {
	return &ReviewPostController{service.NewReviewPostService(db)}
}

func (c *ReviewPostController) Store(ctx *gin.Context) {
	classID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messages": "failed to bind parameters"})
		return
	}
	param := input.ReviewClassRequest{}
	if err := ctx.BindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind query"))
		return
	}

	//TODO:User認証
	user := entity.User{
		ID:   1,
		Name: "User1",
	}

	err = c.ReviewPostService.Store(classID, &user, &param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.New("failed to store"))
		return
	}

	ctx.JSON(http.StatusCreated, nil)
	return
}

func (c *ReviewPostController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messages": "failed to bind parameters"})
		return
	}

	if err := c.ReviewPostService.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.New("failed to delete review"))
		return
	}

	ctx.JSON(http.StatusOK, nil)
	return
}
